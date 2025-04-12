package postgres

import (
	. "Server-Scrapper/internal/fileManager"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(filePath string) error {
	config := LoadConfig(filePath)
	var err error
	DB, err = gorm.Open(postgres.Open(config.PostgresData), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
