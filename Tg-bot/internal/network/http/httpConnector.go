package http

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	dto "tg-bot/pkg/dto"
)

//var client = http.Client{Timeout: 3 * time.Second}

func SendRequest(chatId int64, username string) {
	done := make(chan<- interface{})
	go func() {
		user := dto.Client{
			ChatId:   chatId,
			UserName: username,
		}

		userData, err := json.Marshal(user)
		if err != nil {
			log.Fatalf("Ошибка сериализации структуры: %v", err)
		}

		resp, err := http.Post("http://localhost:8080/users", "application/json", bytes.NewBuffer(userData))
		if err != nil {
			log.Fatalf("Ошибка отправки запроса: %v", err)
		}
		defer resp.Body.Close()

		done <- struct{}{}
	}()
}
