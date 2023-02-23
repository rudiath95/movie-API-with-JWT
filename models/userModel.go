package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	User_ID  uuid.UUID
	Username string `gorm:"unique"`
	Password string
}

type UserInfo struct {
	gorm.Model
	User_ID        User   `gorm:"embedded"`
	Email          string `gorm:"unique"`
	FirstName      string
	LastName       string
	CurrentBalance uint `gorm:"default:0"`
}

type VouceherList struct {
	gorm.Model
	VoucherCode   string `gorm:"unique"`
	VoucherAmount uint
	VoucherStatus bool `gorm:"default:true"`
}

type VoucherRedeem struct {
	gorm.Model

	User_ID  User `gorm:"embedded"`
	Username User `gorm:"embedded"`

	RedeemedVoucher string

	VoucherCode   VouceherList `gorm:"embedded"`
	VoucherAmount VouceherList `gorm:"embedded"`
	VoucherStatus VouceherList `gorm:"embedded"`

	CurrentBalance UserInfo `gorm:"embedded"`
}
