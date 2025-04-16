package commands

import (
	"fmt"
	"log"
	"strings"
	"tg-bot/internal/bot"
)

type HelpCommand struct{}

func (cmd *HelpCommand) Execute(_ CommandContext) string {
	commands, err := bot.GetBotCommand()
	if err != nil {
		log.Printf("[HelpCommand] ошибка получения списка команд: %v", err)
		return "Произошла ошибка, попробуйте еще раз!"
	}

	var result strings.Builder
	for _, command := range commands {
		fmt.Fprintf(&result, "/%s - %s\n", command.Command, command.Description)
	}
	return result.String()
}
