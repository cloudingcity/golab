package global

import (
	"errors"
	"os"
	"testing"

	"github.com/cloudingcity/golab/internal/gitlab/global/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/xanzy/go-gitlab"
)

func TestLint(t *testing.T) {
	t.Run("file not exists", func(t *testing.T) {
		s := &validateService{}
		err := s.Lint("foo/bar")

		assert.Error(t, err)
	})

	t.Run("lint", func(t *testing.T) {
		v := &mocks.GitlabValidateService{}
		v.On("Lint", "HelloWord\n").
			Once().
			Return(&gitlab.LintResult{}, &gitlab.Response{}, errors.New(""))

		s := &validateService{validate: v}
		dir, _ := os.Getwd()
		s.Lint(dir + "/../../../test/.gitlab-ci.yml")

		v.AssertExpectations(t)
	})
}
