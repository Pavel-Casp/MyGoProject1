package handler

import (
	"net/http"

	"github.com/Pavel-Casp/MyGoProject1/internal/service" // Добавьте этот импорт
	"github.com/labstack/echo/v4"
)

type Request struct {
	Numbers []float64 `json:"numbers"`
}

type Response struct {
	Result float64 `json:"result"`
}

func SumHandler(c echo.Context) error {
	var req Request
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}

	sum := service.Sum(req.Numbers)
	return c.JSON(http.StatusOK, Response{Result: sum})
}
