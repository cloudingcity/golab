package global

import (
	"io"
	"io/ioutil"

	"github.com/cloudingcity/golab/internal/gitlab/contract"
	"github.com/cloudingcity/golab/internal/gitlab/render"
)

type validateService struct {
	validate contract.GitlabValidate
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
