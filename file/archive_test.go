package file

import (
	"github.com/rwiv/utils-go/path"
	"path/filepath"
	"testing"
)

func TestCompressTarGz(t *testing.T) {
	root, err := path.ProjectRoot()
	if err != nil {
		t.Fatal(err)
	}
	srcPath := filepath.Join(root, "tests", "archive")
	destPath := filepath.Join(root, "tests", "archive.tar.gz")

	err = CompressTarGz(srcPath, destPath)
	if err != nil {
		t.Fatal(err)
	}
}

func TestExtractTarGz(t *testing.T) {
	root, err := path.ProjectRoot()
	if err != nil {
		t.Fatal(err)
	}
	srcPath := filepath.Join(root, "tests", "archive")
	destPath := filepath.Join(root, "tests", "archive.tar.gz")

	err = CompressTarGz(srcPath, destPath)
	if err != nil {
		t.Fatal(err)
	}

	outPath := filepath.Join(root, "tests")
	err = ExtractTarGz(destPath, outPath, "archive (2)")
	if err != nil {
		t.Fatal(err)
	}
}
