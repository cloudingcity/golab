package group

import (
	"bytes"
	"errors"
	"testing"

	"github.com/cloudingcity/golab/internal/gitlab/contract/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/xanzy/go-gitlab"
)

func TestInspect(t *testing.T) {
	group := "foo"
	pkg := "foo/cool"

	mockGitlabGroup := &mocks.GitlabGroup{}
	mockGitlabGroupOpt := &gitlab.ListGroupProjectsOptions{
		ListOptions: gitlab.ListOptions{Page: 1, PerPage: 100},
		Simple:      gitlab.Bool(true),
		WithShared:  gitlab.Bool(false),
	}
	mockGitlabGroup.On("ListGroupProjects", group, mockGitlabGroupOpt).
		Once().
		Return(
			[]*gitlab.Project{
				{
					Name:          "cool",
					DefaultBranch: "master",
					WebURL:        "https://gitlab.com/foo/cool",
				},
			},
			&gitlab.Response{NextPage: 0},
			nil,
		)

	mockProcessor := func(project *gitlab.Project, pkg string) (version string) {
		return "v1.2.3"
	}

	buf := &bytes.Buffer{}
	s := &dependService{
		group:       group,
		gitlabGroup: mockGitlabGroup,
		out:         buf,
	}
	s.inspect(pkg, mockProcessor)

	got := buf.String()
	wants := []string{"cool", "v1.2.3", "master", "https://gitlab.com/foo/cool"}
	for _, want := range wants {
		assert.Contains(t, got, want)
	}

	mockGitlabGroup.AssertExpectations(t)
}

func TestInspectErr(t *testing.T) {
	mockGitlabGroup := &mocks.GitlabGroup{}
	mockGitlabGroup.On("ListGroupProjects", mock.Anything, mock.Anything).
		Once().
		Return(nil, nil, errors.New(""))
	s := &dependService{
		gitlabGroup: mockGitlabGroup,
	}
	err := s.inspect("", nil)

	assert.Error(t, err)
	mockGitlabGroup.AssertExpectations(t)
}

func TestPHPProcessor(t *testing.T) {
	pkg := "foo/cool"
	project := &gitlab.Project{
		ID:            999,
		Name:          "cool",
		DefaultBranch: "master",
		WebURL:        "https://gitlab.com/foo/cool",
	}

	mockGitlabRepoFile := &mocks.GitlabRepoFile{}
	mockGitlabRepoFileOpt := &gitlab.GetRawFileOptions{
		Ref: gitlab.String(project.DefaultBranch),
	}
	mockGitlabRepoFile.On("GetRawFile", project.ID, "composer.json", mockGitlabRepoFileOpt).
		Once().
		Return([]byte(`{"require":{"foo/cool": "v1.2.3"}}`), nil, nil)

	s := &dependService{
		gitlabRepoFile: mockGitlabRepoFile,
	}

	want := "v1.2.3"
	got := s.phpProcessor(project, pkg)
	assert.Equal(t, want, got)

	mockGitlabRepoFile.AssertExpectations(t)
}

func TestGOProcessor(t *testing.T) {
	pkg := "foo/cool"
	project := &gitlab.Project{
		ID:            999,
		Name:          "cool",
		DefaultBranch: "master",
		WebURL:        "https://gitlab.com/foo/cool",
	}

	mockGitlabRepoFile := &mocks.GitlabRepoFile{}
	mockGitlabRepoFileOpt := &gitlab.GetRawFileOptions{
		Ref: gitlab.String(project.DefaultBranch),
	}
	mockGitlabRepoFile.On("GetRawFile", project.ID, "go.mod", mockGitlabRepoFileOpt).
		Once().
		Return([]byte(`module foo/bar

go 1.13

require foo/cool v1.2.3`), nil, nil)

	s := &dependService{
		gitlabRepoFile: mockGitlabRepoFile,
	}

	want := "v1.2.3"
	got := s.goProcessor(project, pkg)
	assert.Equal(t, want, got)

	mockGitlabRepoFile.AssertExpectations(t)
}
