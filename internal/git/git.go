package git

import (
	"bytes"
	"log"
	"os"
	"os/exec"

	"github.com/cloudingcity/golab/internal/utils"
)

var command = func(args ...string) *exec.Cmd {
	return exec.Command("git", args...)
}

// Push update remote refs along with associated objects
func Push(ref string) error {
	cmd := command("push", "--set-upstream", "origin", ref)
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// Clone clone a repository form GitLab.
func Clone(repo, dir string) error {
	args := []string{"clone", repo}
	if len(dir) != 0 {
		args = append(args, dir)
	}

	cmd := command(args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

// CurrentRepo returns current repo.
func CurrentRepo() string {
	output, err := command("config", "--get", "remote.origin.url").Output()
	if err != nil {
		log.Fatal("fatal: not a git repository (or any of the parent directories): .git")
	}
	return utils.ParseGitProject(string(output))
}

// CurrentBranch returns current branch.
func CurrentBranch() string {
	output, err := command("rev-parse", "--abbrev-ref", "HEAD").CombinedOutput()
	if err != nil {
		log.Fatal(string(output))
	}
	return string(bytes.Trim(output, "\n"))
}
