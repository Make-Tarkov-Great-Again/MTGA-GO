package tools

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/goccy/go-json"
)

const (
	FILE_DOES_NOT_EXIST      string = "file does not exist"
	FAIL_TO_READ_FILE        string = "failed to read file"
	DIRECTORY_EXISTS         string = "directory already exists"
	DIRECTORY_DOES_NOT_EXIST string = "directory does not exist"
	FAIL_TO_READ_DIRECTORY   string = "failed to read directory"
)

const (
	SSW_FORMAT string = "%s: %s: %w"
	SS_FORMAT  string = "%s: %s"
)

// WriteToFile writes the given string of data to the specified file path
func WriteToFile(filePath string, data any) error {
	path := GetAbsolutePathFrom(filePath)
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
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

	err = encoder.Encode(data)
	if err != nil {
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
		return fmt.Errorf(SS_FORMAT, DIRECTORY_EXISTS, filePath)
	}

	err := os.MkdirAll(path, 0755)
	if err != nil {
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
		return nil, fmt.Errorf(SS_FORMAT, FILE_DOES_NOT_EXIST, filePath)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf(SSW_FORMAT, FAIL_TO_READ_FILE, filePath, err)
	}

	return data, nil
}

// GetDirectoriesFrom returns a list of directories from a file path
func GetDirectoriesFrom(filePath string) (map[string]*struct{}, error) {
	path := GetAbsolutePathFrom(filePath)
	if !FileExist(path) {
		return nil, fmt.Errorf(SS_FORMAT, FILE_DOES_NOT_EXIST, filePath)
	}

	directory, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf(SSW_FORMAT, FAIL_TO_READ_DIRECTORY, filePath, err)
	}

	files := make(map[string]*struct{})
	for _, file := range directory {
		if file.IsDir() {
			files[file.Name()] = nil
		}
	}
	return files, nil
}

// GetFilesFrom returns a list of files from a file path
func GetFilesFrom(filePath string) (map[string]*struct{}, error) {
	path := GetAbsolutePathFrom(filePath)
	if !FileExist(path) {
		return nil, fmt.Errorf(SS_FORMAT, FILE_DOES_NOT_EXIST, filePath)
	}

	directory, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf(SSW_FORMAT, FAIL_TO_READ_DIRECTORY, filePath, err)
	}

	files := make(map[string]*struct{})
	for _, file := range directory {
		if !file.IsDir() {
			files[file.Name()] = nil
		}
	}
	return files, nil
}

func TransformInterfaceIntoMappedArray(data []any) []map[string]any {
	results := make([]map[string]any, 0, len(data))
	for _, v := range data {
		result := v.(map[string]any)
		results = append(results, result)
	}
	return results
}

func TransformInterfaceIntoMappedObject(data any) map[string]any {
	result := data.(map[string]any)
	return result
}

func AuditArrayCapacity(data []map[string]any) []map[string]any {
	dataLen := len(data)
	results := make([]map[string]any, 0, dataLen)
	for i := 0; i < dataLen; i++ {
		result := data[i]
		results = append(results, result)
	}
	return results
}
