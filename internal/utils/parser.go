package utils

import (
	"regexp"
	"strings"
)

// ParseGitProject parses git remote url and project name.
func ParseGitProject(url string) string {
	var r *regexp.Regexp
	if strings.Contains(url, "@") {
		r = regexp.MustCompile(`.+:(?P<project>.+)\.git`)
	} else {
		r = regexp.MustCompile(`https://[\w.]+/(?P<project>.+)\.git`)
	}
	return r.FindStringSubmatch(url)[1]
}

// ParseMRProject parses merge request url and return project name.
func ParseMRProject(url string) string {
	return regexp.MustCompile(`https://[\w.]+/(?P<project>.+)/merge_requests.+`).
		FindStringSubmatch(url)[1]
}
