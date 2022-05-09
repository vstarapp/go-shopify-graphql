// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vstarapp/go-shopify-graphql/v6 (interfaces: ProductService)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/r0busta/go-shopify-graphql-model/v3/graph/model"
)

// MockProductService is a mock of ProductService interface.
type MockProductService struct {
	ctrl     *gomock.Controller
	recorder *MockProductServiceMockRecorder
}

// MockProductServiceMockRecorder is the mock recorder for MockProductService.
type MockProductServiceMockRecorder struct {
	mock *MockProductService
}

// NewMockProductService creates a new mock instance.
func NewMockProductService(ctrl *gomock.Controller) *MockProductService {
	mock := &MockProductService{ctrl: ctrl}
	mock.recorder = &MockProductServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductService) EXPECT() *MockProductServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProductService) Create(arg0 model.ProductInput) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockProductServiceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProductService)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockProductService) Delete(arg0 model.ProductDeleteInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockProductServiceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProductService)(nil).Delete), arg0)
}

// Get mocks base method.
func (m *MockProductService) Get(arg0 string) (*model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockProductServiceMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockProductService)(nil).Get), arg0)
}

// List mocks base method.
func (m *MockProductService) List(arg0 string) ([]model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].([]model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockProductServiceMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockProductService)(nil).List), arg0)
}

// ListAll mocks base method.
func (m *MockProductService) ListAll() ([]model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll")
	ret0, _ := ret[0].([]model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAll indicates an expected call of ListAll.
func (mr *MockProductServiceMockRecorder) ListAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockProductService)(nil).ListAll))
}

// Update mocks base method.
func (m *MockProductService) Update(arg0 model.ProductInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockProductServiceMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockProductService)(nil).Update), arg0)
}

// VariantsBulkCreate mocks base method.
func (m *MockProductService) VariantsBulkCreate(arg0 string, arg1 []model.ProductVariantsBulkInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VariantsBulkCreate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// VariantsBulkCreate indicates an expected call of VariantsBulkCreate.
func (mr *MockProductServiceMockRecorder) VariantsBulkCreate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VariantsBulkCreate", reflect.TypeOf((*MockProductService)(nil).VariantsBulkCreate), arg0, arg1)
}

// VariantsBulkReorder mocks base method.
func (m *MockProductService) VariantsBulkReorder(arg0 string, arg1 []model.ProductVariantPositionInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VariantsBulkReorder", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// VariantsBulkReorder indicates an expected call of VariantsBulkReorder.
func (mr *MockProductServiceMockRecorder) VariantsBulkReorder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VariantsBulkReorder", reflect.TypeOf((*MockProductService)(nil).VariantsBulkReorder), arg0, arg1)
}

// VariantsBulkUpdate mocks base method.
func (m *MockProductService) VariantsBulkUpdate(arg0 string, arg1 []model.ProductVariantsBulkInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VariantsBulkUpdate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// VariantsBulkUpdate indicates an expected call of VariantsBulkUpdate.
func (mr *MockProductServiceMockRecorder) VariantsBulkUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VariantsBulkUpdate", reflect.TypeOf((*MockProductService)(nil).VariantsBulkUpdate), arg0, arg1)
}
