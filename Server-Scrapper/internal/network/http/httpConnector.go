package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	dto "server-scrapper/pkg/dto"
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

func AddUser(c echo.Context) error {
	var client dto.Client

	if err := c.Bind(&client); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}

	fmt.Printf("User received: %+v\n", client)

	return c.JSON(http.StatusOK, map[string]string{
		"message": "User added successfully",
	})
}
