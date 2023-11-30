// Package main is a package declaration

package main

import (
	"MT-GO/cli"
	"MT-GO/mods"
	"MT-GO/server"
	"fmt"
	"time"

	"MT-GO/data"
)

func main() {
	startTime := time.Now()
	data.SetPrimaryDatabase()

	mods.Init()
	data.LoadBundleManifests()
	data.LoadCustomItems()

	data.SetCache()

	endTime := time.Now()
	fmt.Printf("Database initialized in %s\n\n", endTime.Sub(startTime))

	server.Start()
	cli.Start()
}
