package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"os/signal"
	"sync"
	"tg-bot/internal/fileManager"
)

type Bot interface {
	Start()
	SendMessage(chatId int64, message string)
	Stop()
}
type TgBot struct {
	*tgbotapi.BotAPI
	stopChan chan struct{}
	wg       sync.WaitGroup
}

// TgBot constructor
func NewTgBot(filePath string) *TgBot {
	config, err := fileManager.LoadConfig(filePath)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	bot, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	botCommands, err := fileManager.CommandConverter(config.Commands)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	setCommands := tgbotapi.NewSetMyCommands(botCommands...)
	_, err = bot.Request(setCommands)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	log.Println("Created new bot")

	return &TgBot{BotAPI: bot}
}

func (bot *TgBot) Start() {

	log.Println("Starting Telegram Bot")
	bot.stopChan = make(chan struct{})

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	bot.wg.Add(1)
	go func() {
		defer bot.wg.Done()
		Controller(updates, bot, bot.stopChan)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	bot.Stop()
}

func (bot *TgBot) SendMessage(chatId int64, message string) {
	msg := tgbotapi.NewMessage(chatId, message)
	msg.DisableWebPagePreview = true
	_, err := bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}

func (bot *TgBot) Stop() {
	log.Println("Выключение бота...")
	bot.StopReceivingUpdates()
	close(bot.stopChan)
	bot.wg.Wait()
	log.Println("Бот успешно остановлен.")
}
