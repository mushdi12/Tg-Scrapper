package user

import (
	"sync"
)

var (
	Users = make(map[int64]*User)
	mu    sync.Mutex
)

type UserInterface interface {
	getState() int
}

type User struct {
	ChatId   int64
	State    int
	Link     string
	Filter   string
	Category string
}

func GetState(chatId int64) int {
	mu.Lock()
	userState := Users[chatId].State
	mu.Unlock()
	return userState
}
