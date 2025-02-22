// Code generated by mockery v2.52.2. DO NOT EDIT.

package mocks

import (
	context "context"
	models "merchshop/internal/models"

	mock "github.com/stretchr/testify/mock"
)

// DataBase is an autogenerated mock type for the DataBase type
type DataBase struct {
	mock.Mock
}

// GetCoinHistoryByUserID provides a mock function with given fields: ctx, userID
func (_m *DataBase) GetCoinHistoryByUserID(ctx context.Context, userID int) (*models.CoinHistory, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetCoinHistoryByUserID")
	}

	var r0 *models.CoinHistory
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*models.CoinHistory, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *models.CoinHistory); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.CoinHistory)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCoinsByUserID provides a mock function with given fields: ctx, userID
func (_m *DataBase) GetCoinsByUserID(ctx context.Context, userID int) (int, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetCoinsByUserID")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (int, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) int); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInventoryByUserID provides a mock function with given fields: ctx, userID
func (_m *DataBase) GetInventoryByUserID(ctx context.Context, userID int) (*[]models.Merch, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetInventoryByUserID")
	}

	var r0 *[]models.Merch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*[]models.Merch, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *[]models.Merch); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]models.Merch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewDataBase creates a new instance of DataBase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDataBase(t interface {
	mock.TestingT
	Cleanup(func())
}) *DataBase {
	mock := &DataBase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
