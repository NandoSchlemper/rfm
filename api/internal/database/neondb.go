package database

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"rfmtransportes/internal/database"
	"rfmtransportes/internal/models"
)

type NeonDB struct {
	DB *gorm.DB
}

func (n *NeonDB) Connection() error {
	godotenv.Load()
}
