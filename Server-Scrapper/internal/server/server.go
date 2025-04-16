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

	e.POST("user/addUser", AddUser)
	e.POST("user/addLink", AddLink)
	e.GET("user/getLinks", GetLinks)
	e.DELETE("user/removeLink", RemoveLink)
	return &Server{e}
}
