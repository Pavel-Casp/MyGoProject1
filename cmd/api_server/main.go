package main

import (
	"github.com/Pavel-Casp/MyGoProject1/internal/config"
	"github.com/Pavel-Casp/MyGoProject1/internal/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.POST("/sum", handler.SumHandler)
	e.Logger.Fatal(e.Start(":" + cfg.Server.Port))
}
