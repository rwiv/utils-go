package files

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func CompressTarGz(srcPath, destPath string) error {
	if Exists(destPath) {
		return fmt.Errorf("%s already exist", destPath)
	}
	outFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	gw := gzip.NewWriter(outFile)
	defer gw.Close()

	tw := tar.NewWriter(gw)
	defer tw.Close()

	err = filepath.Walk(srcPath, func(file string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := tar.FileInfoHeader(fi, fi.Name())
		if err != nil {
			return err
		}
		header.Name, err = filepath.Rel(filepath.Dir(srcPath), file)
		if err != nil {
			return err
		}
		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		if !fi.IsDir() {
			file, err := os.Open(file)
			if err != nil {
				return err
			}
			defer file.Close()

			if _, err := io.Copy(tw, file); err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func ExtractTarGz(srcPath, destDirPath, targetName string) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return fmt.Errorf("Error opening files %s: %v\n", srcPath, err)
	}
	defer srcFile.Close()

	uncompressedStream, err := gzip.NewReader(srcFile)
	if err != nil {
		return fmt.Errorf("ExtractTarGz: NewReader failed: %w", err)
	}
	tarReader := tar.NewReader(uncompressedStream)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("ExtractTarGz: Next() failed: %w", err)
		}

		rootPrefix := ""
		parts := strings.Split(header.Name, string(filepath.Separator))
		if len(parts) > 0 {
			rootPrefix = parts[0]
		}
		newHeaderName := filepath.Join(targetName, strings.TrimPrefix(header.Name, rootPrefix))
		targetPath := filepath.Join(destDirPath, newHeaderName)
		if Exists(targetPath) {
			return fmt.Errorf("%s already exist", targetPath)
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(targetPath, 0755); err != nil {
				return fmt.Errorf("ExtractTarGz: Mkdir() failed: %s", err)
			}
		case tar.TypeReg:
			outFile, err := os.Create(targetPath)
			if err != nil {
				return fmt.Errorf("ExtractTarGz: Create() failed: %s", err)
			}
			if _, err := io.Copy(outFile, tarReader); err != nil {
				return fmt.Errorf("ExtractTarGz: Copy() failed: %s", err)
			}
			outFile.Close()
		default:
			return fmt.Errorf("ExtractTarGz: uknown type: %v in %s", header.Typeflag, header.Name)
		}
	}
	return nil
}
