package main

import (
	"github.com/Pavel-Casp/MyGoProject1/internal/handler"
	"github.com/Pavel-Casp/MyGoProject1/internal/service"
	"github.com/Pavel-Casp/MyGoProject1/internal/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Инициализация
	store := storage.NewMemoryStorage()
	calcService := service.NewCalculatorService(store)
	calcHandler := handler.NewCalculatorHandler(calcService)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Роуты
	e.GET("/token", calcHandler.GenerateToken)
	e.POST("/sum", calcHandler.HandleSum)
	e.POST("/multiply", calcHandler.HandleMultiply)
	e.GET("/results/:key", calcHandler.GetResult)
	e.DELETE("/clear", calcHandler.ClearData)

	e.Logger.Fatal(e.Start(":8080"))
}
