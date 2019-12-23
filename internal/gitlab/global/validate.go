package global

import (
	"io"
	"io/ioutil"

	"github.com/cloudingcity/golab/internal/gitlab/render"
	"github.com/xanzy/go-gitlab"
)

type gitlabValidateService interface {
	Lint(content string, options ...gitlab.OptionFunc) (*gitlab.LintResult, *gitlab.Response, error)
}

type validateService struct {
	validate gitlabValidateService
	out      io.Writer
}

// Lint validate .gitlab-ci.yml whether valid.
func (s *validateService) Lint(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	result, _, err := s.validate.Lint(string(file))
	if err != nil {
		return err
	}

	render.New(s.out).LintCI(result)
	return nil
}
