// Code generated by MockGen. DO NOT EDIT.
// Source: core/model.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	sql "database/sql"
	entity "github.com/codefluence-x/altair/entity"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockHasName is a mock of HasName interface
type MockHasName struct {
	ctrl     *gomock.Controller
	recorder *MockHasNameMockRecorder
}

// MockHasNameMockRecorder is the mock recorder for MockHasName
type MockHasNameMockRecorder struct {
	mock *MockHasName
}

// NewMockHasName creates a new mock instance
func NewMockHasName(ctrl *gomock.Controller) *MockHasName {
	mock := &MockHasName{ctrl: ctrl}
	mock.recorder = &MockHasNameMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHasName) EXPECT() *MockHasNameMockRecorder {
	return m.recorder
}

// Name mocks base method
func (m *MockHasName) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockHasNameMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockHasName)(nil).Name))
}

// MockDBExecutable is a mock of DBExecutable interface
type MockDBExecutable struct {
	ctrl     *gomock.Controller
	recorder *MockDBExecutableMockRecorder
}

// MockDBExecutableMockRecorder is the mock recorder for MockDBExecutable
type MockDBExecutableMockRecorder struct {
	mock *MockDBExecutable
}

// NewMockDBExecutable creates a new mock instance
func NewMockDBExecutable(ctrl *gomock.Controller) *MockDBExecutable {
	mock := &MockDBExecutable{ctrl: ctrl}
	mock.recorder = &MockDBExecutableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDBExecutable) EXPECT() *MockDBExecutableMockRecorder {
	return m.recorder
}

// Exec mocks base method
func (m *MockDBExecutable) Exec(query string, args ...interface{}) (sql.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec
func (mr *MockDBExecutableMockRecorder) Exec(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockDBExecutable)(nil).Exec), varargs...)
}

// MockOauthApplicationModel is a mock of OauthApplicationModel interface
type MockOauthApplicationModel struct {
	ctrl     *gomock.Controller
	recorder *MockOauthApplicationModelMockRecorder
}

// MockOauthApplicationModelMockRecorder is the mock recorder for MockOauthApplicationModel
type MockOauthApplicationModelMockRecorder struct {
	mock *MockOauthApplicationModel
}

// NewMockOauthApplicationModel creates a new mock instance
func NewMockOauthApplicationModel(ctrl *gomock.Controller) *MockOauthApplicationModel {
	mock := &MockOauthApplicationModel{ctrl: ctrl}
	mock.recorder = &MockOauthApplicationModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOauthApplicationModel) EXPECT() *MockOauthApplicationModelMockRecorder {
	return m.recorder
}

// Name mocks base method
func (m *MockOauthApplicationModel) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockOauthApplicationModelMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockOauthApplicationModel)(nil).Name))
}

// Paginate mocks base method
func (m *MockOauthApplicationModel) Paginate(ctx context.Context, offset, limit int) ([]entity.OauthApplication, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Paginate", ctx, offset, limit)
	ret0, _ := ret[0].([]entity.OauthApplication)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Paginate indicates an expected call of Paginate
func (mr *MockOauthApplicationModelMockRecorder) Paginate(ctx, offset, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Paginate", reflect.TypeOf((*MockOauthApplicationModel)(nil).Paginate), ctx, offset, limit)
}

// Count mocks base method
func (m *MockOauthApplicationModel) Count(ctx context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count
func (mr *MockOauthApplicationModelMockRecorder) Count(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockOauthApplicationModel)(nil).Count), ctx)
}

// Create mocks base method
func (m *MockOauthApplicationModel) Create(ctx context.Context, data entity.OauthApplicationJSON, txs ...*sql.Tx) (int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, data}
	for _, a := range txs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockOauthApplicationModelMockRecorder) Create(ctx, data interface{}, txs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, data}, txs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOauthApplicationModel)(nil).Create), varargs...)
}