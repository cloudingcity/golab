package git

import (
	"net/url"
	"os"
	"os/exec"
)

// Git is a command struct.
type Git struct {
	cmd *exec.Cmd
	URL *url.URL
}

// New returns an initialized Git instance.
func New(url *url.URL) *Git {
	cmd := exec.Command("git")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return &Git{cmd: cmd, URL: url}
}

// Clone clone a repository form GitLab.
func (g *Git) Clone(project, dir string) *Git {
	repo := "git@" + g.URL.Host + ":" + project
	args := []string{"clone", repo}

	if len(dir) != 0 {
		args = append(args, dir)
	}

	g.setArgs(args)
	return g
}

// Run execute command.
func (g *Git) Run() {
	g.cmd.Run()
}

// String returns a human-readable description of command.
func (g *Git) String() string {
	return g.cmd.String()
}

func (g *Git) setArgs(args []string) {
	g.cmd.Args = append([]string{g.cmd.Path}, args...)
}
