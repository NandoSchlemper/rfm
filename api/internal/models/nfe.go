package models

import "gorm.io/gorm"

type NFE struct {
	gorm.Model
	Numero      string
	DataEmissão string
	Load        *Load
	User        *User
}
