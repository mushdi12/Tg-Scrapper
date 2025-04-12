package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	. "server-scrapper/internal/fileManager"
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
