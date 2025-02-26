// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/github/client (interfaces: IssuesService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	github "github.com/google/go-github/v45/github"
)

// MockIssuesService is a mock of IssuesService interface.
type MockIssuesService struct {
	ctrl     *gomock.Controller
	recorder *MockIssuesServiceMockRecorder
}

// MockIssuesServiceMockRecorder is the mock recorder for MockIssuesService.
type MockIssuesServiceMockRecorder struct {
	mock *MockIssuesService
}

// NewMockIssuesService creates a new mock instance.
func NewMockIssuesService(ctrl *gomock.Controller) *MockIssuesService {
	mock := &MockIssuesService{ctrl: ctrl}
	mock.recorder = &MockIssuesServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIssuesService) EXPECT() *MockIssuesServiceMockRecorder {
	return m.recorder
}

// ListByOrg mocks base method.
func (m *MockIssuesService) ListByOrg(arg0 context.Context, arg1 string, arg2 *github.IssueListOptions) ([]*github.Issue, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByOrg", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*github.Issue)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListByOrg indicates an expected call of ListByOrg.
func (mr *MockIssuesServiceMockRecorder) ListByOrg(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByOrg", reflect.TypeOf((*MockIssuesService)(nil).ListByOrg), arg0, arg1, arg2)
}
