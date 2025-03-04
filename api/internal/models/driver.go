package models

import (
	"gorm.io/gorm"
)

type Driver struct {
	gorm.Model
	Name   string
	Cpf    string `gorm:"unique"`
	UserID int
	User   User `gorm:"foreignkey:UserID;references:ID"`
}

type DriverDTO struct {
	Name   string
	Cpf    string
	UserID int
	User   User `gorm:"foreignkey:UserID;references:ID"`
}
