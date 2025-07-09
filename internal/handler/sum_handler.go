package handler

import (
	"MyGoProject/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SumRequest struct {
	Numbers []float64 `json:"numbers"`
}

type SumResponse struct {
	Result float64 `json:"result"`
}

func SumHandler(c echo.Context) error {
	var req SumRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}

	result := service.CalculateSum(req.Numbers)
	return c.JSON(http.StatusOK, SumResponse{Result: result})
}
