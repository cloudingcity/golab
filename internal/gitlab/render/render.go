package render

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/cloudingcity/golab/internal/utils"
	"github.com/xanzy/go-gitlab"
)

// Render is a basic instance.
type Render struct {
	out io.Writer
}

// New returns render instance.
func New(w io.Writer) *Render {
	return &Render{w}
}

// ProjectMRs renders project's merge requests.
func (r *Render) ProjectMRs(mrs []*gitlab.MergeRequest) {
	var (
		rows   [][]string
		row, h []string
	)

	h = []string{"mrid", "title", "url"}

	for _, mr := range mrs {
		mrID := strconv.Itoa(mr.IID)
		row = []string{mrID, mr.Title, mr.WebURL}
		rows = append(rows, row)
	}

	utils.RenderTable(r.out, h, rows)
}

// GlobalMRs renders global's merge requests.
func (r *Render) GlobalMRs(mrs []*gitlab.MergeRequest) {
	var (
		rows   [][]string
		row, h []string
	)

	h = []string{"pid", "mrid", "project", "title", "url"}

	for _, mr := range mrs {
		pID := strconv.Itoa(mr.ProjectID)
		mrID := strconv.Itoa(mr.IID)
		project := utils.ParseMRProject(mr.WebURL)
		row = []string{pID, mrID, project, mr.Title, mr.WebURL}
		rows = append(rows, row)
	}

	utils.RenderTable(r.out, h, rows)
}

// MR renders single merge request information.
func (r *Render) MR(mr *gitlab.MergeRequest) {
	var assignee string
	if mr.Assignee != nil {
		assignee = mr.Assignee.Name
	}
	createdAt := mr.CreatedAt.Format("2006-01-02 15:04:05")
	updatedAt := mr.UpdatedAt.Format("2006-01-02 15:04:05")

	format := `
%s
--------------------------------------------------
%s
--------------------------------------------------
PID         %d  
MRID        %d
Project     %s
Branch      %s -> %s
State       %s
Author      %s
Assignee    %s
CreatedAt   %s
UpdatedAt   %s
Url         %s
`
	fmt.Fprintf(r.out, format,
		mr.Title,
		mr.Description,
		mr.ProjectID,
		mr.IID,
		utils.ParseMRProject(mr.WebURL),
		mr.SourceBranch, mr.TargetBranch,
		mr.State,
		mr.Author.Name,
		assignee,
		createdAt,
		updatedAt,
		mr.WebURL,
	)
}

// LintCI renders lint ci result.
func (r *Render) LintCI(result *gitlab.LintResult) {
	status := strings.Title(result.Status) + "!"
	fmt.Fprintln(r.out, status)

	if len(result.Errors) > 0 {
		fmt.Fprintln(r.out)
		fmt.Fprintln(r.out, "Errors:")

		for _, e := range result.Errors {
			fmt.Fprintf(r.out, "  - %s\n", e)
		}
	}
}

// DependResult is a dependency result struct.
type DependResult struct {
	Project string
	Version string
	Branch  string
	URL     string
}

// Depends renders dependency results.
func (r *Render) Depends(pkgs []*DependResult) {
	var (
		rows   [][]string
		row, h []string
	)

	h = []string{"project", "version", "branch", "url"}

	for _, pkg := range pkgs {
		row = []string{pkg.Project, pkg.Version, pkg.Branch, pkg.URL}
		rows = append(rows, row)
	}

	utils.RenderTable(r.out, h, rows)
}
