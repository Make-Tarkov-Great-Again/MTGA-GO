package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var pl = fmt.Println

func main() {

	err := InitializeDatabase()
	if err != nil {
		pl("error initialize database: %w", err)
	}

	r := gin.Default()
	r.GET("/serverConfig", func(c *gin.Context) {
		c.JSON(http.StatusOK, Database.core.serverConfig)
	})

	port := strconv.FormatFloat(Database.core.serverConfig["port"].(float64), 'f', -1, 64)
	ipport := fmt.Sprintf("%s:%s", Database.core.serverConfig["ip"].(string), port)
	r.Run(ipport)
}
