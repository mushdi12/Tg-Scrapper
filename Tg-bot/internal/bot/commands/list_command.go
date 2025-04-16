package commands

import (
	"log"
	"tg-bot/internal/network"
)

type ListCommand struct{}

func (cmd *ListCommand) Execute(ctx CommandContext) string {
	userLinks, err := network.UsersLinkRequest(ctx.ChatId)
	if err != nil {
		log.Printf("[ListCommand] ошибка получения ссылок: %v", err)
		return "Произошла ошибка, попробуйте позже"
	}
	return userLinks
}
