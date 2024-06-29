package files

import (
	"io"
	"os"
)

func Exists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func CopyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close() // ignore error: file was opened read-only.

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() {
		cerr := dstFile.Close()
		if err == nil {
			err = cerr
		}
	}()

	_, err = io.Copy(dstFile, srcFile)
	return err
}
