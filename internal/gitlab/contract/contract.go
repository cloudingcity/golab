package contract

import "github.com/xanzy/go-gitlab"

// GitlabMergeRequests is go-gitlab merge request service interface.
type GitlabMergeRequests interface {
	ListMergeRequests(opt *gitlab.ListMergeRequestsOptions, options ...gitlab.OptionFunc) ([]*gitlab.MergeRequest, *gitlab.Response, error)
	ListProjectMergeRequests(pid interface{}, opt *gitlab.ListProjectMergeRequestsOptions, options ...gitlab.OptionFunc) ([]*gitlab.MergeRequest, *gitlab.Response, error)
	GetMergeRequest(pid interface{}, mergeRequest int, opt *gitlab.GetMergeRequestsOptions, options ...gitlab.OptionFunc) (*gitlab.MergeRequest, *gitlab.Response, error)
}

// GitlabSearch is go-gitlab search service interface.
type GitlabSearch interface {
	MergeRequestsByProject(pid interface{}, query string, opt *gitlab.SearchOptions, options ...gitlab.OptionFunc) ([]*gitlab.MergeRequest, *gitlab.Response, error)
	MergeRequestsByGroup(gid interface{}, query string, opt *gitlab.SearchOptions, options ...gitlab.OptionFunc) ([]*gitlab.MergeRequest, *gitlab.Response, error)
	MergeRequests(query string, opt *gitlab.SearchOptions, options ...gitlab.OptionFunc) ([]*gitlab.MergeRequest, *gitlab.Response, error)
}

// GitlabValidate is go-gitlab validate service interface.
type GitlabValidate interface {
	Lint(content string, options ...gitlab.OptionFunc) (*gitlab.LintResult, *gitlab.Response, error)
}
