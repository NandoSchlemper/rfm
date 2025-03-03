package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Load struct {
	gorm.Model
	Frete       decimal.Decimal
	Description string
	User        *User
}

type LoadDTO struct {
	Frete       decimal.Decimal
	Description string
	User        *User
}
