package utils

import "strings"

// ParseRepo parses git remote url.
func ParseRepo(url string) string {
	var elems []string

	if strings.Contains(url, "@") {
		elems = strings.Split(url, ":")
	} else {
		elems = strings.Split(url, "//")
		elems = strings.SplitN(elems[1], "/", 2)
	}

	return strings.TrimSuffix(elems[1], ".git")
}
