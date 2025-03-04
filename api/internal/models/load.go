package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Load struct {
	gorm.Model
	Frete       decimal.Decimal
	Description string
	UserID      int
	User        User `gorm:"foreignkey:UserID;references:ID"`
}

type LoadDTO struct {
	Frete       decimal.Decimal
	Description string
	UserID      int
	User        User `gorm:"foreignkey:UserID;references:ID"`
}
