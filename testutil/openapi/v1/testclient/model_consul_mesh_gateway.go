/*
 * Nomad
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.1.3
 * Contact: support@hashicorp.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package testclient

import (
	"encoding/json"
)

// ConsulMeshGateway struct for ConsulMeshGateway
type ConsulMeshGateway struct {
	Mode *string `json:"Mode,omitempty"`
}

// NewConsulMeshGateway instantiates a new ConsulMeshGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsulMeshGateway() *ConsulMeshGateway {
	this := ConsulMeshGateway{}
	return &this
}

// NewConsulMeshGatewayWithDefaults instantiates a new ConsulMeshGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsulMeshGatewayWithDefaults() *ConsulMeshGateway {
	this := ConsulMeshGateway{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *ConsulMeshGateway) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsulMeshGateway) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *ConsulMeshGateway) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *ConsulMeshGateway) SetMode(v string) {
	o.Mode = &v
}

func (o ConsulMeshGateway) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Mode != nil {
		toSerialize["Mode"] = o.Mode
	}
	return json.Marshal(toSerialize)
}

type NullableConsulMeshGateway struct {
	value *ConsulMeshGateway
	isSet bool
}

func (v NullableConsulMeshGateway) Get() *ConsulMeshGateway {
	return v.value
}

func (v *NullableConsulMeshGateway) Set(val *ConsulMeshGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableConsulMeshGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableConsulMeshGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsulMeshGateway(val *ConsulMeshGateway) *NullableConsulMeshGateway {
	return &NullableConsulMeshGateway{value: val, isSet: true}
}

func (v NullableConsulMeshGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsulMeshGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
