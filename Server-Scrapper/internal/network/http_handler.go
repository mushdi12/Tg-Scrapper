package network

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RemoveLinkHandler(c echo.Context) error {
	var err error
	// принимаем данные
	return err
}

func AddLinkHandler(c echo.Context) error {
	var userLink Client

	if err := c.Bind(&userLink); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}

	// Запускаем сохранение в отдельной горутине
	go func() {
		if err := SaveClientLink(userLink); err != nil {
			fmt.Printf("Error saving client: %v\n", err)
		}
	}()

	return c.JSON(http.StatusOK, map[string]string{
		"message": "User add initiated",
	})
}

func AddUserHandler(c echo.Context) error {
	var client Client

	if err := c.Bind(&client); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}

	// Запускаем сохранение в отдельной горутине
	go func() {
		if err := SaveClient(client); err != nil {
			fmt.Printf("Error saving client: %v\n", err)
		}
	}()

	return c.JSON(http.StatusOK, map[string]string{
		"message": "User add initiated",
	})
}

func GetLinksHandler(c echo.Context) error {
	var userLink Client

	if err := c.Bind(&userLink); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}

	// Запускаем сохранение в отдельной горутине
	go func() {
		if err := SaveClientLink(userLink); err != nil {
			fmt.Printf("Error saving client: %v\n", err)
		}
	}()

	return c.JSON(http.StatusOK, map[string]string{
		"message": "User add initiated",
	})
}
