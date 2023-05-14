package main

import (
	"compress/zlib"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var pl = fmt.Println

func main() {

	err := InitializeDatabase()
	if err != nil {
		pl("error initialize database: %w", err)
	}

	r := gin.New()
	mtga := r.Group("/")
	//mtga.Use(jsonContentTypeParser())

	mtga.GET("/serverConfig", func(c *gin.Context) {
		c.JSON(http.StatusOK, Database.core.serverConfig)
	})

	mtga.GET("/clientSettings", func(c *gin.Context) {
		c.JSON(http.StatusOK, Database.core.clientSettings)
	})

	mtga.GET("/globals", func(c *gin.Context) {
		c.JSON(http.StatusOK, Database.core.globals)
	})

	port := strconv.FormatFloat(Database.core.serverConfig["port"].(float64), 'f', -1, 64)
	ipport := fmt.Sprintf("%s:%s", Database.core.serverConfig["ip"].(string), port)
	r.Run(ipport)
}

// jsonContentTypeParser parses the body of a request and sets it to the context.
func jsonContentTypeParser() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "GET" {
			c.Next()
		}

		userAgent := c.Request.Header.Get("User-Agent")

		if strings.Contains(userAgent, "Unity") {
			err := parseUnityBody(c)
			if err != nil {
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}
		} else {
			err := parseNormalBody(c)
			if err != nil {
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}
		}

		c.Next()
	}
}

// parseUnityBody parses the body of a request from the Unity client.
func parseUnityBody(c *gin.Context) error {
	reader, err := zlib.NewReader(c.Request.Body)
	if err != nil {
		return fmt.Errorf("error zlib reader: %w", err)
	}
	defer reader.Close()

	bufferString, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("error read body: %w", err)
	}
	c.Set("body", bufferString)
	return nil

	/* 	if len(bufferString) > 0 {
	   		unityBody, err := tools.ParseJSON(&bufferString)
	   		if err != nil {
	   			return fmt.Errorf("error parse json: %w", err)
	   		}
	   		c.Set("body", unityBody)
	   	}
	   	return nil */
}

// parseNormalBody parses the body of a request from a browser or other client.
func parseNormalBody(c *gin.Context) error {
	return nil

	/* 	body, err := io.ReadAll(c.Request.Body)
	   	if err != nil {
	   		return fmt.Errorf("error read body: %w", err)
	   	}
	   	defer c.Request.Body.Close()

	   	normalBody, err := tools.ParseJSON(&body)
	   	if err != nil {
	   		return fmt.Errorf("error parse json: %w", err)
	   	}

	   	c.Set("body", normalBody)
	   	return nil */
}
