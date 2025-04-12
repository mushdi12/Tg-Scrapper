package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tg-bot/internal/user"
)

func Controller(updates tgbotapi.UpdatesChannel, bot *TgBot, stopChan <-chan struct{}) {
	for {
		select {
		case update, ok := <-updates:
			if !ok {
				return
			}

			if update.Message == nil {
				continue
			}

			user := update.Message.From
			chatID := update.Message.Chat.ID

			if update.Message.IsCommand() {
				command := update.Message.Command()
				handleCommand(user.UserName, command, chatID, bot)
			} else {
				message := update.Message.Text
				handleMessage(message, chatID, bot)
			}

		case <-stopChan:
			return
		}
	}
}

func handleCommand(username string, cmd string, chatID int64, bot *TgBot) {
	if command, exists := CommandRegistry[cmd]; exists {
		command.Execute(username, chatID, bot)
	} else {
		bot.SendMessage(chatID, "Неизвестная команда! Используй /help")
	}
}

func handleMessage(message string, chatID int64, bot *TgBot) {
	state := user.GetState(chatID)
	switch state {
	case user.WaitingUrl, user.WaitingFilter, user.WaitingHashtag:
		HandleAsync(bot, chatID, func(user user.User) (user.User, string) {
			return RealizationTrack(user, message)
		})
	case user.WaitingUrlForRemove:
		HandleAsync(bot, chatID, func(user user.User) (user.User, string) {
			return RealizationUnTrack(user, message)
		})
	default:
		bot.SendMessage(chatID, "Неизвестная команда! Используй /help")
	}
}
