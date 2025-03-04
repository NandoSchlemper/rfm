package services

import (
	"fmt"

	"gorm.io/gorm"

	"rfmtransportes/internal/middlewares"
	"rfmtransportes/internal/models"
)

type AuthService struct {
	DB *gorm.DB
}

func (a *AuthService) Register(userDTO *models.User) (string, error) {
	// Verificar se User DB, colcoar user DB, pegar payload, retornar token
	payload := &models.User{Username: userDTO.Username, Password: userDTO.Password}
	validateError := a.DB.First(&payload)
	if validateError != nil {
		return "", validateError.Error
	}

	createError := a.DB.Create(&payload)
	if createError.Error != nil {
		return "", createError.Error
	}

	token, err := middlewares.GenerateToken(payload)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *AuthService) Login(userDTO *models.User) (string, error) {
	// pegar payload, validar,  retornar token
	payload := &models.User{Username: userDTO.Username, Password: userDTO.Password}
	validateError := a.DB.First(&payload)
	if validateError.Error != nil {
		return "", validateError.Error
	}

	if payload.Password != userDTO.Password {
		return "", fmt.Errorf("Senha divergente")
	}

	token, err := middlewares.GenerateToken(payload)
	if err != nil {
		return "", err
	}
	return token, nil
}
