package bot

import (
	"fmt"
	"strings"
	"sync"
	. "tg-bot/internal/network/http"
	. "tg-bot/internal/user"
)

var (
	wg sync.WaitGroup
	mu sync.Mutex
)

func SendList(bot *TgBot, chatId int64) { // переписать
	result := getUserLinks(chatId)
	bot.SendMessage(chatId, result)
}

func SendHelp(bot *TgBot, chatId int64) { // переписать
	commands, err := bot.GetMyCommands()
	if err != nil {
		bot.SendMessage(chatId, "Произошла ошибка, попробуйте еще раз")
	}

	var sb strings.Builder
	for _, command := range commands {
		_, _ = fmt.Fprintf(&sb, "/%s - %s\n", command.Command, command.Description)
	}

	messageText := sb.String()

	bot.SendMessage(chatId, messageText)
}

func runAsync(fn func()) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		fn()
	}()
}

func HandleAsync(bot *TgBot, chatID int64, fn func(User) (User, string)) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()
		u := Users[chatID]
		mu.Unlock()

		updateUser, answer := fn(u)

		mu.Lock()
		Users[chatID] = updateUser
		mu.Unlock()

		bot.SendMessage(chatID, answer)

	}()
}

func RealizationTrack(user User, message string) (User, string) { // переписать
	switch user.State {
	case NONE:
		user.State = WaitingUrl
		return user, "Пришлите мне ссылку"
	case WaitingUrl:
		user.State = WaitingFilter
		user.Link = message
		return user, "Напишите фильтр для ссылки"
	case WaitingFilter:
		user.State = WaitingHashtag
		user.Filter = message
		return user, "Назовите категорию ссылки"
	case WaitingHashtag:
		user.State = NONE
		user.Category = message
		// отправка на сервер
		return user, "Ваша ссылка: " + user.Link + "\n" + "Ваш фильтр: " + user.Filter + "\n" + "Категория ссылки: " + user.Category
	default:
		user.State = NONE
		return user, "Ошибка, действие вашей предыдущей команды отменено!❌!"
	}
}

func RealizationUnTrack(user User, message string) (User, string) { // переписать
	switch user.State {
	case NONE:
		user.State = WaitingUrlForRemove
		return user, "Пришлите мне ссылку"
	case WaitingUrlForRemove:
		user.State = NONE
		user.Link = message
		return user, "Ваша ссылка: " + user.Link + " удалена!"
	default:
		user.State = NONE
		return user, "Ошибка, действие вашей предыдущей команды отменено!❌"
	}
}

func CheckLogin(username string, bot *TgBot, chatId int64) {
	mu.Lock()
	if _, ok := Users[chatId]; !ok {
		Users[chatId] = User{State: NONE}
		SendRequest(chatId, username) // <-----
		bot.SendMessage(chatId, "Вы успешно зарегистрировались!")
	} else {
		bot.SendMessage(chatId, "Пользователь уже зарегистрирован!")
	}
	mu.Unlock()
}

func getUserLinks(chatId int64) string {
	if _, exists := Users[chatId]; exists { // replace with server connection
		user := Users[chatId]
		return fmt.Sprintf("Ваши ссылки :\n" + "Категория: #" + user.Category + "\n" + user.Link)
	}
	return "Ошибка, попробуйте еще раз!"
}
