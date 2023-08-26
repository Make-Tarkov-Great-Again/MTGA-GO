package tools

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
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
func WriteToFile(filePath string, data interface{}) error {
	// Get the absolute path of the file
	path := GetAbsolutePathFrom(filePath)

	// Create the directory if it doesn't exist
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return err
	}

	// Create or truncate the file
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a custom JSON encoder that doesn't escape Unicode
	encoder := json.NewEncoder(file)
	encoder.SetEscapeHTML(false)

	// Encode and write the data to the file
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
func GetDirectoriesFrom(filePath string) ([]string, error) {
	path := GetAbsolutePathFrom(filePath)
	if !FileExist(path) {
		return nil, fmt.Errorf(SS_FORMAT, FILE_DOES_NOT_EXIST, filePath)
	}

	directory, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf(SSW_FORMAT, FAIL_TO_READ_DIRECTORY, filePath, err)
	}

	files := make([]string, 0, len(directory))
	for _, file := range directory {
		if file.IsDir() {
			files = append(files, file.Name())
		}
	}
	return files, nil
}

// GetFilesFrom returns a list of files from a file path
func GetFilesFrom(filePath string) ([]string, error) {
	path := GetAbsolutePathFrom(filePath)
	if !FileExist(path) {
		return nil, fmt.Errorf(SS_FORMAT, FILE_DOES_NOT_EXIST, filePath)
	}

	directory, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf(SSW_FORMAT, FAIL_TO_READ_DIRECTORY, filePath, err)
	}

	files := make([]string, 0, len(directory))
	for _, file := range directory {
		if !file.IsDir() {
			files = append(files, file.Name())
		}
	}
	return files, nil
}

func TransformInterfaceIntoMappedArray(data []interface{}) []map[string]interface{} {
	results := make([]map[string]interface{}, 0, len(data))
	for _, v := range data {
		result := v.(map[string]interface{})
		results = append(results, result)
	}
	return results
}

func TransformInterfaceIntoMappedObject(data interface{}) map[string]interface{} {
	result := data.(map[string]interface{})
	return result
}

func AuditArrayCapacity(data []map[string]interface{}) []map[string]interface{} {
	dataLen := len(data)
	results := make([]map[string]interface{}, 0, dataLen)
	for i := 0; i < dataLen; i++ {
		result := data[i]
		results = append(results, result)
	}
	return results
}
