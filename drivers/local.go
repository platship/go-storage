package drivers

import (
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/fasthey/go-storage"

	"github.com/fasthey/go-utils/osx"
)

type Local struct {
	Path string `json:"path"`
}

func (l *Local) Init() error {
	if l.Path == "" {
		return errors.New("path undefined")
	}
	return osx.CreateDirIsNotExist(l.Path, 0777)
}

func (l *Local) Close() error {
	return nil
}

func (l *Local) Get(key string) (*storage.GetValue, error) {
	if l.Path == "" {
		return nil, errors.New("path undefined")
	}
	f, err := os.Open(l.filePath(key))
	if err != nil {
		return nil, err
	}
	return storage.NewGetValue(f), nil
}

func (l *Local) Set(key string, val *storage.SetValue) (err error) {
	if l.Path == "" {
		return errors.New("path undefined")
	}
	path := l.filePath(key)
	_ = l.createDir(path)

	f, err := os.Create(path)
	if err != nil {
		return
	}
	defer f.Close()
	_, err = io.Copy(f, val.Reader)
	return
}

func (l *Local) Delete(key string) error {
	if l.Path == "" {
		return errors.New("path undefined")
	}
	path := l.filePath(key)
	if path == "" || path == "/" {
		return errors.New("path is empty or root dir")
	}
	return os.Remove(path)
}

func (l *Local) createDir(path string) error {
	dir, _ := filepath.Split(path)
	return osx.CreateDirIsNotExist(dir, 0777)
}

func (l *Local) filePath(key string) string {
	return filepath.Join(l.Path, key)
}
