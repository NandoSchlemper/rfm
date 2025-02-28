package services

import (
	"rfmtransportes/internal/database"
	m "rfmtransportes/internal/models"
)

type DriverService struct{}

func (d *DriverService) Create(
	dto m.DriverDTO,
	db *database.NeonDB,
) error {
	conn := db.Connection()
	driver := m.Driver{
		Name: dto.Name,
		Cpf:  dto.Cpf,
	}

	result := conn.Create(&driver)
	return result.Error
}
