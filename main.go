package main

import (
	"MT-GO/database"
)

func main() {

	database.InitializeDatabase()
	/* 	dbErr := database.InitializeDatabase()
	   	if dbErr != nil {
	   		log.Fatalf("error initializing database: %v", dbErr)
	   	} */
}
