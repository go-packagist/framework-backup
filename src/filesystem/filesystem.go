package filesystem

import (
	"io/ioutil"
	"os"
	"time"
)

// Exists checks if a file or directory exists
func Exists(path string) bool {
	_, err := Stat(path)

	return err == nil
}

// Size file size
func Size(path string) (int64, error) {
	info, err := Stat(path)

	if err != nil {
		return 0, err
	}

	return info.Size(), nil
}

// IsDir checks if a path is a directory
func IsDir(path string) bool {
	info, err := Stat(path)

	if err != nil {
		return false
	}

	return info.IsDir()
}

// IsFile checks if a path is a file
func IsFile(path string) bool {
	info, err := Stat(path)

	if err != nil {
		return false
	}

	return !info.IsDir()
}

// FileInfo file info, alias Stat
func FileInfo(path string) (os.FileInfo, error) {
	return Stat(path)
}

// Info file info, alias Stat
func Info(path string) (os.FileInfo, error) {
	return Stat(path)
}

// Stat file info
func Stat(path string) (os.FileInfo, error) {
	return os.Stat(path)
}

// ReadDir read directory
func ReadDir(path string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(path)
}

// Get get file content, alias ReadFile
func Get(path string) (string, error) {
	return ReadFile(path)
}

// Read read file content, alias ReadFile
func Read(path string) (string, error) {
	return ReadFile(path)
}

// ReadFile read file content
func ReadFile(path string) (string, error) {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(content), nil
}

func Put(path, content string) error {
	return WriteFile(path, content)
}

func WriteFile(path, content string) error {
	return ioutil.WriteFile(path, []byte(content), 0644)
}

func Mkdir(path string) error {
	return os.Mkdir(path, 0755)
}

func Delete(path string) error {
	return Remove(path)
}

func Remove(path string) error {
	return os.Remove(path)
}

func Move(src, dst string) error {
	return Rename(src, dst)
}

func Rename(src, dst string) error {
	return os.Rename(src, dst)
}

func Prepend(path, contents string) error {
	oldContents, err := Read(path)

	if err != nil {
		return err
	}

	return WriteFile(path, contents+oldContents)
}

func Append(path, contents string) error {
	oldContents, err := Read(path)

	if err != nil {
		return err
	}

	return WriteFile(path, oldContents+contents)
}

func Copy(src, dst string) error {
	return CopyFile(src, dst)
}

func CopyFile(src, dst string) error {
	contents, err := ReadFile(src)

	if err != nil {
		return err
	}

	return WriteFile(dst, contents)
}

func LastModified(path string) (time.Time, error) {
	fileinfo, err := Stat(path)

	if err != nil {
		return time.Time{}, err
	}

	return fileinfo.ModTime(), nil
}
