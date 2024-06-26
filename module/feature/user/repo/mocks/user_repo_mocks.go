// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/agusheryanto182/go-health-record/module/feature/user (interfaces: UserRepoInterface)

// Package repo is a generated GoMock package.
package repo

import (
	reflect "reflect"

	entities "github.com/agusheryanto182/go-health-record/module/entities"
	dto "github.com/agusheryanto182/go-health-record/module/feature/user/dto"
	gomock "github.com/golang/mock/gomock"
)

// MockUserRepoInterface is a mock of UserRepoInterface interface.
type MockUserRepoInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepoInterfaceMockRecorder
}

// MockUserRepoInterfaceMockRecorder is the mock recorder for MockUserRepoInterface.
type MockUserRepoInterfaceMockRecorder struct {
	mock *MockUserRepoInterface
}

// NewMockUserRepoInterface creates a new mock instance.
func NewMockUserRepoInterface(ctrl *gomock.Controller) *MockUserRepoInterface {
	mock := &MockUserRepoInterface{ctrl: ctrl}
	mock.recorder = &MockUserRepoInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepoInterface) EXPECT() *MockUserRepoInterfaceMockRecorder {
	return m.recorder
}

// CheckUserByIdAndRole mocks base method.
func (m *MockUserRepoInterface) CheckUserByIdAndRole(arg0, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserByIdAndRole", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUserByIdAndRole indicates an expected call of CheckUserByIdAndRole.
func (mr *MockUserRepoInterfaceMockRecorder) CheckUserByIdAndRole(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserByIdAndRole", reflect.TypeOf((*MockUserRepoInterface)(nil).CheckUserByIdAndRole), arg0, arg1)
}

// DeleteUserNurse mocks base method.
func (m *MockUserRepoInterface) DeleteUserNurse(arg0 *dto.DeleteUserNurse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserNurse", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserNurse indicates an expected call of DeleteUserNurse.
func (mr *MockUserRepoInterfaceMockRecorder) DeleteUserNurse(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserNurse", reflect.TypeOf((*MockUserRepoInterface)(nil).DeleteUserNurse), arg0)
}

// GetUser mocks base method.
func (m *MockUserRepoInterface) GetUser(arg0 int64, arg1 string) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserRepoInterfaceMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserRepoInterface)(nil).GetUser), arg0, arg1)
}

// GetUserByFilters mocks base method.
func (m *MockUserRepoInterface) GetUserByFilters(arg0 *dto.UserFilter) ([]*dto.UserFilterResponses, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByFilters", arg0)
	ret0, _ := ret[0].([]*dto.UserFilterResponses)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByFilters indicates an expected call of GetUserByFilters.
func (mr *MockUserRepoInterfaceMockRecorder) GetUserByFilters(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByFilters", reflect.TypeOf((*MockUserRepoInterface)(nil).GetUserByFilters), arg0)
}

// GetUserByID mocks base method.
func (m *MockUserRepoInterface) GetUserByID(arg0 string) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", arg0)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockUserRepoInterfaceMockRecorder) GetUserByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockUserRepoInterface)(nil).GetUserByID), arg0)
}

// IsNipExist mocks base method.
func (m *MockUserRepoInterface) IsNipExist(arg0 int64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsNipExist", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsNipExist indicates an expected call of IsNipExist.
func (mr *MockUserRepoInterfaceMockRecorder) IsNipExist(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNipExist", reflect.TypeOf((*MockUserRepoInterface)(nil).IsNipExist), arg0)
}

// Register mocks base method.
func (m *MockUserRepoInterface) Register(arg0 *entities.User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockUserRepoInterfaceMockRecorder) Register(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockUserRepoInterface)(nil).Register), arg0)
}

// SetPasswordNurse mocks base method.
func (m *MockUserRepoInterface) SetPasswordNurse(arg0 *dto.SetPasswordNurse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPasswordNurse", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPasswordNurse indicates an expected call of SetPasswordNurse.
func (mr *MockUserRepoInterfaceMockRecorder) SetPasswordNurse(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPasswordNurse", reflect.TypeOf((*MockUserRepoInterface)(nil).SetPasswordNurse), arg0)
}

// UpdateUserNurse mocks base method.
func (m *MockUserRepoInterface) UpdateUserNurse(arg0 *dto.UpdateUserNurse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserNurse", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserNurse indicates an expected call of UpdateUserNurse.
func (mr *MockUserRepoInterfaceMockRecorder) UpdateUserNurse(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserNurse", reflect.TypeOf((*MockUserRepoInterface)(nil).UpdateUserNurse), arg0)
}
