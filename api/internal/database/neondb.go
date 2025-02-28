package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type NeonDB struct {
	URI string
}

func NewNeonDB() *NeonDB {
	return &NeonDB{
		URI: os.Getenv("DATABASE_URL"),
	}
}

func (database *NeonDB) Connection() *gorm.DB {
	db, err := gorm.Open(postgres.Open(database.URI), &gorm.Config{})
	if err != nil {
		fmt.Print("Erro ao conectar no DB")
	}
	return db
}
