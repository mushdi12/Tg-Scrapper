package commands

import (
	"log"
	"tg-bot/internal/network"
	. "tg-bot/internal/user"
)

type TrackCommand struct{}

func (cmd *TrackCommand) Execute(ctx CommandContext) string {
	mu.Lock()
	defer mu.Unlock()

	user, ok := Users[ctx.ChatId]
	if !ok {
		log.Printf("[TrackCommand.Execute] пользователь не найден в Users")
		return "Сначала зарегистрируйтесь -> /start"
	}

	state := user.State

	if ctx.Message == "" && state != NONE {
		ResetState(ctx.ChatId)
		return "Ошибка! Действие Команды отменено"
	}

	stmf := AddStates[state]

	if stmf.FieldtoChange != "" {
		setUserField(user, stmf.FieldtoChange, ctx.Message)
	}

	user.State = stmf.NextState

	if state == WaitingHashtag {
		err := network.AddLinkRequest(ctx.ChatId, user.Link, user.Category, user.Filter)
		if err != nil {
			log.Printf("[TrackCommand.Execute] ошибка при добавлении ссылки: %v", err)
			return "Не удалось добавить ссылку, попробуйте снова"
		}
	}

	return stmf.Message
}
