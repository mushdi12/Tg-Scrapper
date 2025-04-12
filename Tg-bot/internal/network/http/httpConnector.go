package http

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

var client = http.Client{Timeout: 3 * time.Second}

type user struct {
	ChatID   int64  `json:"chatID"`
	Username string `json:"username"`
}

func SendRequest() {
	done := make(chan<- interface{})
	go func() {
		user := user{
			ChatID:   123456789,
			Username: "new_user",
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
