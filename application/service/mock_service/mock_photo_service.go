// Code generated by MockGen. DO NOT EDIT.
// Source: application/service/photo_service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	gomock "github.com/golang/mock/gomock"
	model "github.com/photoshelf/photoshelf-storage/model"
	reflect "reflect"
)

// MockPhotoService is a mock of PhotoService interface
type MockPhotoService struct {
	ctrl     *gomock.Controller
	recorder *MockPhotoServiceMockRecorder
}

// MockPhotoServiceMockRecorder is the mock recorder for MockPhotoService
type MockPhotoServiceMockRecorder struct {
	mock *MockPhotoService
}

// NewMockPhotoService creates a new mock instance
func NewMockPhotoService(ctrl *gomock.Controller) *MockPhotoService {
	mock := &MockPhotoService{ctrl: ctrl}
	mock.recorder = &MockPhotoServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPhotoService) EXPECT() *MockPhotoServiceMockRecorder {
	return m.recorder
}

// Save mocks base method
func (m *MockPhotoService) Save(photo model.Photo) (*model.Identifier, error) {
	ret := m.ctrl.Call(m, "Save", photo)
	ret0, _ := ret[0].(*model.Identifier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save
func (mr *MockPhotoServiceMockRecorder) Save(photo interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockPhotoService)(nil).Save), photo)
}

// Find mocks base method
func (m *MockPhotoService) Find(id model.Identifier) (*model.Photo, error) {
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*model.Photo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockPhotoServiceMockRecorder) Find(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockPhotoService)(nil).Find), id)
}

// Delete mocks base method
func (m *MockPhotoService) Delete(id model.Identifier) error {
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockPhotoServiceMockRecorder) Delete(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPhotoService)(nil).Delete), id)
}
