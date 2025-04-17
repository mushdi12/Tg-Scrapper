package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	. "server-scrapper/internal/network"
)

type Server struct {
	*echo.Echo
}

func NewServer() *Server {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover()) // для восстановления из паники

	e.POST("user/addUser", AddUserHandler)
	e.POST("user/addLink", AddLinkHandler)
	e.GET("user/getLinks", GetLinksHandler)
	e.DELETE("user/removeLink", RemoveLinkHandler)
	return &Server{e}
}
