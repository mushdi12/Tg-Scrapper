package network

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	client := Client{ChatId: chatId, UserName: username}
	return maskeRequest(ServerURL+"addUser", POST, client)
}

func AddLinkRequest(chatId int64, link, category, filter string) error {
	linkForAdd := AddLink{ChatId: chatId, Link: link, Category: category, Filters: filter}
	return maskeRequest(ServerURL+"addLink", POST, linkForAdd)
}

func RemoveLinkRequest(chatId int64, link string) error {
	linkForRemove := RemoveLink{ChatId: chatId, Link: link}
	return maskeRequest(ServerURL+"removeLink", POST, linkForRemove)
}

func UsersLinkRequest(chatId int64) (string, error) {
	url := fmt.Sprintf("%sgetLinks?chatId=%d", ServerURL, chatId)
	return makeGetRequest(url)
}

func formatUserLinks(category, filter, link string) string {
	return fmt.Sprintf("Ваши ссылки:\nКатегория: #%s\nФильтр: %s\n%s", category, filter, link)
}

func makeGetRequest(url string) (string, error) {
	req, err := http.NewRequest(GET, url, nil)
	if err != nil {
		log.Printf("[UsersLinkRequest] ошибка создания GET-запроса: %v", err)
		return "", err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Printf("[UsersLinkRequest] ошибка отправки GET-запроса: %v", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[UsersLinkRequest] ошибка чтения ответа: %v", err)
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("[UsersLinkRequest] сервер вернул статус %d: %s", resp.StatusCode, string(body))
		return "", err
	}

	var response struct {
		Link     string `json:"link"`
		Category string `json:"category"`
		Filters  string `json:"filters"`
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Printf("[UsersLinkRequest] ошибка разбора JSON-ответа: %v", err)
		return "", err
	}

	return formatUserLinks(response.Category, response.Filters, response.Link), nil
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
