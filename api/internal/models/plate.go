package models

import "gorm.io/gorm"

type Plate struct {
	gorm.Model
	Plate       string
	Description string
	Mark        string
	User        *User
}

type PlateDTO struct {
	Plate       string
	Descriptiom string
	Mark        string
	User        *User
}
