// Code generated by MockGen. DO NOT EDIT.
// Source: ./cache/registry.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	cache "github.com/catalystgo/cache-go/cache"
	gomock "github.com/golang/mock/gomock"
)

// MockRegistry is a mock of Registry interface.
type MockRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockRegistryMockRecorder
}

// MockRegistryMockRecorder is the mock recorder for MockRegistry.
type MockRegistryMockRecorder struct {
	mock *MockRegistry
}

// NewMockRegistry creates a new mock instance.
func NewMockRegistry(ctrl *gomock.Controller) *MockRegistry {
	mock := &MockRegistry{ctrl: ctrl}
	mock.recorder = &MockRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegistry) EXPECT() *MockRegistryMockRecorder {
	return m.recorder
}

// GetByName mocks base method.
func (m *MockRegistry) GetByName(name string) (cache.NamedCache, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", name)
	ret0, _ := ret[0].(cache.NamedCache)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName.
func (mr *MockRegistryMockRecorder) GetByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockRegistry)(nil).GetByName), name)
}

// MaybeRegister mocks base method.
func (m *MockRegistry) MaybeRegister(cache cache.NamedCache, err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "MaybeRegister", cache, err)
}

// MaybeRegister indicates an expected call of MaybeRegister.
func (mr *MockRegistryMockRecorder) MaybeRegister(cache, err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaybeRegister", reflect.TypeOf((*MockRegistry)(nil).MaybeRegister), cache, err)
}

// MustRegister mocks base method.
func (m *MockRegistry) MustRegister(cache cache.NamedCache, err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "MustRegister", cache, err)
}

// MustRegister indicates an expected call of MustRegister.
func (mr *MockRegistryMockRecorder) MustRegister(cache, err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MustRegister", reflect.TypeOf((*MockRegistry)(nil).MustRegister), cache, err)
}

// Register mocks base method.
func (m *MockRegistry) Register(caches ...cache.NamedCache) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range caches {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Register", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockRegistryMockRecorder) Register(caches ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockRegistry)(nil).Register), caches...)
}
