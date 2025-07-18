// Code generated by mockery v2.53.4. DO NOT EDIT.

package mocks

import (
	domain "chapa/internal/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// TransferRepository is an autogenerated mock type for the TransferRepository type
type TransferRepository struct {
	mock.Mock
}

// CreateTransfer provides a mock function with given fields: ctx, transfer
func (_m *TransferRepository) CreateTransfer(ctx context.Context, transfer *domain.Transfer) error {
	ret := _m.Called(ctx, transfer)

	if len(ret) == 0 {
		panic("no return value specified for CreateTransfer")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Transfer) error); ok {
		r0 = rf(ctx, transfer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTransferByRef provides a mock function with given fields: ctx, ref
func (_m *TransferRepository) GetTransferByRef(ctx context.Context, ref string) (*domain.Transfer, error) {
	ret := _m.Called(ctx, ref)

	if len(ret) == 0 {
		panic("no return value specified for GetTransferByRef")
	}

	var r0 *domain.Transfer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*domain.Transfer, error)); ok {
		return rf(ctx, ref)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.Transfer); ok {
		r0 = rf(ctx, ref)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Transfer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, ref)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTransferStatus provides a mock function with given fields: ctx, ref, status
func (_m *TransferRepository) UpdateTransferStatus(ctx context.Context, ref string, status string) error {
	ret := _m.Called(ctx, ref, status)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTransferStatus")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, ref, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewTransferRepository creates a new instance of TransferRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTransferRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *TransferRepository {
	mock := &TransferRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
