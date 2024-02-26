package ditto

import (
	"path/filepath"
	"strings"
)

func projectFilenameClean(filename string) string {
	if len(filename) <= 0 {
		return filename
	}
	return strings.ReplaceAll(filepath.Clean(filepath.Clean(filename)), "\\", "/")
}
