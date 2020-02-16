package git

import (
	"os"
	"os/exec"
)

var command = func(args ...string) *exec.Cmd {
	return exec.Command("git", args...)
}

// Clone clone a repository form GitLab.
func Clone(repo, dir string) error {
	args := []string{"clone", repo}
	if len(dir) != 0 {
		args = append(args, dir)
	}

	cmd := command(args...)
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
