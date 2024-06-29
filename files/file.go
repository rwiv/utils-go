package files

import (
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

type Info struct {
	Name        string
	AbsPath     string
	IsDirectory bool
	Size        int64
	Mode        fs.FileMode
	ModeTime    time.Time
}

func ReadDir(abs string) ([]*Info, error) {
	entries, err := os.ReadDir(abs)
	if err != nil {
		return nil, err
	}
	result := make([]*Info, 0)
	for _, entry := range entries {
		info, err := ParseDirEntry(abs, entry)
		if err != nil {
			return nil, err
		}
		result = append(result, info)
	}
	return result, nil
}

func ParseDirEntry(base string, entry os.DirEntry) (*Info, error) {
	file, err := entry.Info()
	if err != nil {
		return nil, err
	}
	name := file.Name()
	return &Info{
		Name:        name,
		AbsPath:     base + string(filepath.Separator) + name,
		IsDirectory: file.IsDir(),
		Size:        file.Size(),
		Mode:        file.Mode(),
		ModeTime:    file.ModTime(),
	}, nil
}

func Exists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}
