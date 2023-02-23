package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rudiath95/movie-API-with-JWT/controllers"
	"github.com/rudiath95/movie-API-with-JWT/ini"
	"github.com/rudiath95/movie-API-with-JWT/middleware"
)

func init() {
	ini.LoadEnvVariables()
	ini.ConnecttoDB()
	ini.SyncDatabases()
}

func main() {

	r := gin.Default()

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequiredAuth, controllers.Validate)
	r.POST("/profile", middleware.RequiredAuth, controllers.UserInfo)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
