package main

import (
	"errors"
	"log"
	"log/slog"
	"net/http"
	. "server-scrapper/internal/network"
	. "server-scrapper/internal/server"
)

func main() {
	err := Connect("configs/bd_config.json")
	if err != nil {
		log.Fatal(err)
	}
	e := NewServer()
	if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
	}

}
