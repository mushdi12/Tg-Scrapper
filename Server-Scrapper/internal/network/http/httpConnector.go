package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RemoveLink(c echo.Context) error {
	var err error
	// принимаем данные
	return err
}

func AddLink(c echo.Context) error {
	var err error

	return err
}

type user struct {
	ChatID   int64  `json:"chatID"`
	Username string `json:"username"`
	RegDate  string `json:"reg_date"`
}

func AddUser(c echo.Context) error {
	var user user

	// Применяем Bind, чтобы распарсить тело запроса в структуру user
	if err := c.Bind(&user); err != nil {
		// Если ошибка при парсинге, возвращаем ошибку
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}

	// Выводим полученные данные на сервер
	fmt.Printf("User received: %+v\n", user)

	// Возвращаем успешный ответ
	return c.JSON(http.StatusOK, map[string]string{
		"message": "User added successfully",
	})
}
