package base

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func FileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

func DirExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func RemoveDir(path string) error {
	return os.RemoveAll(path)
}

func MakeDirs(path string) error {
	return os.MkdirAll(path, 0775)
}

func CurrentDir() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	return dir, nil
}

func WriteFile(path string, data []byte) (err error) {
	err = ioutil.WriteFile(path, data, 0644)
	return
}
