// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import db "github.com/derek-elliott/url-shortener/db"
import mock "github.com/stretchr/testify/mock"

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// CollectStats provides a mock function with given fields:
func (_m *Store) CollectStats() (*db.Stats, error) {
	ret := _m.Called()

	var r0 *db.Stats
	if rf, ok := ret.Get(0).(func() *db.Stats); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*db.Stats)
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

// CreateShortURL provides a mock function with given fields: shortURL
func (_m *Store) CreateShortURL(shortURL *db.ShortURL) error {
	ret := _m.Called(shortURL)

	var r0 error
	if rf, ok := ret.Get(0).(func(*db.ShortURL) error); ok {
		r0 = rf(shortURL)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteShortURL provides a mock function with given fields: token
func (_m *Store) DeleteShortURL(token string) error {
	ret := _m.Called(token)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllURLTokens provides a mock function with given fields:
func (_m *Store) GetAllURLTokens() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
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

// GetShortURL provides a mock function with given fields: token
func (_m *Store) GetShortURL(token string) (*db.ShortURL, error) {
	ret := _m.Called(token)

	var r0 *db.ShortURL
	if rf, ok := ret.Get(0).(func(string) *db.ShortURL); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*db.ShortURL)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InitDB provides a mock function with given fields: user, pass, name, host, port
func (_m *Store) InitDB(user string, pass string, name string, host string, port int) error {
	ret := _m.Called(user, pass, name, host, port)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string, int) error); ok {
		r0 = rf(user, pass, name, host, port)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateShortURL provides a mock function with given fields: shortURL
func (_m *Store) UpdateShortURL(shortURL *db.ShortURL) error {
	ret := _m.Called(shortURL)

	var r0 error
	if rf, ok := ret.Get(0).(func(*db.ShortURL) error); ok {
		r0 = rf(shortURL)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
