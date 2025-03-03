package models

import (
	"gorm.io/gorm"
)

type Driver struct {
	gorm.Model
	Name string
	Cpf  string `gorm:"unique"`
	User *User
}

type DriverDTO struct {
	Name string
	Cpf  string
	User *User
}
