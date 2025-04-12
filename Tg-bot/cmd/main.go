package main

import (
	"github.com/mushdi12/Tg-LinkTracker/tree/main/Server-Scrapper/pkg/dto "
	. "link-tracker/internal/bot"
)

const filePath = "configs/config.json" // or switch to $PATH

func main() {
	P
	bot := NewTgBot(filePath)
	bot.Start()
}
