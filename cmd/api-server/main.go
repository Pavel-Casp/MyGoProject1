package main

import (
	"MyGoProject/internal/config"
	"MyGoProject/internal/server"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	cfg := config.Load()
	server.Setup(e, cfg)

	e.Logger.Fatal(e.Start(":8080"))
}
