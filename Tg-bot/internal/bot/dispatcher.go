package bot

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	commands2 "tg-bot/internal/commands"
	user "tg-bot/internal/user"
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
	if command, exists := commands2.CommandRegistry[cmd]; exists {

		commandCtx := commands2.CommandContext{ChatId: chatID, Username: username}
		messageForUser := make(chan string)

		if cmd == "help" {
			botCmds, err := bot.GetMyCommands()
			if err != nil {
				log.Println(err.Error())
				commandCtx.BotCmd = nil
			} else {
				commandCtx.BotCmd = botCmds
			}
		}

		go func() {
			msg := command.Execute(commandCtx)
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
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	state := user.GetState(chatID)
	_, addFlag := user.AddStates[state]
	_, removeFlag := user.AddStates[state]

	if addFlag || removeFlag {
		messageForUser := make(chan string)
		var command commands2.Command
		if addFlag {
			command = &commands2.TrackCommand{}
		} else {
			command = &commands2.UnTrackCommand{}
		}
		go func() {
			msg := command.Execute(commands2.CommandContext{ChatId: chatID, Message: message})
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
