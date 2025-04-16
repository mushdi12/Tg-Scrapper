package network

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"tg-bot/internal/user"
	dto "tg-bot/pkg/dto"
	"time"
)

const (
	POST = "POST"
	GET  = "GET"
)

var (
	httpClient = http.Client{Timeout: 30 * time.Second}
	ServerURL  = "ServerURL"
)

func AddUserRequest(chatId int64, username string) error {
	client := dto.Client{ChatId: chatId, UserName: username}
	return maskeRequest(ServerURL+"addUser", POST, client)
}

func AddLinkRequest(chatId int64, link, category, filter string) error {
	linkForAdd := dto.AddLink{ChatId: chatId, Link: link, Category: category, Filters: filter}
	return maskeRequest(ServerURL+"addLink", POST, linkForAdd)
}

func RemoveLinkRequest(chatId int64, link string) error {
	linkForRemove := dto.RemoveLink{ChatId: chatId, Link: link}
	return maskeRequest(ServerURL+"removeLink", POST, linkForRemove)
}

// сделать нормальный запрос
func UsersLinkRequest(chatId int64) (string, error) {
	result := getUserLinks(chatId)
	return result, nil
}

// replace with server connection
func getUserLinks(chatId int64) string {
	if _, exists := user.Users[chatId]; exists {
		localUser := user.Users[chatId]
		return fmt.Sprintf("Ваши ссылки :\n" + "Категория: #" + localUser.Category + "\n" + "Фильтр: " + localUser.Filter + "\n" + localUser.Link)
	}
	return "Ошибка, попробуйте еще раз!"
}

func maskeRequest(url string, htttpMethod string, payload any) error {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Ошибка сериализации структуры: %v", err)
		return err
	}

	req, err := http.NewRequest(htttpMethod, url, bytes.NewBuffer(data))
	if err != nil {
		log.Printf("Ошибка создания запроса: %v", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Printf("Ошибка отправки запроса: %v", err)
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Ошибка чтения ответа: %v", err)
		return err
	}
	if resp.StatusCode != http.StatusOK {
		log.Printf("Ответ сервера: %s", string(body))
		return errors.New(string(body))
	}

	return nil
}
