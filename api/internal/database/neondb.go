package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	m "rfmtransportes/internal/models"
)

type NeonDB struct{}

func (n *NeonDB) Migrate() error {
	db := n.Connection()

	driver := m.Driver{}

	err := db.AutoMigrate(
		&driver,
	)
	if err != nil {
		return fmt.Errorf("Erro ao migrar o banco\n%s", err.Error())
	}
	return nil
}

func (n *NeonDB) Connection() *gorm.DB {
	godotenv.Load()
	dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	dbSQL, err := db.DB()
	if err != nil {
		log.Fatal(err.Error())
	}

	dbSQL.SetMaxIdleConns(10)
	dbSQL.SetMaxOpenConns(50)
	dbSQL.SetConnMaxLifetime(10 * time.Minute)

	return db
}
