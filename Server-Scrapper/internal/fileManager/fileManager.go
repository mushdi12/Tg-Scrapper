package fileManager

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type ServerConfig struct {
	PostgresData string `json:"postgres_data"`
}

func LoadConfig(filePath string) *ServerConfig {

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Ошибка чтения конфигурационного файла: %v", err)
	}

	var config ServerConfig
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil
	}
	return &config
}
