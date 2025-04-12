package main

import (
	. "Server-Scrapper/internal/db/postgres"
	. "Server-Scrapper/internal/server"
	"errors"
	"log"
	"log/slog"
	"net/http"
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
