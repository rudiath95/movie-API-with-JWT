package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rudiath95/movie-API-with-JWT/ini"
)

func init() {
	ini.LoadEnvVariables()
	ini.ConnecttoDB()
	// initializers.SyncDatabases()
}

func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
