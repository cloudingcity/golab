package global

import (
	"io"
	"io/ioutil"

	"github.com/cloudingcity/golab/internal/gitlab/render"
	"github.com/xanzy/go-gitlab"
)

// GitlabValidateService is go-gitlab validate service interface.
type GitlabValidateService interface {
	Lint(content string, options ...gitlab.OptionFunc) (*gitlab.LintResult, *gitlab.Response, error)
}

type validateService struct {
	validate GitlabValidateService
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
