package client

import (
	"time"

	metrics "github.com/armon/go-metrics"
	"github.com/hashicorp/nomad/acl"
	"github.com/hashicorp/nomad/nomad/structs"
)

const (
	// policyCacheSize is the number of ACL policies to keep cached. Policies have a fetching cost
	// so we keep the hot policies cached to reduce the ACL token resolution time.
	policyCacheSize = 64

	// aclCacheSize is the number of ACL objects to keep cached. ACLs have a parsing and
	// construction cost, so we keep the hot objects cached to reduce the ACL token resolution time.
	aclCacheSize = 64

	// tokenCacheSize is the number of ACL tokens to keep cached. Tokens have a fetching cost,
	// so we keep the hot tokens cached to reduce the lookups.
	tokenCacheSize = 64

	// roleCacheSize is the number of ACL roles to keep cached. Looking up
	// roles requires an RPC call, so we keep the hot roles cached to reduce
	// the number of lookups.
	roleCacheSize = 64
)

// clientACLResolver holds the state required for client resolution
// of ACLs
type clientACLResolver struct {
	// aclCache is used to maintain the parsed ACL objects
	aclCache *structs.ACLCache[*acl.ACL]

	// policyCache is used to maintain the fetched policy objects
	policyCache *structs.ACLCache[*structs.ACLPolicy]

	// tokenCache is used to maintain the fetched token objects
	tokenCache *structs.ACLCache[*structs.ACLToken]

	// roleCache is used to maintain a cache of the fetched ACL roles. Each
	// entry is keyed by the role ID.
	roleCache *structs.ACLCache[*structs.ACLRole]
}

// init is used to setup the client resolver state
func (c *clientACLResolver) init() {
	// Create the ACL object caches
	c.aclCache = structs.NewACLCache[*acl.ACL](aclCacheSize)
	c.policyCache = structs.NewACLCache[*structs.ACLPolicy](policyCacheSize)
	c.tokenCache = structs.NewACLCache[*structs.ACLToken](tokenCacheSize)
	c.roleCache = structs.NewACLCache[*structs.ACLRole](roleCacheSize)
}

//// cachedACLValue is used to manage ACL Token, Policy, or Role cache entries
//// and their TTLs.
//type cachedACLValue struct {
//	Token     *structs.ACLToken
//	Policy    *structs.ACLPolicy
//	Role      *structs.ACLRole
//	CacheTime time.Time
//}

//// Age is the time since the token was cached
//func (c *cachedACLValue) Age() time.Duration {
//	return time.Since(c.CacheTime)
//}

// ResolveToken is used to translate an ACL Token Secret ID into
// an ACL object, nil if ACLs are disabled, or an error.
func (c *Client) ResolveToken(secretID string) (*acl.ACL, error) {
	a, _, err := c.resolveTokenAndACL(secretID)
	return a, err
}

func (c *Client) ResolveSecretToken(secretID string) (*structs.ACLToken, error) {
	_, t, err := c.resolveTokenAndACL(secretID)
	return t, err
}

func (c *Client) resolveTokenAndACL(secretID string) (*acl.ACL, *structs.ACLToken, error) {
	// Fast-path if ACLs are disabled
	if !c.GetConfig().ACLEnabled {
		return nil, nil, nil
	}
	defer metrics.MeasureSince([]string{"client", "acl", "resolve_token"}, time.Now())

	// Resolve the token value
	token, err := c.resolveTokenValue(secretID)
	if err != nil {
		return nil, nil, err
	}
	if token == nil {
		return nil, nil, structs.ErrTokenNotFound
	}

	// Give the token expiry some slight leeway in case the client and server
	// clocks are skewed.
	if token.IsExpired(time.Now().Add(2 * time.Second)) {
		return nil, nil, structs.ErrTokenExpired
	}

	// Check if this is a management token
	if token.Type == structs.ACLManagementToken {
		return acl.ManagementACL, token, nil
	}

	// Resolve the policy links within the token ACL roles.
	policyNames, err := c.resolveTokenACLRoles(secretID, token.Roles)
	if err != nil {
		return nil, nil, err
	}

	// Generate a slice of all policy names included within the token, taken
	// from both the ACL roles and the direct assignments.
	policyNames = append(policyNames, token.Policies...)

	// Resolve the policies
	policies, err := c.resolvePolicies(token.SecretID, policyNames)
	if err != nil {
		return nil, nil, err
	}

	// Resolve the ACL object
	aclObj, err := structs.CompileACLObject(c.aclCache, policies)
	if err != nil {
		return nil, nil, err
	}
	return aclObj, token, nil
}

// resolveTokenValue is used to translate a secret ID into an ACL token with caching
// We use a local cache up to the TTL limit, and then resolve via a server. If we cannot
// reach a server, but have a cached value we extend the TTL to gracefully handle outages.
func (c *Client) resolveTokenValue(secretID string) (*structs.ACLToken, error) {
	// Hot-path the anonymous token
	if secretID == "" {
		return structs.AnonymousACLToken, nil
	}

	// Lookup the token entry in the cache
	value, ok := c.tokenCache.Get(secretID)
	if ok {
		if value.Age() <= c.GetConfig().ACLTokenTTL {
			return value.First, nil
		}
	}

	// Lookup the token
	req := structs.ResolveACLTokenRequest{
		SecretID: secretID,
		QueryOptions: structs.QueryOptions{
			Region:     c.Region(),
			AllowStale: true,
		},
	}
	var resp structs.ResolveACLTokenResponse
	if err := c.RPC("ACL.ResolveToken", &req, &resp); err != nil {
		// If we encounter an error but have a cached value, mask the error and extend the cache
		if ok {
			c.logger.Warn("failed to resolve token, using expired cached value", "error", err)
			cached := value
			return cached.First, nil
		}
		return nil, err
	}

	// Cache the response (positive or negative)
	c.tokenCache.Add(secretID, resp.Token)
	return resp.Token, nil
}

// resolvePolicies is used to translate a set of named ACL policies into the objects.
// We cache the policies locally, and fault them from a server as necessary. Policies
// are cached for a TTL, and then refreshed. If a server cannot be reached, the cache TTL
// will be ignored to gracefully handle outages.
func (c *Client) resolvePolicies(secretID string, policies []string) ([]*structs.ACLPolicy, error) {
	var out []*structs.ACLPolicy
	var expired []*structs.ACLPolicy
	var missing []string

	// Scan the cache for each policy
	for _, policyName := range policies {
		// Lookup the policy in the cache
		value, ok := c.policyCache.Get(policyName)
		if !ok {
			missing = append(missing, policyName)
			continue
		}

		// Check if the cached value is valid or expired
		if value.Age() <= c.GetConfig().ACLPolicyTTL {
			out = append(out, value.Item())
		} else {
			expired = append(expired, value.Item())
		}
	}

	// Hot-path if we have no missing or expired policies
	if len(missing)+len(expired) == 0 {
		return out, nil
	}

	// Lookup the missing and expired policies
	fetch := missing
	for _, p := range expired {
		fetch = append(fetch, p.Name)
	}
	req := structs.ACLPolicySetRequest{
		Names: fetch,
		QueryOptions: structs.QueryOptions{
			Region:     c.Region(),
			AuthToken:  secretID,
			AllowStale: true,
		},
	}
	var resp structs.ACLPolicySetResponse
	if err := c.RPC("ACL.GetPolicies", &req, &resp); err != nil {
		// If we encounter an error but have cached policies, mask the error and extend the cache
		if len(missing) == 0 {
			c.logger.Warn("failed to resolve policies, using expired cached value", "error", err)
			out = append(out, expired...)
			return out, nil
		}
		return nil, err
	}

	// Handle each output
	for _, policy := range resp.Policies {
		c.policyCache.Add(policy.Name, &cachedACLValue{ // YOU ARE HERE
			Policy:    policy,
			CacheTime: time.Now(),
		})
		out = append(out, policy)
	}

	// Return the valid policies
	return out, nil
}

// resolveTokenACLRoles is used to unpack an ACL roles and their policy
// assignments into a list of ACL policy names. This can then be used to
// compile an ACL object.
//
// When roles need to be looked up from state via server RPC, we may use the
// expired cache version. This can only occur if we can fully resolve the role
// via the cache.
func (c *Client) resolveTokenACLRoles(secretID string, roleLinks []*structs.ACLTokenRoleLink) ([]string, error) {

	var (
		// policyNames tracks the resolved ACL policies which are linked to the
		// role. This is the output object and represents the authorisation
		// this role provides token bearers.
		policyNames []string

		// missingRoleIDs are the roles linked which are not found within our
		// cache. These must be looked up from the server via and RPC, so we
		// can correctly identify the policy links.
		missingRoleIDs []string

		// expiredRoleIDs are the roles linked which have been found within our
		// cache, but are expired. These must be looked up from the server via
		// and RPC, so we can correctly identify the policy links.
		expiredRoleIDs []string
	)

	for _, roleLink := range roleLinks {

		// Look within the cache to see if the role is already present. If we
		// do not find it, add the ID to our tracking, so we look this up via
		// RPC.
		raw, ok := c.roleCache.Get(roleLink.ID)
		if !ok {
			missingRoleIDs = append(missingRoleIDs, roleLink.ID)
			continue
		}

		// If the cached value is expired, add the ID to our tracking, so we
		// look this up via RPC. Otherwise, iterate the policy links and add
		// each policy name to our return object tracking.
		cached := raw.(*cachedACLValue)
		if cached.Age() <= c.GetConfig().ACLRoleTTL {
			for _, policyLink := range cached.Role.Policies {
				policyNames = append(policyNames, policyLink.Name)
			}
		} else {
			expiredRoleIDs = append(expiredRoleIDs, cached.Role.ID)
		}
	}

	// Hot-path: we were able to resolve all ACL roles via the cache and
	// generate a list of linked policy names. Therefore, we can avoid making
	// any RPC calls.
	if len(missingRoleIDs)+len(expiredRoleIDs) == 0 {
		return policyNames, nil
	}

	// Created a combined list of role IDs that we need to lookup from server
	// state.
	roleIDsToFetch := missingRoleIDs
	roleIDsToFetch = append(roleIDsToFetch, expiredRoleIDs...)

	// Generate an RPC request to detail all the ACL roles that we did not find
	// or were expired within the cache.
	roleByIDReq := structs.ACLRolesByIDRequest{
		ACLRoleIDs: roleIDsToFetch,
		QueryOptions: structs.QueryOptions{
			Region:     c.Region(),
			AuthToken:  secretID,
			AllowStale: true,
		},
	}

	var roleByIDResp structs.ACLRolesByIDResponse

	// Perform the RPC call to detail the required ACL roles. If the RPC call
	// fails, and we are only updating expired cache entries, use the expired
	// entries. This allows use to handle intermittent failures.
	err := c.RPC(structs.ACLGetRolesByIDRPCMethod, &roleByIDReq, &roleByIDResp)
	if err != nil {
		if len(missingRoleIDs) == 0 {
			c.logger.Warn("failed to resolve ACL roles, using expired cached value", "error", err)
			for _, aclRole := range roleByIDResp.ACLRoles {
				for _, rolePolicyLink := range aclRole.Policies {
					policyNames = append(policyNames, rolePolicyLink.Name)
				}
			}
			return policyNames, nil
		}
		return nil, err
	}

	// Generate a timestamp for the cache entry. We do not need to use a
	// timestamp per ACL role response integration.
	now := time.Now()

	for _, aclRole := range roleByIDResp.ACLRoles {

		// Add an entry to the cache using the generated timestamp for future
		// expiry calculations. Any existing, expired entry will be
		// overwritten.
		c.roleCache.Add(aclRole.ID, &cachedACLValue{Role: aclRole, CacheTime: now})

		// Iterate the role policy links, extracting the name and adding this
		// to our return response tracking.
		for _, rolePolicyLink := range aclRole.Policies {
			policyNames = append(policyNames, rolePolicyLink.Name)
		}
	}

	return policyNames, nil
}
