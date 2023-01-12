// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	book "cleanarch/features/book"

	mock "github.com/stretchr/testify/mock"
)

// BookData is an autogenerated mock type for the BookData type
type BookData struct {
	mock.Mock
}

// Add provides a mock function with given fields: userID, newBook
func (_m *BookData) Add(userID uint, newBook book.BookCore) (book.BookCore, error) {
	ret := _m.Called(userID, newBook)

	var r0 book.BookCore
	if rf, ok := ret.Get(0).(func(uint, book.BookCore) book.BookCore); ok {
		r0 = rf(userID, newBook)
	} else {
		r0 = ret.Get(0).(book.BookCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, book.BookCore) error); ok {
		r1 = rf(userID, newBook)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: userID, bookID
func (_m *BookData) Delete(userID uint, bookID uint) error {
	ret := _m.Called(userID, bookID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint) error); ok {
		r0 = rf(userID, bookID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: bookID, updateBook
func (_m *BookData) Update(bookID uint, updateBook book.BookCore) (book.BookCore, error) {
	ret := _m.Called(bookID, updateBook)

	var r0 book.BookCore
	if rf, ok := ret.Get(0).(func(uint, book.BookCore) book.BookCore); ok {
		r0 = rf(bookID, updateBook)
	} else {
		r0 = ret.Get(0).(book.BookCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, book.BookCore) error); ok {
		r1 = rf(bookID, updateBook)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewBookData interface {
	mock.TestingT
	Cleanup(func())
}

// NewBookData creates a new instance of BookData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBookData(t mockConstructorTestingTNewBookData) *BookData {
	mock := &BookData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}