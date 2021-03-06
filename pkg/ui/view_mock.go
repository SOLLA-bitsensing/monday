// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/ui/view.go

// Package ui is a generated GoMock package.
package ui

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockView is a mock of View interface.
type MockView struct {
	ctrl     *gomock.Controller
	recorder *MockViewMockRecorder
}

// MockViewMockRecorder is the mock recorder for MockView.
type MockViewMockRecorder struct {
	mock *MockView
}

// NewMockView creates a new mock instance.
func NewMockView(ctrl *gomock.Controller) *MockView {
	mock := &MockView{ctrl: ctrl}
	mock.recorder = &MockViewMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockView) EXPECT() *MockViewMockRecorder {
	return m.recorder
}

// GetName mocks base method.
func (m *MockView) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockViewMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockView)(nil).GetName))
}

// Write mocks base method.
func (m *MockView) Write(str string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Write", str)
}

// Write indicates an expected call of Write.
func (mr *MockViewMockRecorder) Write(str interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockView)(nil).Write), str)
}

// Writef mocks base method.
func (m *MockView) Writef(str string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{str}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Writef", varargs...)
}

// Writef indicates an expected call of Writef.
func (mr *MockViewMockRecorder) Writef(str interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{str}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Writef", reflect.TypeOf((*MockView)(nil).Writef), varargs...)
}
