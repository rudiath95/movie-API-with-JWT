package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/rudiath95/movie-API-with-JWT/ini"
	"github.com/rudiath95/movie-API-with-JWT/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	//Get the email/pass off req body
	var body struct {
		Username string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	//HashPassword
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	//Create the user
	user := models.User{Username: body.Username, Password: string(hash)}
	result := ini.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user or Email already used",
		})
		return
	}

	//Respond
	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {
	//Get the email/pass off req body
	var body struct {
		Username string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	//Look up requested user
	var user models.User
	ini.DB.First(&user, "username = ?", body.Username)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	//Compare sent in pass with saved user pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	//Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	//Sign and get the token
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fail to Create Token",
		})
		return
	}

	//sent it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30 /*path*/, "" /*domain_name*/, "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	//example for argument
	//user.(models.User).Email = "xxx"

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func UserInfo(c *gin.Context) {
	user, _ := c.Get("user")

	var body struct {
		User_ID   int
		Email     string
		FirstName string
		LastName  string
	}

	c.Bind(&body)

	//Create Query
	post := models.UserInfo{User_ID: int(user.(models.User).ID), Email: body.Email, FirstName: body.FirstName, LastName: body.LastName}

	result := ini.DB.Create(&post)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to post Userinfo",
		})
		return
	}

	//example for argument
	//user.(models.User).Email = "xxx"

	//Return
	c.JSON(200, gin.H{
		"post": post,
	})
}

func UpdateUserInfo(c *gin.Context) {
	//get id
	id := c.Param("id")

	//Call DB
	var body struct {
		Email     string
		FirstName string
		LastName  string
	}

	c.Bind(&body)
	var post models.UserInfo
	ini.DB.Find(&post, id)

	//Update
	ini.DB.Model(&post).Updates(models.UserInfo{
		Email:     body.Email,
		FirstName: body.FirstName,
		LastName:  body.LastName,
	})

	//Return
	c.JSON(200, gin.H{
		"post": post,
	})
}

func ChargeBalance(c *gin.Context) {
	//get id
	user, _ := c.Get("user")
	id := user.(models.User).ID

	//Call DB
	var body struct {
		CurrentBalance uint
	}

	c.Bind(&body)
	var post models.UserInfo
	ini.DB.Raw("SELECT current_balance FROM user_infos WHERE user_id = ?", id).Scan(&post)
	// ini.DB.Find(&post, 2)

	total := post.CurrentBalance + body.CurrentBalance
	ini.DB.Raw("UPDATE user_infos SET current_balance = ? WHERE user_id = ? ",
		total, id).Scan(&post)

	//Return
	c.JSON(200, gin.H{
		"current_balance": total,
	})
}
