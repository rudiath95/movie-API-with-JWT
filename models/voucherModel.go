package models

import "gorm.io/gorm"

type VoucherList struct {
	gorm.Model
	VoucherCode   string `gorm:"unique"`
	VoucherAmount uint
	VoucherStatus bool `gorm:"default:true"`
}

type VoucherRedeem struct {
}
