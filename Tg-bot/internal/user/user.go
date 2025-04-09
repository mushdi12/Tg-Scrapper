package user

import (
	"sync"
)

var (
	Users = make(map[int64]User)
	mu    sync.Mutex
)

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
	mu.Lock()
	userState := Users[chatId].State
	mu.Unlock()
	return userState
}
