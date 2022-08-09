// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cq-provider-aws/client (interfaces: SecretsManagerClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	secretsmanager "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	gomock "github.com/golang/mock/gomock"
)

// MockSecretsManagerClient is a mock of SecretsManagerClient interface.
type MockSecretsManagerClient struct {
	ctrl     *gomock.Controller
	recorder *MockSecretsManagerClientMockRecorder
}

// MockSecretsManagerClientMockRecorder is the mock recorder for MockSecretsManagerClient.
type MockSecretsManagerClientMockRecorder struct {
	mock *MockSecretsManagerClient
}

// NewMockSecretsManagerClient creates a new mock instance.
func NewMockSecretsManagerClient(ctrl *gomock.Controller) *MockSecretsManagerClient {
	mock := &MockSecretsManagerClient{ctrl: ctrl}
	mock.recorder = &MockSecretsManagerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretsManagerClient) EXPECT() *MockSecretsManagerClientMockRecorder {
	return m.recorder
}

// DescribeSecret mocks base method.
func (m *MockSecretsManagerClient) DescribeSecret(arg0 context.Context, arg1 *secretsmanager.DescribeSecretInput, arg2 ...func(*secretsmanager.Options)) (*secretsmanager.DescribeSecretOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeSecret", varargs...)
	ret0, _ := ret[0].(*secretsmanager.DescribeSecretOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSecret indicates an expected call of DescribeSecret.
func (mr *MockSecretsManagerClientMockRecorder) DescribeSecret(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSecret", reflect.TypeOf((*MockSecretsManagerClient)(nil).DescribeSecret), varargs...)
}

// GetResourcePolicy mocks base method.
func (m *MockSecretsManagerClient) GetResourcePolicy(arg0 context.Context, arg1 *secretsmanager.GetResourcePolicyInput, arg2 ...func(*secretsmanager.Options)) (*secretsmanager.GetResourcePolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetResourcePolicy", varargs...)
	ret0, _ := ret[0].(*secretsmanager.GetResourcePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResourcePolicy indicates an expected call of GetResourcePolicy.
func (mr *MockSecretsManagerClientMockRecorder) GetResourcePolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourcePolicy", reflect.TypeOf((*MockSecretsManagerClient)(nil).GetResourcePolicy), varargs...)
}

// ListSecrets mocks base method.
func (m *MockSecretsManagerClient) ListSecrets(arg0 context.Context, arg1 *secretsmanager.ListSecretsInput, arg2 ...func(*secretsmanager.Options)) (*secretsmanager.ListSecretsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSecrets", varargs...)
	ret0, _ := ret[0].(*secretsmanager.ListSecretsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecrets indicates an expected call of ListSecrets.
func (mr *MockSecretsManagerClientMockRecorder) ListSecrets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecrets", reflect.TypeOf((*MockSecretsManagerClient)(nil).ListSecrets), varargs...)
}