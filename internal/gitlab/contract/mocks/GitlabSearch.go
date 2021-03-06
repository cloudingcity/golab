// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	gitlab "github.com/xanzy/go-gitlab"

	mock "github.com/stretchr/testify/mock"
)

// GitlabSearch is an autogenerated mock type for the GitlabSearch type
type GitlabSearch struct {
	mock.Mock
}

// MergeRequests provides a mock function with given fields: query, opt, options
func (_m *GitlabSearch) MergeRequests(query string, opt *gitlab.SearchOptions, options ...gitlab.OptionFunc) ([]*gitlab.MergeRequest, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, query, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*gitlab.MergeRequest
	if rf, ok := ret.Get(0).(func(string, *gitlab.SearchOptions, ...gitlab.OptionFunc) []*gitlab.MergeRequest); ok {
		r0 = rf(query, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*gitlab.MergeRequest)
		}
	}

	var r1 *gitlab.Response
	if rf, ok := ret.Get(1).(func(string, *gitlab.SearchOptions, ...gitlab.OptionFunc) *gitlab.Response); ok {
		r1 = rf(query, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, *gitlab.SearchOptions, ...gitlab.OptionFunc) error); ok {
		r2 = rf(query, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MergeRequestsByGroup provides a mock function with given fields: gid, query, opt, options
func (_m *GitlabSearch) MergeRequestsByGroup(gid interface{}, query string, opt *gitlab.SearchOptions, options ...gitlab.OptionFunc) ([]*gitlab.MergeRequest, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, gid, query, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*gitlab.MergeRequest
	if rf, ok := ret.Get(0).(func(interface{}, string, *gitlab.SearchOptions, ...gitlab.OptionFunc) []*gitlab.MergeRequest); ok {
		r0 = rf(gid, query, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*gitlab.MergeRequest)
		}
	}

	var r1 *gitlab.Response
	if rf, ok := ret.Get(1).(func(interface{}, string, *gitlab.SearchOptions, ...gitlab.OptionFunc) *gitlab.Response); ok {
		r1 = rf(gid, query, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(interface{}, string, *gitlab.SearchOptions, ...gitlab.OptionFunc) error); ok {
		r2 = rf(gid, query, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MergeRequestsByProject provides a mock function with given fields: pid, query, opt, options
func (_m *GitlabSearch) MergeRequestsByProject(pid interface{}, query string, opt *gitlab.SearchOptions, options ...gitlab.OptionFunc) ([]*gitlab.MergeRequest, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, query, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*gitlab.MergeRequest
	if rf, ok := ret.Get(0).(func(interface{}, string, *gitlab.SearchOptions, ...gitlab.OptionFunc) []*gitlab.MergeRequest); ok {
		r0 = rf(pid, query, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*gitlab.MergeRequest)
		}
	}

	var r1 *gitlab.Response
	if rf, ok := ret.Get(1).(func(interface{}, string, *gitlab.SearchOptions, ...gitlab.OptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, query, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(interface{}, string, *gitlab.SearchOptions, ...gitlab.OptionFunc) error); ok {
		r2 = rf(pid, query, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Projects provides a mock function with given fields: query, opt, options
func (_m *GitlabSearch) Projects(query string, opt *gitlab.SearchOptions, options ...gitlab.OptionFunc) ([]*gitlab.Project, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, query, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*gitlab.Project
	if rf, ok := ret.Get(0).(func(string, *gitlab.SearchOptions, ...gitlab.OptionFunc) []*gitlab.Project); ok {
		r0 = rf(query, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*gitlab.Project)
		}
	}

	var r1 *gitlab.Response
	if rf, ok := ret.Get(1).(func(string, *gitlab.SearchOptions, ...gitlab.OptionFunc) *gitlab.Response); ok {
		r1 = rf(query, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, *gitlab.SearchOptions, ...gitlab.OptionFunc) error); ok {
		r2 = rf(query, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ProjectsByGroup provides a mock function with given fields: gid, query, opt, options
func (_m *GitlabSearch) ProjectsByGroup(gid interface{}, query string, opt *gitlab.SearchOptions, options ...gitlab.OptionFunc) ([]*gitlab.Project, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, gid, query, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*gitlab.Project
	if rf, ok := ret.Get(0).(func(interface{}, string, *gitlab.SearchOptions, ...gitlab.OptionFunc) []*gitlab.Project); ok {
		r0 = rf(gid, query, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*gitlab.Project)
		}
	}

	var r1 *gitlab.Response
	if rf, ok := ret.Get(1).(func(interface{}, string, *gitlab.SearchOptions, ...gitlab.OptionFunc) *gitlab.Response); ok {
		r1 = rf(gid, query, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(interface{}, string, *gitlab.SearchOptions, ...gitlab.OptionFunc) error); ok {
		r2 = rf(gid, query, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
