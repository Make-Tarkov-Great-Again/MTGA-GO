package main

import (
	"MT-GO/database"
	"MT-GO/tools"
	"fmt"
	"github.com/goccy/go-json"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	MTGOUserMods = "%s\"MT-GO/user/mods/%s\""
	//MTGO_SERVER    = "\"MT-GO/server\""
	ModNameMod        = "%s.Mod(\"%s\")"
	BundlesToLoad     = "var bundlesToLoad = []string{%s,\n}"
	BundlesToLoadLoop = "for _, path := range bundlesToLoad {\n\t\tformattedPath := strings.Replace(path, \"\\\\\\\\\", \"\\\\\", -1)\n\t\tdatabase.AddModBundleDirPath(formattedPath)\n\t}"
)

func main() {
	// Get the path of the "mods" folder in the same directory as the executable.

	wd, err := os.Getwd()

	modDir := filepath.Join(wd, "user", "mods")
	fmt.Println("Mod directory:", modDir)
	if !tools.FileExist(modDir) {
		if err := tools.CreateDirectory(modDir); err != nil {
			log.Fatalln(err)
		}
	}

	// List all subdirectories in the "mods" folder.
	modSubDirs, err := tools.GetDirectoriesFrom(modDir)
	if err != nil {
		fmt.Println("Error listing subdirectories in the 'mods' folder:", err)
		return
	}

	// Create an array to store the mod imports and function calls.
	imports := []string{"\"fmt\"", "\"time\""}
	calls := make([]string, 0)
	variables := make([]string, 0)
	bundlesToLoad := make([]string, 0)

	//var modAdvanced []string
	var modConfig *database.ModInfo

	var bundleLoader bool

	if len(modSubDirs) != 0 {
		for name := range modSubDirs {
			fmt.Println("Checking directory:", name)

			// Check if there's a "mod-info.json" file in the subdirectory.
			ModInfoPath := filepath.Join(modDir, name, "mod-info.json")
			if !tools.FileExist(ModInfoPath) {
				fmt.Println("Did not find 'mod-info.json' in:", name, ", continuing...")
				continue
			}

			// Read and parse the mod-info.json file.
			data := tools.GetJSONRawMessage(ModInfoPath)

			if err := json.Unmarshal(data, &modConfig); err != nil {
				fmt.Printf("Error parsing mod-info.json in %s: %v\n", name, err)
				continue
			}
			// Construct the mod import and function call with alias.

			dir := filepath.Join(modDir, name)
			if tools.FileExist(filepath.Join(dir, "bundles")) {
				if !bundleLoader {
					imports = append(imports, "\"MT-GO/database\"", "\"strings\"")
					calls = append(calls, BundlesToLoadLoop)

					bundleLoader = true
				}

				//TODO: See if we can make this better because golly-fuckin-gee

				bundleName := filepath.Join(dir, "bundles")
				fixed := "\n\t\"" + strings.Replace(bundleName, "\\", "\\\\", -1) + "\""
				bundlesToLoad = append(bundlesToLoad, fixed)
				fmt.Println()
			}

			var modImport string
			var modCall string
			if modConfig.PackageAlias == modConfig.PackageName {
				modImport = fmt.Sprintf(MTGOUserMods, "", modConfig.PackageName)
				modCall = fmt.Sprintf(ModNameMod, modConfig.PackageName, strings.Replace(dir, "\\", "\\\\", -1))
			} else {
				modImport = fmt.Sprintf(MTGOUserMods, modConfig.PackageAlias, modConfig.PackageName)
				modCall = fmt.Sprintf(ModNameMod, modConfig.PackageAlias, strings.Replace(dir, "\\", "\\\\", -1))
			}

			imports = append(imports, modImport)
			calls = append(calls, modCall)
		}

		if len(bundlesToLoad) != 0 {
			bundles := strings.Join(bundlesToLoad, ",")
			bundlesVariable := fmt.Sprintf(BundlesToLoad, bundles)
			variables = append(variables, bundlesVariable)
		}
	}

	// Update the "mods.go" file.
	modFile := filepath.Join(modDir, "mods.go")
	fmt.Println("Updating 'mods.go' file:", modFile)
	err = updateModsFile(modFile, imports, variables, calls)
	if err != nil {
		fmt.Println("Error updating mods.go:", err)
		return
	}

	exeDir, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting the executable path:", err)
		return
	}

	// Start the main Go instance in a new cmd instance.
	mainGoFile := filepath.Join(filepath.Dir(exeDir), "server.go")
	fmt.Println("Starting main Go instance:", mainGoFile)
	cmd := exec.Command("cmd", "/c", "start", "cmd", "/k", "go run "+"server.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the new cmd instance to start the main Go program.
	if err := cmd.Run(); err != nil {
		fmt.Println("Error starting the main Go instance:", err)
		return
	}

	// Close the updater executable.
	fmt.Println("Updater finished.")
	os.Exit(1)

}

func updateModsFile(filePath string, imports []string, variables []string, calls []string) error {
	// Create the new content with updated imports and function calls.

	newContent := []byte(fmt.Sprintf(
		`package mods

//TODO: DO NOT DELETE OR MANUALLY EDIT THIS FILE
// This file is automatically generated for mod functionality

import (
	%s
)

%s

func Init() {
	startTime := time.Now()

	%s
        
	endTime := time.Now()
	fmt.Printf("\n[MOD LOADER : COMPLETE] in %%s\n", endTime.Sub(startTime))
}`,
		strings.Join(imports, "\n\t"),
		strings.Join(variables, "\n"),
		strings.Join(calls, "\n\t"),
	))

	// Write the updated content to "mods.go" or create the file if it doesn't exist.
	err := os.WriteFile(filePath, newContent, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
