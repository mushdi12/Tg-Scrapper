package network

import (
	"fmt"
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

func SaveClient(client Client) error {
	if DB == nil {
		return fmt.Errorf("database not initialized")
	}
	return DB.Create(&client).Error
}

func SaveClientLink(usersLink Client) error {
	if DB == nil {
		return fmt.Errorf("database not initialized")
	}
	return DB.Create(&usersLink).Error
}
