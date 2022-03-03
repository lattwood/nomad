package agent

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/go-memdb"
	"github.com/hashicorp/nomad/nomad/mock"
	"github.com/hashicorp/nomad/nomad/structs"
	"github.com/stretchr/testify/require"
)

func TestHTTPServer_ServiceRegistrationListRequest(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		testFn func(srv *TestAgent)
		name   string
	}{
		{
			testFn: func(s *TestAgent) {

				// Grab the state, so we can manipulate it and test against it.
				testState := s.Agent.server.State()

				// Generate service registrations and upsert.
				serviceRegs := mock.ServiceRegistrations()
				require.NoError(t, testState.UpsertServiceRegistrations(
					structs.MsgTypeTestSetup, 10, serviceRegs))

				// Build the HTTP request.
				req, err := http.NewRequest(http.MethodGet, "/v1/services", nil)
				require.NoError(t, err)
				respW := httptest.NewRecorder()

				// Send the HTTP request.
				obj, err := s.Server.ServiceRegistrationListRequest(respW, req)
				require.NoError(t, err)
				require.NotNil(t, obj)

				// Check the index is not zero.
				require.EqualValues(t, "10", respW.Header().Get("X-Nomad-Index"))
				require.ElementsMatch(t, []*structs.ServiceRegistrationListStub{
					{
						Namespace: "default",
						Services: []*structs.ServiceRegistrationStub{
							{
								ServiceName: "example-cache",
								Tags:        []string{"foo"},
							},
						},
					},
				}, obj.([]*structs.ServiceRegistrationListStub))
			},
			name: "list default namespace",
		},
		{
			testFn: func(s *TestAgent) {

				// Grab the state, so we can manipulate it and test against it.
				testState := s.Agent.server.State()

				// Generate service registrations and upsert.
				serviceRegs := mock.ServiceRegistrations()
				require.NoError(t, testState.UpsertServiceRegistrations(
					structs.MsgTypeTestSetup, 10, serviceRegs))

				// Build the HTTP request.
				req, err := http.NewRequest(http.MethodGet, "/v1/services?namespace=platform", nil)
				require.NoError(t, err)
				respW := httptest.NewRecorder()

				// Send the HTTP request.
				obj, err := s.Server.ServiceRegistrationListRequest(respW, req)
				require.NoError(t, err)
				require.NotNil(t, obj)

				// Check the index is not zero.
				require.EqualValues(t, "10", respW.Header().Get("X-Nomad-Index"))
				require.ElementsMatch(t, []*structs.ServiceRegistrationListStub{
					{
						Namespace: "platform",
						Services: []*structs.ServiceRegistrationStub{
							{
								ServiceName: "countdash-api",
								Tags:        []string{"bar"},
							},
						},
					},
				}, obj.([]*structs.ServiceRegistrationListStub))
			},
			name: "list platform namespace",
		},
		{
			testFn: func(s *TestAgent) {

				// Grab the state, so we can manipulate it and test against it.
				testState := s.Agent.server.State()

				// Generate service registrations and upsert.
				serviceRegs := mock.ServiceRegistrations()
				require.NoError(t, testState.UpsertServiceRegistrations(
					structs.MsgTypeTestSetup, 10, serviceRegs))

				// Build the HTTP request.
				req, err := http.NewRequest(http.MethodGet, "/v1/services?namespace=*", nil)
				require.NoError(t, err)
				respW := httptest.NewRecorder()

				// Send the HTTP request.
				obj, err := s.Server.ServiceRegistrationListRequest(respW, req)
				require.NoError(t, err)
				require.NotNil(t, obj)

				// Check the index is not zero.
				require.EqualValues(t, "10", respW.Header().Get("X-Nomad-Index"))
				require.ElementsMatch(t, []*structs.ServiceRegistrationListStub{
					{
						Namespace: "default",
						Services: []*structs.ServiceRegistrationStub{
							{
								ServiceName: "example-cache",
								Tags:        []string{"foo"},
							},
						},
					},
					{
						Namespace: "platform",
						Services: []*structs.ServiceRegistrationStub{
							{
								ServiceName: "countdash-api",
								Tags:        []string{"bar"},
							},
						},
					},
				}, obj.([]*structs.ServiceRegistrationListStub))
			},
			name: "list wildcard namespace",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			httpTest(t, nil, tc.testFn)
		})
	}
}

func TestHTTPServer_ServiceRegistrationRequest(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		testFn func(srv *TestAgent)
		name   string
	}{
		{
			testFn: func(s *TestAgent) {

				// Grab the state, so we can manipulate it and test against it.
				testState := s.Agent.server.State()

				// Generate a service registration and upsert this.
				serviceReg := mock.ServiceRegistrations()[0]
				require.NoError(t, testState.UpsertServiceRegistrations(
					structs.MsgTypeTestSetup, 10, []*structs.ServiceRegistration{serviceReg}))

				// Build the HTTP request.
				path := fmt.Sprintf("/v1/service/%s/%s", serviceReg.ServiceName, serviceReg.ID)
				req, err := http.NewRequest(http.MethodDelete, path, nil)
				require.NoError(t, err)
				respW := httptest.NewRecorder()

				// Send the HTTP request.
				obj, err := s.Server.ServiceRegistrationRequest(respW, req)
				require.NoError(t, err)
				require.Nil(t, obj)

				// Check the index is not zero.
				require.NotZero(t, respW.Header().Get("X-Nomad-Index"))

				// Check that the service is not found within state.
				out, err := testState.GetServiceRegistrationByID(memdb.NewWatchSet(), serviceReg.Namespace, serviceReg.ID)
				require.Nil(t, out)
				require.NoError(t, err)
			},
			name: "delete by ID",
		},
		{
			testFn: func(s *TestAgent) {

				// Grab the state, so we can manipulate it and test against it.
				testState := s.Agent.server.State()

				// Generate a service registration and upsert this.
				serviceReg := mock.ServiceRegistrations()[0]
				require.NoError(t, testState.UpsertServiceRegistrations(
					structs.MsgTypeTestSetup, 10, []*structs.ServiceRegistration{serviceReg}))

				// Build the HTTP request.
				path := fmt.Sprintf("/v1/service/%s", serviceReg.ServiceName)
				req, err := http.NewRequest(http.MethodGet, path, nil)
				require.NoError(t, err)
				respW := httptest.NewRecorder()

				// Send the HTTP request.
				obj, err := s.Server.ServiceRegistrationRequest(respW, req)
				require.NoError(t, err)

				// Check the index is not zero and that we see the service
				// registration.
				require.NotZero(t, respW.Header().Get("X-Nomad-Index"))
				require.Equal(t, serviceReg, obj.([]*structs.ServiceRegistration)[0])
			},
			name: "get service by name",
		},
		{
			testFn: func(s *TestAgent) {

				// Build the HTTP request.
				req, err := http.NewRequest(http.MethodGet, "/v1/service/foo/bar/baz/bonkers", nil)
				require.NoError(t, err)
				respW := httptest.NewRecorder()

				// Send the HTTP request.
				obj, err := s.Server.ServiceRegistrationRequest(respW, req)
				require.Error(t, err)
				require.Contains(t, err.Error(), "invalid URI")
				require.Nil(t, obj)
			},
			name: "incorrect URI format",
		},
		{
			testFn: func(s *TestAgent) {

				// Build the HTTP request.
				req, err := http.NewRequest(http.MethodGet, "/v1/service/", nil)
				require.NoError(t, err)
				respW := httptest.NewRecorder()

				// Send the HTTP request.
				obj, err := s.Server.ServiceRegistrationRequest(respW, req)
				require.Error(t, err)
				require.Contains(t, err.Error(), "missing service name")
				require.Nil(t, obj)
			},
			name: "get service empty name",
		},
		{
			testFn: func(s *TestAgent) {

				// Build the HTTP request.
				req, err := http.NewRequest(http.MethodHead, "/v1/service/foo", nil)
				require.NoError(t, err)
				respW := httptest.NewRecorder()

				// Send the HTTP request.
				obj, err := s.Server.ServiceRegistrationRequest(respW, req)
				require.Error(t, err)
				require.Contains(t, err.Error(), "Invalid method")
				require.Nil(t, obj)
			},
			name: "get service incorrect method",
		},
		{
			testFn: func(s *TestAgent) {

				// Build the HTTP request.
				req, err := http.NewRequest(http.MethodDelete, "/v1/service/foo/", nil)
				require.NoError(t, err)
				respW := httptest.NewRecorder()

				// Send the HTTP request.
				obj, err := s.Server.ServiceRegistrationRequest(respW, req)
				require.Error(t, err)
				require.Contains(t, err.Error(), "missing service id")
				require.Nil(t, obj)
			},
			name: "delete service empty id",
		},
		{
			testFn: func(s *TestAgent) {

				// Build the HTTP request.
				req, err := http.NewRequest(http.MethodHead, "/v1/service/foo/bar", nil)
				require.NoError(t, err)
				respW := httptest.NewRecorder()

				// Send the HTTP request.
				obj, err := s.Server.ServiceRegistrationRequest(respW, req)
				require.Error(t, err)
				require.Contains(t, err.Error(), "Invalid method")
				require.Nil(t, obj)
			},
			name: "delete service incorrect method",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			httpTest(t, nil, tc.testFn)
		})
	}
}
