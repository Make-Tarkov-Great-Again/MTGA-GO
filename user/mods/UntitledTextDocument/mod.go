package UntitledTextDocument

import (
	"MT-GO/database"
	"MT-GO/tools"
	"fmt"
	"github.com/goccy/go-json"
	"path/filepath"
	"runtime"
	"strings"
)

var modInfo = SetModConfig()

func Mod() {
	Load()
}

/* ---------------------- Boring mod bindings below lol --------------------- */
//TODO: Save directory path to reduce imports

func SetModConfig() database.ModInfo {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic(fmt.Errorf("failed to get the current file's path"))
	}

	directory := filepath.Dir(filename)
	readFile, err := tools.ReadFile(filepath.Join(directory, "mod-info.json"))
	if err != nil {
		panic(err)
	}

	config := new(database.ModInfo)
	err = json.Unmarshal(readFile, &config)
	if err != nil {
		panic(err)
	}

	config.Dir = directory
	return *config
}

var bundlesDir string

func Load() {
	bundlesDir = filepath.Join(modInfo.Dir, "bundles", "test9")

	addToAssort(filepath.Join(bundlesDir, "traders", "assort.json"))
	addToLocales(filepath.Join(bundlesDir, "locales"))
	addToHandbook(filepath.Join(bundlesDir, "templates", "handbook", "test9.json"))
	addToItems(filepath.Join(bundlesDir, "templates", "items", "test9.json"))
}

func addToAssort(path string) {
	itemAssort := make(map[string]database.Assort)
	data := tools.GetJSONRawMessage(path)
	if err := json.Unmarshal(data, &itemAssort); err != nil {
		fmt.Println(err)
		return
	}

	for tid, assort := range itemAssort {
		traderAssort := database.GetTraderByUID(tid).Assort

		for _, item := range assort.Items {
			traderAssort.Items = append(traderAssort.Items, item)
		}

		for id, scheme := range assort.BarterScheme {
			traderAssort.BarterScheme[id] = scheme
		}

		for id, level := range assort.LoyalLevelItems {
			traderAssort.LoyalLevelItems[id] = level
		}
	}
}

type locale struct {
	Name        string `json:"Name"`
	ShortName   string `json:"ShortName"`
	Description string `json:"Description"`
}

func addToLocales(path string) {
	files, err := tools.GetFilesFrom(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	for file := range files {
		templates := map[string]locale{}
		data := tools.GetJSONRawMessage(filepath.Join(path, file))
		if err := json.Unmarshal(data, &templates); err != nil {
			fmt.Println(err)
			return
		}

		localeName := strings.TrimSuffix(file, ".json")
		locale := database.GetLocalesLocaleByName(localeName)

		for key, value := range templates {
			name := fmt.Sprintf("%s %s", key, "Name")
			locale[name] = value.Name

			shortName := fmt.Sprintf("%s %s", key, "ShortName")
			locale[shortName] = value.ShortName

			description := fmt.Sprintf("%s %s", key, "Description")
			locale[description] = value.Description
		}
	}

}

func addToHandbook(path string) {
	data := tools.GetJSONRawMessage(path)

	entry := new(database.HandbookItem)
	if err := json.Unmarshal(data, &entry); err != nil {
		fmt.Println(err)
		return
	}

	database.SetHandbookItemEntry(*entry)
}

func addToItems(path string) {
	data := tools.GetJSONRawMessage(path)

	entry := new(database.DatabaseItem)
	if err := json.Unmarshal(data, &entry); err != nil {
		fmt.Println(err)
		return
	}

	database.SetNewItem(*entry)
}

//TODO: Attempt parallelism later
/*
var wg sync.WaitGroup
var completionCh = make(chan struct{})
var workerCh = make(chan struct{}, tools.CalculateWorkers()/4)

var startTime time.Time

var tasks = []func(){
	bundles,
}

func Load(passed *database.ModInfo) {
	startTime = time.Now()

	for _, task := range tasks {
		wg.Add(1)
		go func(taskFunc func()) {
			defer wg.Done()
			workerCh <- struct{}{}
			taskFunc()
			<-workerCh
			completionCh <- struct{}{}
		}(task)
	}

	go func() {
		wg.Wait()
		close(completionCh)
	}()

	for range tasks {
		<-completionCh
	}

	endTime := time.Now()
	fmt.Printf("\n\nThis shit initialized in %s with %d workers\n", endTime.Sub(startTime), len(workerCh))

}
*/
