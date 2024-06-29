package files

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestReadDir(t *testing.T) {
	root, err := testProjectRoot()
	if err != nil {
		t.Fatal(err)
	}
	dirPath := filepath.Join(root, "tests", "recur")

	// test ReadDir
	fmt.Println("ReadDir")
	infos, err := ReadDir(dirPath)
	if err != nil {
		t.Fatal(err)
	}
	for _, info := range infos {
		fmt.Println(info)
	}

	// test ReadDirRecur
	fmt.Println("ReadDirRecur")
	infos, err = ReadDirRecur(dirPath)
	if err != nil {
		t.Fatal(err)
	}
	for _, info := range infos {
		fmt.Println(info)
	}
}
