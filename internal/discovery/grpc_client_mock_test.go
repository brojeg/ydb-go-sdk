// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ydb-platform/ydb-go-genproto/Ydb_Discovery_V1 (interfaces: DiscoveryServiceClient)
//
// Generated by this command:
//
//	mockgen -destination grpc_client_mock_test.go -package discovery -write_package_comment=false github.com/ydb-platform/ydb-go-genproto/Ydb_Discovery_V1 DiscoveryServiceClient
package discovery

import (
	context "context"
	reflect "reflect"

	Ydb_Discovery "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Discovery"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockDiscoveryServiceClient is a mock of DiscoveryServiceClient interface.
type MockDiscoveryServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockDiscoveryServiceClientMockRecorder
}

// MockDiscoveryServiceClientMockRecorder is the mock recorder for MockDiscoveryServiceClient.
type MockDiscoveryServiceClientMockRecorder struct {
	mock *MockDiscoveryServiceClient
}

// NewMockDiscoveryServiceClient creates a new mock instance.
func NewMockDiscoveryServiceClient(ctrl *gomock.Controller) *MockDiscoveryServiceClient {
	mock := &MockDiscoveryServiceClient{ctrl: ctrl}
	mock.recorder = &MockDiscoveryServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDiscoveryServiceClient) EXPECT() *MockDiscoveryServiceClientMockRecorder {
	return m.recorder
}

// ListEndpoints mocks base method.
func (m *MockDiscoveryServiceClient) ListEndpoints(arg0 context.Context, arg1 *Ydb_Discovery.ListEndpointsRequest, arg2 ...grpc.CallOption) (*Ydb_Discovery.ListEndpointsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListEndpoints", varargs...)
	ret0, _ := ret[0].(*Ydb_Discovery.ListEndpointsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEndpoints indicates an expected call of ListEndpoints.
func (mr *MockDiscoveryServiceClientMockRecorder) ListEndpoints(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEndpoints", reflect.TypeOf((*MockDiscoveryServiceClient)(nil).ListEndpoints), varargs...)
}

// WhoAmI mocks base method.
func (m *MockDiscoveryServiceClient) WhoAmI(arg0 context.Context, arg1 *Ydb_Discovery.WhoAmIRequest, arg2 ...grpc.CallOption) (*Ydb_Discovery.WhoAmIResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WhoAmI", varargs...)
	ret0, _ := ret[0].(*Ydb_Discovery.WhoAmIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WhoAmI indicates an expected call of WhoAmI.
func (mr *MockDiscoveryServiceClientMockRecorder) WhoAmI(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WhoAmI", reflect.TypeOf((*MockDiscoveryServiceClient)(nil).WhoAmI), varargs...)
}
