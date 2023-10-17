package database

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	serverFolder := filepath.Dir(wd)
	tempFolder := "AppData/temp" // Adjust the path as needed
	assetsFolder := "assets"     // Adjust the path as needed

	// Create the temporary folder
	err = os.MkdirAll(tempFolder, 0755)
	if err != nil {
		fmt.Printf("Error creating temporary folder: %v\n", err)
		return
	}

	// Copy the entire assets folder to the temporary folder
	err = copyFolder(filepath.Join(serverFolder, assetsFolder), filepath.Join(tempFolder, assetsFolder))
	if err != nil {
		fmt.Printf("Error copying assets folder: %v\n", err)
		return
	}

	// Now mods can freely work with the contents of the temporary assets folder
	// ...

	// Database initialization using files from the temporary assets folder
	// Clean up the temporary folder after initialization if needed
	// ...
	// Define the path to the "mods" folder
	modsFolder := "mods" // Adjust the path as needed

	// Get a list of files in the "mods" folder
	files, err := getJSFiles(modsFolder)
	if err != nil {
		fmt.Printf("Error listing files in the 'mods' folder: %v\n", err)
		return
	}

	for _, jsFilePath := range files {
		// Execute the JavaScript code for each mod using Node.js
		cmd := exec.Command("node", jsFilePath)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin

		// Run the JavaScript code for the mod
		if err := cmd.Run(); err != nil {
			fmt.Printf("Error running the JavaScript code for %s: %v\n", jsFilePath, err)
			continue // Continue to the next mod
		}

		fmt.Printf("JavaScript code for %s executed successfully.\n", jsFilePath)
	}
}

func copyFolder(src, dst string) error {
	sourceInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	if err = os.MkdirAll(dst, sourceInfo.Mode()); err != nil {
		return err
	}

	directory, err := os.Open(src)
	if err != nil {
		return err
	}

	objects, err := directory.Readdir(-1)
	for _, obj := range objects {
		sourceFilePath := filepath.Join(src, obj.Name())
		destinationFilePath := filepath.Join(dst, obj.Name())

		if obj.IsDir() {
			err = copyFolder(sourceFilePath, destinationFilePath)
			if err != nil {
				return err
			}
		} else {
			err = copyFile(sourceFilePath, destinationFilePath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Function to copy a file
func copyFile(src, dst string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}

// Function to get a list of JavaScript files in a folder
func getJSFiles(folderPath string) ([]string, error) {
	var jsFiles []string

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".js" {
			jsFiles = append(jsFiles, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return jsFiles, nil
}
