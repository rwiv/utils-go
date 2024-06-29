package files

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCopyFile(t *testing.T) {
	root, err := ProjectRoot()
	if err != nil {
		t.Fatal(err)
	}
	srcPath := filepath.Join(root, "tests", "test.txt")
	dstPath := filepath.Join(root, "tests", "copied.txt")

	err = CopyFile(srcPath, dstPath)
	if err != nil {
		t.Fatal(err)
	}
}

func ProjectRoot() (string, error) {
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
		return "", fmt.Errorf("not found pkgIdx")
	}
	return strings.Join(chunks[0:pkgIdx+1], string(os.PathSeparator)), nil
}
