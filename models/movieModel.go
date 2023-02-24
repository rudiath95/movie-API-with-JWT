package models

import (
	"time"

	"gorm.io/gorm"
)

type Director struct {
	gorm.Model
	Name  string  `gorm:"not null"`
	Movie []Movie `gorm:"foreignkey:Director_ID"`
}

type Movie struct {
	gorm.Model
	Director_ID uint
	Director    Director
	Title       string
	Price       uint `gorm:"default:0"`
	Owned       bool `gorm:"default:false"`
	YearRelease time.Time
	Genres      []Genre `gorm:"foreignkey:Movie_ID"`
}

type Genre struct {
	gorm.Model
	Name     string
	Movie_ID uint
}
