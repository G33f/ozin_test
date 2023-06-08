// Code generated by MockGen. DO NOT EDIT.
// Source: internal/shortener/repo.go

// Package mock_shortener is a generated GoMock package.
package mock_shortener

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// AddShortURL mocks base method.
func (m *MockRepo) AddShortURL(ctx context.Context, url, shortURL string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddShortURL", ctx, url, shortURL)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddShortURL indicates an expected call of AddShortURL.
func (mr *MockRepoMockRecorder) AddShortURL(ctx, url, shortURL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddShortURL", reflect.TypeOf((*MockRepo)(nil).AddShortURL), ctx, url, shortURL)
}

// GetURL mocks base method.
func (m *MockRepo) GetURL(ctx context.Context, shortURL string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetURL", ctx, shortURL)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetURL indicates an expected call of GetURL.
func (mr *MockRepoMockRecorder) GetURL(ctx, shortURL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetURL", reflect.TypeOf((*MockRepo)(nil).GetURL), ctx, shortURL)
}
