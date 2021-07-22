package wrapper

import (
	"io/fs"
	"os"
)

// OperatingSystem is an interface which provides method signatures for interacting with the operating system
type OperatingSystem interface {
	OpenFile(string, int, fs.FileMode) (*os.File, error)
	ReadFile(string) ([]byte, error)
	WriteFile(string, []byte, os.FileMode) error
	IsNotExist(error) bool
	Stat(string) (os.FileInfo, error)
	Remove(string) error
}

// DefaultOS is the wrapper for the default OS
type DefaultOS struct {
}

var (
	_ OperatingSystem = (*DefaultOS)(nil)
)

// OpenFile is a wrapper for os.OpenFile
func (d *DefaultOS) OpenFile(name string, flag int, perm fs.FileMode) (*os.File, error) {
	return os.OpenFile(name, flag, perm)
}

// ReadFile is a wrapper for os.ReadFile
func (d *DefaultOS) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

// WriteFile is a wrapper for os.WriteFile
func (d *DefaultOS) WriteFile(name string, data []byte, perm os.FileMode) error {
	return os.WriteFile(name, data, perm)
}

// IsNotExist is a wrapper for os.IsNotExist
func (d *DefaultOS) IsNotExist(err error) bool {
	return os.IsNotExist(err)
}

// Stat is a wrapper for os.Stat
func (d *DefaultOS) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

// Remove is a wrapper for os.Remove
func (d *DefaultOS) Remove(name string) error {
	return os.Remove(name)
}
