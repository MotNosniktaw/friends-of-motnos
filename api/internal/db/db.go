package db

import (
	"github.com/motnosniktaw/nflfantasy/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost port=5432 user=postgres password=password dbname=fom sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to database")
	}

	DB = db

	DB.AutoMigrate(&models.Player{})
}
