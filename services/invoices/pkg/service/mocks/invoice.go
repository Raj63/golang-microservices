// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/service/invoice.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	model "github.com/Raj63/golang-microservices/services/invoices/pkg/model"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockInvoice is a mock of Invoice interface.
type MockInvoice struct {
	ctrl     *gomock.Controller
	recorder *MockInvoiceMockRecorder
}

// MockInvoiceMockRecorder is the mock recorder for MockInvoice.
type MockInvoiceMockRecorder struct {
	mock *MockInvoice
}

// NewMockInvoice creates a new mock instance.
func NewMockInvoice(ctrl *gomock.Controller) *MockInvoice {
	mock := &MockInvoice{ctrl: ctrl}
	mock.recorder = &MockInvoiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInvoice) EXPECT() *MockInvoiceMockRecorder {
	return m.recorder
}

// Approve mocks base method.
func (m *MockInvoice) Approve(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Approve", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Approve indicates an expected call of Approve.
func (mr *MockInvoiceMockRecorder) Approve(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Approve", reflect.TypeOf((*MockInvoice)(nil).Approve), ctx, id)
}

// Create mocks base method.
func (m *MockInvoice) Create(ctx context.Context, invoice *model.Invoice) (*model.Invoice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, invoice)
	ret0, _ := ret[0].(*model.Invoice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockInvoiceMockRecorder) Create(ctx, invoice interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockInvoice)(nil).Create), ctx, invoice)
}

// Get mocks base method.
func (m *MockInvoice) Get(ctx context.Context, id uuid.UUID) (*model.Invoice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*model.Invoice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockInvoiceMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockInvoice)(nil).Get), ctx, id)
}
