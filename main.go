package main

import (
	"MT-GO/database"
	"MT-GO/server"
	"log"
)

func main() {

	dbErr := database.InitializeDatabase()
	if dbErr != nil {
		log.Fatalf("error initializing database: %v", dbErr)
	}

	ginErr := server.SetGin()
	if ginErr != nil {
		log.Fatalf("error setting gin: %v", ginErr)
	}
}
