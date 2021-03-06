// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	gitlab "github.com/xanzy/go-gitlab"

	mock "github.com/stretchr/testify/mock"
)

// GitlabGroup is an autogenerated mock type for the GitlabGroup type
type GitlabGroup struct {
	mock.Mock
}

// ListGroupProjects provides a mock function with given fields: gid, opt, options
func (_m *GitlabGroup) ListGroupProjects(gid interface{}, opt *gitlab.ListGroupProjectsOptions, options ...gitlab.OptionFunc) ([]*gitlab.Project, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, gid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*gitlab.Project
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.ListGroupProjectsOptions, ...gitlab.OptionFunc) []*gitlab.Project); ok {
		r0 = rf(gid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*gitlab.Project)
		}
	}

	var r1 *gitlab.Response
	if rf, ok := ret.Get(1).(func(interface{}, *gitlab.ListGroupProjectsOptions, ...gitlab.OptionFunc) *gitlab.Response); ok {
		r1 = rf(gid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(interface{}, *gitlab.ListGroupProjectsOptions, ...gitlab.OptionFunc) error); ok {
		r2 = rf(gid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
