package main

import (
	"MT-GO/database"
)

func main() {

	database.LoadCore()
	/* 	dbErr := database.InitializeDatabase()
	   	if dbErr != nil {
	   		log.Fatalf("error initializing database: %v", dbErr)
	   	} */
}
