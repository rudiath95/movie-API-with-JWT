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
	//USER
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequiredAuth, controllers.Validate)
	r.POST("/profile", middleware.RequiredAuth, controllers.UserInfo)
	r.PUT("/profile/:id", middleware.RequiredAuth, controllers.UpdateUserInfo)
	r.PUT("/charge_balance", middleware.RequiredAuth, controllers.ChargeBalance)

	//VOUCHER
	r.POST("/voucher", middleware.RequiredAuth, controllers.AddVoucher)

	r.PUT("/redeem", middleware.RequiredAuth, controllers.RedeemVoucher)

	r.Run()
}
