package wrapper

import (
	"io/fs"
	"os"
)

type OperatingSystem interface {
	OpenFile(string, int, fs.FileMode) (*os.File, error)
	ReadFile(string) ([]byte, error)
	WriteFile(string, []byte, os.FileMode) error
	IsNotExist(error) bool
	Stat(string) (os.FileInfo, error)
	Remove(string) error
}

type DefaultOS struct {
}

var (
	_ OperatingSystem = (*DefaultOS)(nil)
)

func (d *DefaultOS) OpenFile(name string, flag int, perm fs.FileMode) (*os.File, error) {
	return os.OpenFile(name, flag, perm)
}

func (d *DefaultOS) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func (d *DefaultOS) WriteFile(name string, data []byte, perm os.FileMode) error {
	return os.WriteFile(name, data, perm)
}

func (d *DefaultOS) IsNotExist(err error) bool {
	return os.IsNotExist(err)
}

func (d *DefaultOS) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

func (d *DefaultOS) Remove(name string) error {
	return os.Remove(name)
}
