package commands

import (
	"fmt"
	"strings"
)

type HelpCommand struct{}

func (cmd *HelpCommand) Execute(ctx CommandContext) string {
	commands := ctx.BotCmd
	if commands != nil {
		return "Произошла ошибка, попробуйте еще раз!"
	}

	var result strings.Builder
	for _, command := range commands {
		fmt.Fprintf(&result, "/%s - %s\n", command.Command, command.Description)
	}
	return result.String()
}
