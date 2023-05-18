package main

import (
	"log"
)

func main() {

	dbErr := initializeDatabase()
	if dbErr != nil {
		log.Fatalf("error initializing database: %v", dbErr)
	}

	ginErr := setGin()
	if ginErr != nil {
		log.Fatalf("error setting gin: %v", ginErr)
	}
}
