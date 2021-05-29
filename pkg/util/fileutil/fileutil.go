package fileutil

import (
	"io/ioutil"
	"os"
)

type FileInfo os.FileInfo

const (
	NewFilePerm = 0644
)

// GetFileInfo returns a FileInfo describing the named file
func GetFileInfo(name string) (FileInfo, error) {
	fi, err := os.Stat(name)
	return fi, err
}

// Exists checks if the given filename exists
func Exists(filename string) (bool, error) {
	if _, err := GetFileInfo(filename); err != nil {
		if os.IsNotExist(err) {
			return false, err
		}
	}
	return true, nil
}

// Mkdir creates a new directory with the specified name and permission bits.
func Mkdir(name string, perm os.FileMode) error {
	return os.Mkdir(name, perm)
}

// MkdirAll creates a directory named path, along with any necessary parents.
func MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

// MkdirP creates a directory named path, along with any necessary parents.
// MkdirP is equivalent to MkdirAll.
func MkdirP(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

// Chmod changes the mode of the named file to mode.
func Chmod(name string, mode os.FileMode) error {
	return os.Chmod(name, mode)
}

// Chown changes the numeric uid and gid of the named file.
func Chown(name string, uid, gid int) error {
	return os.Chown(name, uid, gid)
}

// Write writes data to a file named by filename.
func Write(filename string, content []byte) error {
	return ioutil.WriteFile(filename, content, NewFilePerm)
}

// Read reads the file named by filename and returns the contents.
func Read(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}
