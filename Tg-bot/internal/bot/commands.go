package bot

import (
	"link-tracker/internal/user"
)

var CommandRegistry = map[string]Command{
	"start":   &StartCommand{},
	"help":    &HelpCommand{},
	"list":    &ListCommand{},
	"track":   &TrackCommand{},
	"untrack": &UnTrackCommand{},
}

type Command interface {
	Execute(chatID int64, bot *TgBot)
}

type StartCommand struct{}

type HelpCommand struct{}

type ListCommand struct{}

type TrackCommand struct{}

type UnTrackCommand struct{}

func (cmd *StartCommand) Execute(chatID int64, bot *TgBot) {
	runAsync(func() { CheckLogin(bot, chatID) })
}

func (cmd *HelpCommand) Execute(chatID int64, bot *TgBot) {
	runAsync(func() { SendHelp(bot, chatID) })
}

func (cmd *ListCommand) Execute(chatID int64, bot *TgBot) {
	runAsync(func() { SendList(bot, chatID) })
}

func (cmd *TrackCommand) Execute(chatID int64, bot *TgBot) {
	HandleAsync(bot, chatID, func(user user.User) (user.User, string) {
		return RealizationTrack(user, "") // добавить канал done
	})
}

func (cmd *UnTrackCommand) Execute(chatID int64, bot *TgBot) {
	HandleAsync(bot, chatID, func(user user.User) (user.User, string) {
		return RealizationUnTrack(user, "") // добавить канал done
	})
}
