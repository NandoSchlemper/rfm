package services

import (
	"fmt"

	"gorm.io/gorm"

	m "rfmtransportes/internal/models"
)

type DriverService struct {
	DB *gorm.DB
}

func (d *DriverService) CreateDriver(dto m.DriverDTO) error {
	requestBody := m.DriverDTO{Name: dto.Name, Cpf: dto.Cpf, User: dto.User}

	//error := d.DB.First(&requestBody.Cpf)
	//if error != nil {
	//	fmt.Println("Erro ao cadastrar motorista")
	//	return fmt.Errorf("Motorista já cadastrado no sistema")
	//}

	err := d.DB.Create(&requestBody)
	if err != nil {
		fmt.Print("Erro ao cadastrar motorista")
		return fmt.Errorf("Erro ao criar motorista!\n%s", err.Error)
	}

	fmt.Print("Motorista cadastrado com sucesso!")
	return nil
}
