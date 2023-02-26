// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/MeganViga/Plipkart/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	db "github.com/MeganViga/Plipkart/db/sqlc"
	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// AddProduct mocks base method.
func (m *MockStore) AddProduct(arg0 context.Context, arg1 db.AddProductParams) (db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProduct", arg0, arg1)
	ret0, _ := ret[0].(db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddProduct indicates an expected call of AddProduct.
func (mr *MockStoreMockRecorder) AddProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProduct", reflect.TypeOf((*MockStore)(nil).AddProduct), arg0, arg1)
}

// AddProductCategory mocks base method.
func (m *MockStore) AddProductCategory(arg0 context.Context, arg1 db.AddProductCategoryParams) (db.ProductCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProductCategory", arg0, arg1)
	ret0, _ := ret[0].(db.ProductCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddProductCategory indicates an expected call of AddProductCategory.
func (mr *MockStoreMockRecorder) AddProductCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProductCategory", reflect.TypeOf((*MockStore)(nil).AddProductCategory), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.UserDatum, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.UserDatum)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// CreateUserAddress mocks base method.
func (m *MockStore) CreateUserAddress(arg0 context.Context, arg1 db.CreateUserAddressParams) (db.UserAddress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserAddress", arg0, arg1)
	ret0, _ := ret[0].(db.UserAddress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUserAddress indicates an expected call of CreateUserAddress.
func (mr *MockStoreMockRecorder) CreateUserAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserAddress", reflect.TypeOf((*MockStore)(nil).CreateUserAddress), arg0, arg1)
}

// DeleteUserAddress mocks base method.
func (m *MockStore) DeleteUserAddress(arg0 context.Context, arg1 int64) (db.UserAddress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserAddress", arg0, arg1)
	ret0, _ := ret[0].(db.UserAddress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUserAddress indicates an expected call of DeleteUserAddress.
func (mr *MockStoreMockRecorder) DeleteUserAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserAddress", reflect.TypeOf((*MockStore)(nil).DeleteUserAddress), arg0, arg1)
}

// GetProductByColor mocks base method.
func (m *MockStore) GetProductByColor(arg0 context.Context, arg1 db.GetProductByColorParams) ([]db.GetProductByColorRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductByColor", arg0, arg1)
	ret0, _ := ret[0].([]db.GetProductByColorRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductByColor indicates an expected call of GetProductByColor.
func (mr *MockStoreMockRecorder) GetProductByColor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductByColor", reflect.TypeOf((*MockStore)(nil).GetProductByColor), arg0, arg1)
}

// GetProductById mocks base method.
func (m *MockStore) GetProductById(arg0 context.Context, arg1 int64) (db.GetProductByIdRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductById", arg0, arg1)
	ret0, _ := ret[0].(db.GetProductByIdRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductById indicates an expected call of GetProductById.
func (mr *MockStoreMockRecorder) GetProductById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductById", reflect.TypeOf((*MockStore)(nil).GetProductById), arg0, arg1)
}

// GetProductByName mocks base method.
func (m *MockStore) GetProductByName(arg0 context.Context, arg1 string) ([]db.GetProductByNameRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductByName", arg0, arg1)
	ret0, _ := ret[0].([]db.GetProductByNameRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductByName indicates an expected call of GetProductByName.
func (mr *MockStoreMockRecorder) GetProductByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductByName", reflect.TypeOf((*MockStore)(nil).GetProductByName), arg0, arg1)
}

// GetProducts mocks base method.
func (m *MockStore) GetProducts(arg0 context.Context) ([]db.GetProductsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProducts", arg0)
	ret0, _ := ret[0].([]db.GetProductsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProducts indicates an expected call of GetProducts.
func (mr *MockStoreMockRecorder) GetProducts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProducts", reflect.TypeOf((*MockStore)(nil).GetProducts), arg0)
}

// GetUserByName mocks base method.
func (m *MockStore) GetUserByName(arg0 context.Context, arg1 string) (db.UserDatum, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByName", arg0, arg1)
	ret0, _ := ret[0].(db.UserDatum)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByName indicates an expected call of GetUserByName.
func (mr *MockStoreMockRecorder) GetUserByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByName", reflect.TypeOf((*MockStore)(nil).GetUserByName), arg0, arg1)
}

// ListUserAddresses mocks base method.
func (m *MockStore) ListUserAddresses(arg0 context.Context, arg1 int64) ([]db.UserAddress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUserAddresses", arg0, arg1)
	ret0, _ := ret[0].([]db.UserAddress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUserAddresses indicates an expected call of ListUserAddresses.
func (mr *MockStoreMockRecorder) ListUserAddresses(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserAddresses", reflect.TypeOf((*MockStore)(nil).ListUserAddresses), arg0, arg1)
}
