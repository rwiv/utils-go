package files

import (
	"path/filepath"
	"testing"
)

func TestCopyFile(t *testing.T) {
	root, err := testProjectRoot()
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
