package main

import (
	"github.com/Pavel-Casp/MyGoProject1/internal/handler"
	"github.com/Pavel-Casp/MyGoProject1/internal/service"
	"github.com/Pavel-Casp/MyGoProject1/internal/storage"
	"github.com/labstack/echo/v4"
)

func main() {
	// Инициализация с передачей хранилища
	storage := storage.NewMemoryStorage()
	sumService := service.NewSumService(storage) // Теперь передаем storage
	sumHandler := handler.NewSumHandler(sumService)

	e := echo.New()
	e.POST("/sum", sumHandler.Handle)
	e.Logger.Fatal(e.Start(":8080"))
}
