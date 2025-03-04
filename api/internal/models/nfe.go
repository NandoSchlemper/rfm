package models

import "gorm.io/gorm"

type NFE struct {
	gorm.Model
	Numero      string
	DataEmissão string
	UserID      int
	LoadID      int
	Load        Load `gorm:"foreignkey:LoadID;references:ID"`
	User        User `gorm:"foreignkey:UserID;references:ID"`
}
