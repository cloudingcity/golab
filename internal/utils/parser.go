package utils

import (
	"regexp"
	"strings"
)

// ParseGitProject parses git remote url and project name.
func ParseGitProject(url string) string {
	var elems []string

	if strings.Contains(url, "@") {
		elems = strings.Split(url, ":")
	} else {
		elems = strings.Split(url, "//")
		elems = strings.SplitN(elems[1], "/", 2)
	}

	return strings.TrimSuffix(elems[1], ".git")
}

// ParseMRProject parses merge request url and return project name.
func ParseMRProject(url string) string {
	return regexp.MustCompile(`https://[\w.]+/(?P<project>.+)/merge_requests.+`).
		FindStringSubmatch(url)[1]
}
