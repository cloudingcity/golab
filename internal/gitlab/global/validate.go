package global

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"

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

	lint, _, err := s.validate.Lint(string(file))
	if err != nil {
		return err
	}

	s.renderLint(lint)
	return nil
}

func (s *validateService) renderLint(lint *gitlab.LintResult) {
	status := strings.Title(lint.Status) + "!"
	fmt.Fprintln(s.out, status)

	if len(lint.Errors) > 0 {
		fmt.Fprintln(s.out)
		fmt.Fprintln(s.out, "Errors:")

		for _, e := range lint.Errors {
			fmt.Fprintf(s.out, "  - %s\n", e)
		}
	}
}
