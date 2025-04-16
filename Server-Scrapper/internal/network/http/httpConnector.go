package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	dbclient "server-scrapper/internal/database/postgres"
	"server-scrapper/pkg/dto"
)

func RemoveLink(c echo.Context) error {
	var err error
	// принимаем данные
	return err
}

func AddLink(c echo.Context) error {
	var userLink dto.UsersLinks

	if err := c.Bind(&userLink); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}

	// Запускаем сохранение в отдельной горутине
	go func() {
		if err := dbclient.SaveClientLink(userLink); err != nil {
			fmt.Printf("Error saving client: %v\n", err)
		}
	}()

	return c.JSON(http.StatusOK, map[string]string{
		"message": "User add initiated",
	})
}

func AddUser(c echo.Context) error {
	var client dto.User

	if err := c.Bind(&client); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}

	// Запускаем сохранение в отдельной горутине
	go func() {
		if err := dbclient.SaveClient(client); err != nil {
			fmt.Printf("Error saving client: %v\n", err)
		}
	}()

	return c.JSON(http.StatusOK, map[string]string{
		"message": "User add initiated",
	})
}

func GetLinks(c echo.Context) error {
	var userLink dto.UsersLinks

	if err := c.Bind(&userLink); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}

	// Запускаем сохранение в отдельной горутине
	go func() {
		if err := dbclient.SaveClientLink(userLink); err != nil {
			fmt.Printf("Error saving client: %v\n", err)
		}
	}()

	return c.JSON(http.StatusOK, map[string]string{
		"message": "User add initiated",
	})
}
