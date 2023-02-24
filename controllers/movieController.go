package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rudiath95/movie-API-with-JWT/ini"
	"github.com/rudiath95/movie-API-with-JWT/models"
)

func AddDirector(c *gin.Context) {
	var body struct {
		Name string
	}

	c.Bind(&body)

	//Create Query
	post := models.Director{Name: body.Name}

	result := ini.DB.Create(&post)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to post",
		})
		return
	}

	// Return
	c.JSON(200, gin.H{
		"post": post,
	})

}

func GetDirector(c *gin.Context) {

	//Call DB
	var post []models.Director
	ini.DB.Find(&post)

	//Return
	c.JSON(200, gin.H{
		"post": post,
	})
}

func UpdateDirector(c *gin.Context) {
	//get id
	id := c.Param("id")

	//Call DB
	var body struct {
		Name string
	}

	c.Bind(&body)
	var post models.Director
	ini.DB.Find(&post, id)

	//Update
	ini.DB.Model(&post).Updates(models.Director{
		Name: body.Name,
	})

	//Return
	c.JSON(200, gin.H{
		"post": post,
	})
}

func DeleteDirector(c *gin.Context) {
	//get id
	id := c.Param("id")

	//Delete
	ini.DB.Delete(&models.Director{}, id)

	//Return
	c.JSON(200, gin.H{
		"post": "Deleted",
	})
}
