package models

import "gorm.io/gorm"

type Plate struct {
	gorm.Model
	Plate       string
	Description string
	Mark        string
	UserID      int
	User        User `gorm:"foreignkey:UserID;references:ID"`
}

type PlateDTO struct {
	Plate       string
	Descriptiom string
	Mark        string
	UserID      int
	User        User `gorm:"foreignkey:UserID;references:ID"`
}
