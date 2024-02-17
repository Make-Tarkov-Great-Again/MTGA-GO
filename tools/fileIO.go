package tools

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/goccy/go-json"
)

const (
	FileDoesNotExist    string = "file does not exist"
	FailToReadFile      string = "failed to read file"
	DirectoryExists     string = "directory already exists"
	FailToReadDirectory string = "failed to read directory"
	SswFormat           string = "%s: %s: %w"
	SsFormat            string = "%s: %s"
)

// WriteToFile writes the given string of data to the specified file path
func WriteToFile(filePath string, data any) error {
	path := GetAbsolutePathFrom(filePath)
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return err
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetEscapeHTML(false)  // don't escape Unicode
	encoder.SetIndent("", "    ") //4 space indentation

	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}

// GetAbsolutePathFrom returns the absolute path from a relative path
func GetAbsolutePathFrom(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	cwd, _ := os.Getwd()
	return filepath.Join(cwd, path)
}

// CreateDirectory creates a directory at the specified path
func CreateDirectory(filePath string) error {
	path := GetAbsolutePathFrom(filePath)

	if FileExist(path) {
		return fmt.Errorf(SsFormat, DirectoryExists, filePath)
	}

	if err := os.MkdirAll(path, 0755); err != nil {
		log.Fatal(err)
	}
	return nil
}

// FileExist checks if a file exists at the specified path
func FileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil || !os.IsNotExist(err)
}

// ReadFile reads the file at filePath and returns its contents as a byte slice.
func ReadFile(filePath string) ([]byte, error) {
	path := GetAbsolutePathFrom(filePath)

	if !FileExist(path) {
		return nil, fmt.Errorf(SsFormat, FileDoesNotExist, filePath)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf(SswFormat, FailToReadFile, filePath, err)
	}

	return data, nil
}

// GetDirectoriesFrom returns a list of directories from a file path
func GetDirectoriesFrom(filePath string) (map[string]struct{}, error) {
	path := GetAbsolutePathFrom(filePath)
	if !FileExist(path) {
		return nil, fmt.Errorf(SsFormat, FileDoesNotExist, filePath)
	}

	directory, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf(SswFormat, FailToReadDirectory, filePath, err)
	}

	files := make(map[string]struct{})
	for _, file := range directory {
		if file.IsDir() {
			files[file.Name()] = struct{}{}
		}
	}
	return files, nil
}

// GetFilesFrom returns a list of files from a file path
func GetFilesFrom(filePath string) (map[string]struct{}, error) {
	path := GetAbsolutePathFrom(filePath)
	if !FileExist(path) {
		return nil, fmt.Errorf(SsFormat, FileDoesNotExist, filePath)
	}

	directory, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf(SswFormat, FailToReadDirectory, filePath, err)
	}

	files := make(map[string]struct{})
	for _, file := range directory {
		if !file.IsDir() {
			files[file.Name()] = struct{}{}
		}
	}
	return files, nil
}
