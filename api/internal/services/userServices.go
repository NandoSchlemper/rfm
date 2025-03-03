package services

import (
	"fmt"

	"gorm.io/gorm"

	"rfmtransportes/internal/models"
)

type UserService struct {
	DB *gorm.DB
}

func (u *UserService) Create(dto *models.UserDTO) error {
	requestBody := &models.UserDTO{Username: dto.Username, Password: dto.Password}

	searchUser := u.DB.First(&requestBody)
	if searchUser.Error != nil {
		return fmt.Errorf("Usuário já cadastrado no sistema!\n%s", searchUser.Error)
	}

	createUser := u.DB.Create(&requestBody)
	if createUser.Error != nil {
		return fmt.Errorf("Erro ao registrar o usuário\n%s", createUser.Error)
	}

	fmt.Printf("Usuário registrado com sucesso!")
	return nil
}
