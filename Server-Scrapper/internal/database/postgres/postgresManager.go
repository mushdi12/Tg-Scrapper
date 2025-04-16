package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	. "server-scrapper/internal/fileManager"
	"server-scrapper/pkg/dto"
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

func SaveClient(client dto.User) error {
	if DB == nil {
		return fmt.Errorf("database not initialized")
	}
	return DB.Create(&client).Error
}

func SaveClientLink(usersLink dto.UsersLinks) error {
	if DB == nil {
		return fmt.Errorf("database not initialized")
	}
	return DB.Create(&usersLink).Error
}
