package bot

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tg-bot/internal/bot/commands"
	user "tg-bot/internal/user"
	"tg-bot/pkg/async"
	"time"
)

func MainController(updates tgbotapi.UpdatesChannel, bot *TgBot, stopChan <-chan struct{}) {
	for {
		select {
		case update := <-updates:

			chatID := update.Message.Chat.ID
			username := update.Message.From.UserName

			if update.Message.IsCommand() {
				command := update.Message.Command()
				go handleCommand(username, command, chatID, bot)
			} else {
				message := update.Message.Text
				go handleMessage(message, chatID, bot)
			}

		case <-stopChan:
			return
		}
	}
}

func handleCommand(username string, cmd string, chatID int64, bot *TgBot) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	if command, exists := commands.CommandRegistry[cmd]; exists {

		messageForUser := make(chan string)

		go func() {
			//msg := command.Execute(chatID, username)
			messageForUser <- msg
		}()

		select {
		case msg := <-messageForUser:
			bot.SendMessage(chatID, msg)
		case <-ctx.Done():
			bot.SendMessage(chatID, "Команда заняла слишком много времени, попробуйте еще раз!")
		}

	} else {
		bot.SendMessage(chatID, "Неизвестная команда! Воспользуйтесь /help")
	}
}

func handleMessage(message string, chatID int64, bot *TgBot) {
	state := user.GetState(chatID)
	if _, exists := user.AddStates[state]; exists {
		async.RunAsync(func() { RealizationTrack(bot, chatID, state, message) })
	} else if _, exists := user.RemoveStates[state]; exists {
		async.RunAsync(func() { RealizationUnTrack(bot, chatID, state, message) })
	} else {
		bot.SendMessage(chatID, "Неизвестная команда! Воспользуйтесь /help")
	}
}
