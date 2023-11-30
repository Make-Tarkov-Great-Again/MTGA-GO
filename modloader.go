package main

import (
	"MT-GO/data"
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
	ModNameMod        = "%s.Mod()"
	BundlesToLoad     = "var bundlesToLoad = []string{%s,\n}"
	BundlesToLoadLoop = "for _, path := range bundlesToLoad {\n\t\tformattedPath := strings.Replace(path, \"\\\\\\\\\", \"\\\\\", -1)\n\t\tdata.AddModBundleDirPath(formattedPath)\n\t}"
)

func main() {
	// Get the path of the "mods" folder in the same directory as the executable.

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	user := filepath.Join(wd, "user")

	if !tools.FileExist(user) {
		if err := tools.CreateDirectory(user); err != nil {
			log.Fatalln(err)
		}
	}

	profilesPath := filepath.Join(user, "profiles")
	if !tools.FileExist(profilesPath) {
		if err := tools.CreateDirectory(profilesPath); err != nil {
			log.Fatalln(err)
		}
	}

	modDir := filepath.Join(user, "mods")
	log.Println("Mod directory:", modDir)
	if !tools.FileExist(modDir) {
		if err := tools.CreateDirectory(modDir); err != nil {
			log.Fatalln(err)
		}
	}

	// List all subdirectories in the "mods" folder.
	modSubDirs, err := tools.GetDirectoriesFrom(modDir)
	if err != nil {
		log.Fatalln("Error listing subdirectories in the 'mods' folder:", err)
	}

	// Create an array to store the mod imports and function calls.
	imports := []string{"\"fmt\"", "\"time\""}
	calls := make([]string, 0)
	variables := make([]string, 0)
	bundlesToLoad := make([]string, 0)

	//var modAdvanced []string
	var modConfig *data.ModInfo

	var bundleLoader bool

	if len(modSubDirs) != 0 {
		for name := range modSubDirs {
			log.Println("Checking directory:", name)

			// Check if there's a "mod-info.json" file in the subdirectory.
			ModInfoPath := filepath.Join(modDir, name, "mod-info.json")
			if !tools.FileExist(ModInfoPath) {
				log.Println("Did not find 'mod-info.json' in:", name, ", continuing...")
				continue
			}

			// Read and parse the mod-info.json file.
			input := tools.GetJSONRawMessage(ModInfoPath)

			if err := json.Unmarshal(input, &modConfig); err != nil {
				fmt.Printf("Error parsing mod-info.json in %s: %v\n", name, err)
				continue
			}
			// Construct the mod import and function call with alias.

			dir := filepath.Join(modDir, name)
			if tools.FileExist(filepath.Join(dir, "bundles")) {
				if !bundleLoader {
					imports = append(imports, "\"MT-GO/data\"", "\"strings\"")
					calls = append(calls, BundlesToLoadLoop)

					bundleLoader = true
				}

				//TODO: See if we can make this better because golly-fuckin-gee

				bundleName := filepath.Join(dir, "bundles")
				fixed := "\n\t\"" + strings.Replace(bundleName, "\\", "\\\\", -1) + "\""
				bundlesToLoad = append(bundlesToLoad, fixed)
				log.Println()
			}

			var modImport string
			var modCall string
			if modConfig.PackageAlias == modConfig.PackageName {
				modImport = fmt.Sprintf(MTGOUserMods, "", modConfig.PackageName)
				modCall = fmt.Sprintf(ModNameMod, modConfig.PackageName)
			} else {
				modImport = fmt.Sprintf(MTGOUserMods, modConfig.PackageAlias, modConfig.PackageName)
				modCall = fmt.Sprintf(ModNameMod, modConfig.PackageAlias)
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
	log.Println("Updating 'mods.go' file:", modFile)
	if err := updateModsFile(modFile, imports, variables, calls); err != nil {
		log.Fatalln("Error updating mods.go:", err)
	}

	exeDir, err := os.Executable()
	if err != nil {
		log.Fatalln("Error getting the executable path:", err)
	}

	// Start the main Go instance in a new cmd instance.
	mainGoFile := filepath.Join(filepath.Dir(exeDir), "backend.go")
	log.Println("Starting main Go instance:", mainGoFile)
	cmd := exec.Command("cmd", "/c", "start", "cmd", "/k", "go run backend.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the new cmd instance to start the main Go program.
	if err := cmd.Run(); err != nil {
		log.Println("Error starting the main Go instance:", err)
		return
	}

	// Close the updater executable.
	log.Println("Updater finished.")
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
