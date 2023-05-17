package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {

	dbErr := initializeDatabase()
	if dbErr != nil {
		pl("error initialize database: %w", dbErr)
	}

	ginErr := setGin()
	if ginErr != nil {
		pl("error setting gin: %w", ginErr)
	}
}
