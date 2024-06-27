package file

import (
	"path/filepath"
	"testing"

	"github.com/rwiv/utils-go/pkg/path"
)

func TestCompressTarGz(t *testing.T) {
	root, err := path.RootPath()
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
	root, err := path.RootPath()
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
