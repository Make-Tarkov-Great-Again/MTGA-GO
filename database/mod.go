package database

import (
	"MT-GO/tools"
	"fmt"
	"github.com/goccy/go-json"
	"path/filepath"
	"strings"
)

type ModInfo struct {
	PackageName  string
	PackageAlias string
	Dir          string
	Advanced     struct {
		CustomRoutes bool
	}
	Config map[string]interface{}
}

type Manifests struct {
	Manifests []*Manifest `json:"manifest"`
}

type Manifest struct {
	ModPath        string   `json:"modPath"`
	Key            string   `json:"key"`
	Path           string   `json:"path"`
	FilePath       string   `json:"filePath,omitempty"`
	DependencyKeys []string `json:"dependencyKeys"`
}

var modBundleDirPaths = make([]string, 0)

var bundleManifests []*Manifest

func (m *ModInfo) GetConfig() map[string]interface{} {
	return m.Config
}

func GetBundleManifests() []*Manifest {
	return bundleManifests
}

func ClearBundleManifests() {
	bundleManifests = nil
}

func AddModBundleDirPath(modBundleDirPath string) {
	modBundleDirPaths = append(modBundleDirPaths, modBundleDirPath)
}

func SetBundleManifests() {
	if len(modBundleDirPaths) == 0 {
		return
	}
	defer func() {
		modBundleDirPaths = nil
	}()

	isLocal := GetServerConfig().Hostname == "localhost"
	var mainAddress string
	if !isLocal {
		mainAddress = GetMainAddress()
	}

	for _, path := range modBundleDirPaths {
		bundlesJsonPath := filepath.Join(path, "bundles.json")

		var err error
		if !tools.FileExist(bundlesJsonPath) {
			err = fmt.Errorf("bundles.json file not located in %s, returning", path)
			fmt.Println(err)
			return
		}

		bundlesSubDirectories, err := tools.GetDirectoriesFrom(path)
		if err != nil {
			fmt.Println(err)
			return
		}

		manifests := new(Manifests)
		data := tools.GetJSONRawMessage(bundlesJsonPath)
		if err := json.Unmarshal(data, &manifests); err != nil {
			fmt.Println(err)
			return
		}

		for _, manifest := range manifests.Manifests {
			name := strings.Split(manifest.Key, ".")[0]
			if _, ok := bundlesSubDirectories[name]; !ok {
				err = fmt.Errorf("bundle %s does not have a directory in %s", name, path)
				fmt.Println(err)
				return
			}

			bundleDirPath := filepath.Join(path, name)
			bundlePath := filepath.Join(bundleDirPath, manifest.Key)
			if !tools.FileExist(bundlePath) {
				err = fmt.Errorf("bundle %s does not exist in %s", manifest.Key, bundleDirPath)
				fmt.Println(err)
				return
			}

			manifest.ModPath = bundleDirPath
			if isLocal {
				manifest.Path = bundlePath
			} else {
				manifest.Path = filepath.Join(mainAddress, "files", "bundle", manifest.Key)
				manifest.FilePath = manifest.Path
			}

			bundleManifests = append(bundleManifests, manifest)
		}
	}
}
