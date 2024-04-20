package dugtrio

import (
	"path/filepath"
	"strings"
)

var (
	ProjectRoot = ""
)

func init() {
	ProjectRoot = sourceFilepathClean(strings.Trim(ProjectRoot, "\""))
}

func HasProjectRoot() bool {
	return len(ProjectRoot) != 0
}

func sourceFilepathClean(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ReplaceAll(filepath.Clean(s), "\\", "/")
}

func convertSourceFilename(s string) string {
	if len(s) == 0 || len(ProjectRoot) == 0 {
		return s
	}
	s = sourceFilepathClean(s)
	if strings.HasPrefix(s, ProjectRoot) {
		return strings.TrimSpace(strings.TrimLeft(s, ProjectRoot))
	}
	return s
}
