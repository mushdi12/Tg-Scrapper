package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	. "server-scrapper/internal/network/http"
)

type Server struct {
	*echo.Echo
}

func NewServer() *Server {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover()) // для восстановления из паники

	e.POST("users", AddUser)
	e.POST("user/addlink", AddLink)
	e.DELETE("user/removelink", RemoveLink)
	return &Server{e}
}
