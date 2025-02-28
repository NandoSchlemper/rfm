package models

import (
	"gorm.io/gorm"
)

type Driver struct {
	gorm.Model
	Name string
	Cpf  string
}

type DriverDTO struct {
	Name string
	Cpf  string
}
