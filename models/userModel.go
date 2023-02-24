package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}

type UserInfo struct {
	gorm.Model
	User_ID        int
	User           User   `gorm:"foreignKey:User_ID"`
	Email          string `gorm:"unique"`
	FirstName      string
	LastName       string
	CurrentBalance uint `gorm:"default:0"`
}
