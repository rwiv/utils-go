package paths

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetExt(filePath string) string {
	ext := filepath.Ext(filePath)
	if strings.HasPrefix(ext, ".") {
		ext = strings.TrimPrefix(ext, ".")
	}
	return ext
}

func ProjectRoot() (string, error) {
	ex, err := filepath.Abs("")
	if err != nil {
		return "", err
	}
	chunks := strings.Split(ex, string(os.PathSeparator))
	var pkgIdx = -1
	for i, chunk := range chunks {
		if chunk == "pkg" {
			pkgIdx = i
		}
	}
	if pkgIdx == -1 {
		return "", fmt.Errorf("not found pkgIdx")
	}
	return strings.Join(chunks[0:pkgIdx], string(os.PathSeparator)), nil
}
