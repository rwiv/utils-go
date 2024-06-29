package files

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func testProjectRoot() (string, error) {
	ex, err := filepath.Abs("")
	if err != nil {
		return "", err
	}
	chunks := strings.Split(ex, string(os.PathSeparator))
	var pkgIdx = -1
	for i, chunk := range chunks {
		if chunk == "utils-go" {
			pkgIdx = i
		}
	}
	if pkgIdx == -1 {
		return "", fmt.Errorf("not found idx")
	}
	return strings.Join(chunks[0:pkgIdx+1], string(os.PathSeparator)), nil
}
