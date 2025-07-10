package main

import (
	"github.com/Pavel-Casp/MyGoProject1/internal/config"
	"github.com/Pavel-Casp/MyGoProject1/internal/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	// Загрузка конфига
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	// Настройка логгера Echo
	e.Logger.SetLevel(cfg.Logger.Level) // "debug", "info", "warn", "error"

	// Роутинг
	e.POST("/sum", handler.SumHandler)

	// Запуск сервера
	e.Logger.Fatal(e.Start(":" + cfg.Server.Port))
}
