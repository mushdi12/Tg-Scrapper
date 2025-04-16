package fileManager

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io/ioutil"
)

type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

type Config struct {
	Token     string       `json:"token"`
	Commands  []BotCommand `json:"commands"`
	ServerURL string       `json:"server_url"`
}

func LoadConfig(filePath string) (*Config, error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func CommandConverter(commands []BotCommand) ([]tgbotapi.BotCommand, error) {
	var botCommands []tgbotapi.BotCommand
	for _, cmd := range commands {
		botCommands = append(botCommands, tgbotapi.BotCommand{Command: cmd.Command, Description: cmd.Description})
	}
	return botCommands, nil
}
