package group

import (
	"encoding/json"
	"io"
	"sync"

	"github.com/cloudingcity/golab/internal/gitlab/contract"
	"github.com/cloudingcity/golab/internal/gitlab/render"
	"github.com/cloudingcity/gomod"
	"github.com/xanzy/go-gitlab"
)

type processor func(project *gitlab.Project, pkg string) (version string)

type dependService struct {
	group          string
	gitlabGroup    contract.GitlabGroup
	gitlabRepoFile contract.GitlabRepoFile
	out            io.Writer
}

func (s *dependService) inspect(pkg string, pr processor) error {
	var (
		results []*render.DependResult
		mutex   sync.Mutex
	)

	opt := &gitlab.ListGroupProjectsOptions{
		ListOptions: gitlab.ListOptions{Page: 1, PerPage: 100},
		Simple:      gitlab.Bool(true),
		WithShared:  gitlab.Bool(false),
	}

	for opt.Page != 0 {
		projects, resp, err := s.gitlabGroup.ListGroupProjects(s.group, opt)
		if err != nil {
			return err
		}
		opt.Page = resp.NextPage

		var wg sync.WaitGroup
		wg.Add(len(projects))

		for _, project := range projects {
			project := project

			go func() {
				defer wg.Done()

				if version := pr(project, pkg); version != "" {
					result := &render.DependResult{
						Project: project.Name,
						Version: version,
						Branch:  project.DefaultBranch,
						URL:     project.WebURL,
					}

					mutex.Lock()
					results = append(results, result)
					mutex.Unlock()
				}
			}()
		}
		wg.Wait()
	}

	render.New(s.out).Depends(results)
	return nil
}

func (s *dependService) PHP(pkg string) error {
	return s.inspect(pkg, s.phpProcessor)
}

func (s *dependService) phpProcessor(project *gitlab.Project, pkg string) (version string) {
	composer, err := s.getComposer(project)
	if err != nil {
		return ""
	}

	if version, ok := composer.Require[pkg]; ok {
		return version
	}
	return ""
}

type composer struct {
	Require map[string]string `json:"require"`
}

func (s *dependService) getComposer(project *gitlab.Project) (*composer, error) {
	opt := &gitlab.GetRawFileOptions{
		Ref: gitlab.String(project.DefaultBranch),
	}
	file, _, err := s.gitlabRepoFile.GetRawFile(project.ID, "composer.json", opt)
	if err != nil {
		return nil, err
	}

	var composer *composer
	if err := json.Unmarshal(file, &composer); err != nil {
		return nil, err
	}

	return composer, nil
}

func (s *dependService) GO(pkg string) error {
	return s.inspect(pkg, s.goProcessor)
}

func (s *dependService) goProcessor(project *gitlab.Project, pkg string) (version string) {
	mod, err := s.getGoModule(project)
	if err != nil {
		return ""
	}

	for _, r := range mod.Require {
		if r.Path == pkg {
			return r.Version
		}
	}
	return ""
}

func (s *dependService) getGoModule(project *gitlab.Project) (*gomod.GoMod, error) {
	opt := &gitlab.GetRawFileOptions{
		Ref: gitlab.String(project.DefaultBranch),
	}
	file, _, err := s.gitlabRepoFile.GetRawFile(project.ID, "go.mod", opt)
	if err != nil {
		return nil, err
	}

	mod, err := gomod.Parse(file)
	if err != nil {
		return nil, err
	}

	return mod, nil
}
