package tools

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

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
	err := os.MkdirAll(path, 0755)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

// FileExist checks if a file exists at the specified path
func FileExist(filePath string) bool {
	path := GetAbsolutePathFrom(filePath)
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err)
}

func Stringify(data interface{}, oneline bool) string {
	var bytes, err = []byte{}, error(nil)
	if oneline {
		bytes, err = json.Marshal(data)
	} else {
		bytes, err = json.MarshalIndent(data, "", "\t")
	}
	if err != nil {
		return ""
	}
	return string(bytes[:])
}

// ReadParsed reads a file path and parses it into an interface
func ReadParsed(filePath string) (interface{}, error) {
	path := GetAbsolutePathFrom(filePath)

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var result interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetDirectoriesFrom returns a list of directories from a file path
func GetDirectoriesFrom(filePath string) []string {
	path := GetAbsolutePathFrom(filePath)
	if !FileExist(path) {
		log.Fatal("File does not exist")
		return nil
	}

	files := []string{}
	directory, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range directory {
		if file.IsDir() {
			files = append(files, file.Name())
		}
	}
	return files
}

// GetFilesFrom returns a list of files from a file path
func GetFilesFrom(filePath string) []string {
	path := GetAbsolutePathFrom(filePath)
	if !FileExist(path) {
		log.Fatal("File does not exist")
		return nil
	}

	files := []string{}
	directory, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range directory {
		if !file.IsDir() {
			files = append(files, file.Name())
		}
	}
	return files
}
