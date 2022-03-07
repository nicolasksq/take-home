// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	dao "server/app/dao"

	mock "github.com/stretchr/testify/mock"
)

// ClientAPI is an autogenerated mock type for the ClientAPI type
type ClientAPI struct {
	mock.Mock
}

// BatchListMembers provides a mock function with given fields: contacts, listID
func (_m *ClientAPI) BatchListMembers(contacts []dao.Contact, listID string) ([]dao.Contact, error) {
	ret := _m.Called(contacts, listID)

	var r0 []dao.Contact
	if rf, ok := ret.Get(0).(func([]dao.Contact, string) []dao.Contact); ok {
		r0 = rf(contacts, listID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dao.Contact)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]dao.Contact, string) error); ok {
		r1 = rf(contacts, listID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateList provides a mock function with given fields: listName
func (_m *ClientAPI) CreateList(listName *string) error {
	ret := _m.Called(listName)

	var r0 error
	if rf, ok := ret.Get(0).(func(*string) error); ok {
		r0 = rf(listName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetListsByName provides a mock function with given fields: name
func (_m *ClientAPI) GetListsByName(name string) (*dao.List, error) {
	ret := _m.Called(name)

	var r0 *dao.List
	if rf, ok := ret.Get(0).(func(string) *dao.List); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dao.List)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}