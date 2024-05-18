// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/agusheryanto182/go-health-record/utils/hash (interfaces: HashInterface)

// Package hash is a generated GoMock package.
package hash

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHashInterface is a mock of HashInterface interface.
type MockHashInterface struct {
	ctrl     *gomock.Controller
	recorder *MockHashInterfaceMockRecorder
}

// MockHashInterfaceMockRecorder is the mock recorder for MockHashInterface.
type MockHashInterfaceMockRecorder struct {
	mock *MockHashInterface
}

// NewMockHashInterface creates a new mock instance.
func NewMockHashInterface(ctrl *gomock.Controller) *MockHashInterface {
	mock := &MockHashInterface{ctrl: ctrl}
	mock.recorder = &MockHashInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHashInterface) EXPECT() *MockHashInterfaceMockRecorder {
	return m.recorder
}

// CheckPasswordHash mocks base method.
func (m *MockHashInterface) CheckPasswordHash(arg0, arg1 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckPasswordHash", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckPasswordHash indicates an expected call of CheckPasswordHash.
func (mr *MockHashInterfaceMockRecorder) CheckPasswordHash(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckPasswordHash", reflect.TypeOf((*MockHashInterface)(nil).CheckPasswordHash), arg0, arg1)
}

// HashPassword mocks base method.
func (m *MockHashInterface) HashPassword(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashPassword", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HashPassword indicates an expected call of HashPassword.
func (mr *MockHashInterfaceMockRecorder) HashPassword(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashPassword", reflect.TypeOf((*MockHashInterface)(nil).HashPassword), arg0)
}