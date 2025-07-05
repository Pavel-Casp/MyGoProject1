package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Request struct {
	Numbers []float64 `json:"numbers"`
}

type Response struct {
	Result float64 `json:"result"`
}

func sumHandler(c echo.Context) error {
	var req Request
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}

	var sum float64
	for _, n := range req.Numbers {
		sum += n
	}

	return c.JSON(http.StatusOK, Response{Result: sum})
}

func main() {
	e := echo.New()

	// Роутинг
	e.POST("/sum", sumHandler)

	// Запуск сервера
	e.Logger.Fatal(e.Start(":8080"))
}
