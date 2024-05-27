// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	context "context"

	ports "github.com/hanoys/sigma-music/internal/ports"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// TrackObjectStorage is an autogenerated mock type for the ITrackObjectStorage type
type TrackObjectStorage struct {
	mock.Mock
}

// DeleteTrack provides a mock function with given fields: ctx, trackID
func (_m *TrackObjectStorage) DeleteTrack(ctx context.Context, trackID uuid.UUID) error {
	ret := _m.Called(ctx, trackID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTrack")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, trackID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PutTrack provides a mock function with given fields: ctx, req
func (_m *TrackObjectStorage) PutTrack(ctx context.Context, req ports.PutTrackReq) error {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for PutTrack")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ports.PutTrackReq) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewTrackObjectStorage creates a new instance of TrackObjectStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTrackObjectStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *TrackObjectStorage {
	mock := &TrackObjectStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
