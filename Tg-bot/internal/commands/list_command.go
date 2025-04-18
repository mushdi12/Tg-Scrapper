package commands

import (
	"log"
	"tg-bot/internal/network"
	"tg-bot/internal/user"
)

type ListCommand struct{}

func (cmd *ListCommand) Execute(ctx CommandContext) string {
	if !checkUser(ctx.ChatId) {
		return "Сначала зарегистрируйтесь -> /start"
	}
	userLinks, err := network.UsersLinkRequest(ctx.ChatId)
	if err != nil {
		log.Printf("[ListCommand] ошибка получения ссылок: %v", err)
		return "Произошла ошибка, попробуйте позже"
	}
	return userLinks
}

func checkUser(chatId int64) bool {
	mu.Lock()
	defer mu.Unlock()
	_, exist := user.Users[chatId]
	return exist
}
