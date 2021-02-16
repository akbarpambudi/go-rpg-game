// Code generated by MockGen. DO NOT EDIT.
// Source: ../character.go

// Package repositorymock is a generated GoMock package.
package repositorymock

import (
	context "context"
	entity "github.com/akbarpambudi/go-rpg-game/internal/app/entity"
	predicate "github.com/akbarpambudi/go-rpg-game/internal/app/entity/predicate"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCharacter is a mock of Character interface
type MockCharacter struct {
	ctrl     *gomock.Controller
	recorder *MockCharacterMockRecorder
}

// MockCharacterMockRecorder is the mock recorder for MockCharacter
type MockCharacterMockRecorder struct {
	mock *MockCharacter
}

// NewMockCharacter creates a new mock instance
func NewMockCharacter(ctrl *gomock.Controller) *MockCharacter {
	mock := &MockCharacter{ctrl: ctrl}
	mock.recorder = &MockCharacterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCharacter) EXPECT() *MockCharacterMockRecorder {
	return m.recorder
}

// CreateOrUpdate mocks base method
func (m *MockCharacter) CreateOrUpdate(ctx context.Context, chara *entity.Character) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", ctx, chara)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate
func (mr *MockCharacterMockRecorder) CreateOrUpdate(ctx, chara interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockCharacter)(nil).CreateOrUpdate), ctx, chara)
}

// LoadByID mocks base method
func (m *MockCharacter) LoadByID(ctx context.Context, id uint) (*entity.Character, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadByID", ctx, id)
	ret0, _ := ret[0].(*entity.Character)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadByID indicates an expected call of LoadByID
func (mr *MockCharacterMockRecorder) LoadByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadByID", reflect.TypeOf((*MockCharacter)(nil).LoadByID), ctx, id)
}

// LoadMany mocks base method
func (m *MockCharacter) LoadMany(ctx context.Context, predicate predicate.Character) ([]*entity.Character, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadMany", ctx, predicate)
	ret0, _ := ret[0].([]*entity.Character)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadMany indicates an expected call of LoadMany
func (mr *MockCharacterMockRecorder) LoadMany(ctx, predicate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadMany", reflect.TypeOf((*MockCharacter)(nil).LoadMany), ctx, predicate)
}

// RemoveByID mocks base method
func (m *MockCharacter) RemoveByID(ctx context.Context, id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveByID", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveByID indicates an expected call of RemoveByID
func (mr *MockCharacterMockRecorder) RemoveByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveByID", reflect.TypeOf((*MockCharacter)(nil).RemoveByID), ctx, id)
}