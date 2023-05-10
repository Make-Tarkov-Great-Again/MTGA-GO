package tools

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

// WriteToFile writes the given string of data to the specified file path
func WriteToFile(filePath string, data string) error {
	path := GetAbsolutePathFrom(filePath)
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(data)
	if err != nil {
		return err
	}

	err = writer.Flush()
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
	if FileExist(filePath) {
		log.Print("File already exists")
		return nil
	}
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

// ReadFile reads the file at filePath and returns its contents as a byte slice.
func ReadFile(filePath string) ([]byte, error) {
	path := GetAbsolutePathFrom(filePath)

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return data, nil
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
