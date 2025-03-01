package models

import (
	"gorm.io/gorm"
)

type Driver struct {
	gorm.Model
	Name string
	Cpf  string `gorm:"unique"`
}

type DriverDTO struct {
	Name string
	Cpf  string
}
