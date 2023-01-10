// Code generated by MockGen. DO NOT EDIT.
// Source: example.com/sarang-apis/services (interfaces: UserService)

// Package services is a generated GoMock package.
package services

import (
	reflect "reflect"

	dto "example.com/sarang-apis/dto"
	models "example.com/sarang-apis/models"
	gomock "github.com/golang/mock/gomock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// UserDelete mocks base method.
func (m *MockUserService) UserDelete(arg0 primitive.ObjectID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserDelete", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserDelete indicates an expected call of UserDelete.
func (mr *MockUserServiceMockRecorder) UserDelete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserDelete", reflect.TypeOf((*MockUserService)(nil).UserDelete), arg0)
}

// UserGetAll mocks base method.
func (m *MockUserService) UserGetAll() ([]models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserGetAll")
	ret0, _ := ret[0].([]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserGetAll indicates an expected call of UserGetAll.
func (mr *MockUserServiceMockRecorder) UserGetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserGetAll", reflect.TypeOf((*MockUserService)(nil).UserGetAll))
}

// UserInsert mocks base method.
func (m *MockUserService) UserInsert(arg0 models.User) (*dto.UserDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserInsert", arg0)
	ret0, _ := ret[0].(*dto.UserDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserInsert indicates an expected call of UserInsert.
func (mr *MockUserServiceMockRecorder) UserInsert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserInsert", reflect.TypeOf((*MockUserService)(nil).UserInsert), arg0)
}

// UserUpdate mocks base method.
func (m *MockUserService) UserUpdate(arg0 primitive.ObjectID, arg1 primitive.M) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserUpdate", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserUpdate indicates an expected call of UserUpdate.
func (mr *MockUserServiceMockRecorder) UserUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserUpdate", reflect.TypeOf((*MockUserService)(nil).UserUpdate), arg0, arg1)
}