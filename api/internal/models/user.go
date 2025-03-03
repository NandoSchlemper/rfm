package models

import (
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
}

type UserDTO struct {
	Username string
	Password string
}

type UserJwt struct {
	User *User
	jwt.RegisteredClaims
}
