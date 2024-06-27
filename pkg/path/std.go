package path

import stdpath "path"

var ErrBadPattern = stdpath.ErrBadPattern

func Clean(path string) string {
	return stdpath.Clean(path)
}

func Split(path string) (dir, file string) {
	return stdpath.Split(path)
}

func Join(elem ...string) string {
	return stdpath.Join(elem...)
}

func Ext(path string) string {
	return stdpath.Ext(path)
}

func Base(path string) string {
	return stdpath.Base(path)
}

func IsAbs(path string) bool {
	return stdpath.IsAbs(path)
}

func Dir(path string) string {
	return stdpath.Dir(path)
}

func Match(pattern, name string) (matched bool, err error) {
	return stdpath.Match(pattern, name)
}
