package user

import "link-tracker/internal/bot"

var Users = make(map[int64]User)

const (
	NONE = iota
	WaitingUrl
	WaitingFilter
	WaitingHashtag
	WaitingUrlForRemove
)

type UserInterface interface {
	getState() int
}

type User struct {
	State    int
	Link     string
	Filter   string
	Category string
}

func GetState(chatId int64) int {
	bot.mu.Lock()
	userState := Users[chatId].State
	bot.mu.Unlock()
	return userState
}
