package files

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

type FileInfo struct {
	Name         string
	Path         string
	IsDir        bool
	Size         int64
	LastModified time.Time
	Mode         fs.FileMode // permission
}

func ReadDir(dirPath string) ([]*FileInfo, error) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	var result []*FileInfo
	for _, entry := range entries {
		info, err := ParseDirEntry(dirPath, entry)
		if err != nil {
			return nil, err
		}
		result = append(result, info)
	}
	return result, nil
}

func ReadDirRecur(dirPath string) ([]*FileInfo, error) {
	stat, err := os.Stat(dirPath)
	if err != nil {
		return nil, err
	}
	if !stat.IsDir() {
		return nil, errors.New("this file is not directory")
	}

	var result []*FileInfo
	err = filepath.Walk(dirPath, func(path string, osInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if path == dirPath {
			return nil
		}
		info, err := ParseOsFileInfo(dirPath, osInfo)
		if err != nil {
			return err
		}
		result = append(result, info)
		return nil
	})
	return result, err
}

func ParseDirEntry(base string, entry os.DirEntry) (*FileInfo, error) {
	file, err := entry.Info()
	if err != nil {
		return nil, err
	}
	return ParseOsFileInfo(base, file)
}

func ParseOsFileInfo(base string, file os.FileInfo) (*FileInfo, error) {
	name := file.Name()
	return &FileInfo{
		Name:         name,
		Path:         base + string(filepath.Separator) + name,
		IsDir:        file.IsDir(),
		Size:         file.Size(),
		Mode:         file.Mode(),
		LastModified: file.ModTime(),
	}, nil
}
