// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	tag "github.com/limeiralucas/chrono-cli/internal/tag"
	mock "github.com/stretchr/testify/mock"
)

// TagProvider is an autogenerated mock type for the TagProvider type
type TagProvider struct {
	mock.Mock
}

// Create provides a mock function with given fields: name
func (_m *TagProvider) Create(name string) (int, error) {
	ret := _m.Called(name)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *TagProvider) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *TagProvider) GetAll() ([]tag.Tag, error) {
	ret := _m.Called()

	var r0 []tag.Tag
	if rf, ok := ret.Get(0).(func() []tag.Tag); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]tag.Tag)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: id
func (_m *TagProvider) GetById(id int) (tag.Tag, error) {
	ret := _m.Called(id)

	var r0 tag.Tag
	if rf, ok := ret.Get(0).(func(int) tag.Tag); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(tag.Tag)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByName provides a mock function with given fields: name
func (_m *TagProvider) GetByName(name string) (tag.Tag, error) {
	ret := _m.Called(name)

	var r0 tag.Tag
	if rf, ok := ret.Get(0).(func(string) tag.Tag); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(tag.Tag)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateName provides a mock function with given fields: id, name
func (_m *TagProvider) UpdateName(id int, name string) error {
	ret := _m.Called(id, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string) error); ok {
		r0 = rf(id, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewTagProvider interface {
	mock.TestingT
	Cleanup(func())
}

// NewTagProvider creates a new instance of TagProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTagProvider(t mockConstructorTestingTNewTagProvider) *TagProvider {
	mock := &TagProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}