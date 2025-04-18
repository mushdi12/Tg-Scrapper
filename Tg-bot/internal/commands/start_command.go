package commands

import (
	"log"
	"sync"
	"tg-bot/internal/network"
	. "tg-bot/internal/user"
)

var mu sync.Mutex

type StartCommand struct{}

func (cmd *StartCommand) Execute(ctx CommandContext) string {
	mu.Lock()
	defer mu.Unlock()

	if _, ok := Users[ctx.ChatId]; !ok {
		Users[ctx.ChatId] = &User{ChatId: ctx.ChatId, State: NONE}
		err := network.AddUserRequest(ctx.ChatId, ctx.Username)
		if err != nil {
			log.Printf("[StartCommand] ошибка при регистрации: %v", err)
			return "Произошла ошибка, попробуйте еще раз!"
		}
		return "Вы успешно зарегистрировались!"
	}
	return "Пользователь уже зарегистрирован!"
}
