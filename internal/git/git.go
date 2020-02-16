package git

import (
	"log"
	"os"
	"os/exec"

	"github.com/cloudingcity/golab/internal/utils"
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

// CurrentRepo returns current repo.
func CurrentRepo() string {
	output, err := command("config", "--get", "remote.origin.url").Output()
	if err != nil {
		log.Fatal(err)
	}
	return utils.ParseGitProject(string(output))
}
