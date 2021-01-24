// Code generated by MockGen. DO NOT EDIT.
// Source: queries/userQueryService/usecase.go

// Package userqueryservice is a generated GoMock package.
package userqueryservice

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUsecase is a mock of Usecase interface
type MockUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseMockRecorder
}

// MockUsecaseMockRecorder is the mock recorder for MockUsecase
type MockUsecaseMockRecorder struct {
	mock *MockUsecase
}

// NewMockUsecase creates a new mock instance
func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
	mock := &MockUsecase{ctrl: ctrl}
	mock.recorder = &MockUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUsecase) EXPECT() *MockUsecaseMockRecorder {
	return m.recorder
}

// getByID mocks base method
func (m *MockUsecase) getByID(req getByIDRequest) (getByIDResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getByID", req)
	ret0, _ := ret[0].(getByIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getByID indicates an expected call of getByID
func (mr *MockUsecaseMockRecorder) getByID(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getByID", reflect.TypeOf((*MockUsecase)(nil).getByID), req)
}

// getByEmailAndPassword mocks base method
func (m *MockUsecase) getByEmailAndPassword(req getByEmailAndPasswordRequest) (getByEmailAndPasswordResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getByEmailAndPassword", req)
	ret0, _ := ret[0].(getByEmailAndPasswordResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getByEmailAndPassword indicates an expected call of getByEmailAndPassword
func (mr *MockUsecaseMockRecorder) getByEmailAndPassword(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getByEmailAndPassword", reflect.TypeOf((*MockUsecase)(nil).getByEmailAndPassword), req)
}

// MockDataAccessor is a mock of DataAccessor interface
type MockDataAccessor struct {
	ctrl     *gomock.Controller
	recorder *MockDataAccessorMockRecorder
}

// MockDataAccessorMockRecorder is the mock recorder for MockDataAccessor
type MockDataAccessorMockRecorder struct {
	mock *MockDataAccessor
}

// NewMockDataAccessor creates a new mock instance
func NewMockDataAccessor(ctrl *gomock.Controller) *MockDataAccessor {
	mock := &MockDataAccessor{ctrl: ctrl}
	mock.recorder = &MockDataAccessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDataAccessor) EXPECT() *MockDataAccessorMockRecorder {
	return m.recorder
}

// getByID mocks base method
func (m *MockDataAccessor) getByID(req getByIDRequest) (getByIDResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getByID", req)
	ret0, _ := ret[0].(getByIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getByID indicates an expected call of getByID
func (mr *MockDataAccessorMockRecorder) getByID(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getByID", reflect.TypeOf((*MockDataAccessor)(nil).getByID), req)
}

// getByEmail mocks base method
func (m *MockDataAccessor) getByEmail(req getByEmailAndPasswordRequest) (getByEmailAndPasswordResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getByEmail", req)
	ret0, _ := ret[0].(getByEmailAndPasswordResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getByEmail indicates an expected call of getByEmail
func (mr *MockDataAccessorMockRecorder) getByEmail(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getByEmail", reflect.TypeOf((*MockDataAccessor)(nil).getByEmail), req)
}
