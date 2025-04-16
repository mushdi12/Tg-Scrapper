package commands

import (
	"reflect"
	. "tg-bot/internal/user"
)

var CommandRegistry = map[string]Command{
	"start":   &StartCommand{},
	"help":    &HelpCommand{},
	"list":    &ListCommand{},
	"track":   &TrackCommand{},
	"untrack": &UnTrackCommand{},
}

type CommandContext struct {
	ChatId   int64
	Username string
	Message  string // для Track/Untrack
}

type Command interface {
	Execute(ctx CommandContext) string
}

func setUserField(user *User, fieldName string, value string) {
	reflect.ValueOf(user).Elem().FieldByName(fieldName).SetString(value)
}
