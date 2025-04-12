package main

import (
	. "tg-bot/internal/bot"
)

const filePath = "configs/config.json" // or switch to $PATH

func main() {
	bot := NewTgBot(filePath)
	bot.Start()
}
