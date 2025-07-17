package handler

import (
	"net/http"

	"github.com/Pavel-Casp/MyGoProject1/internal/service"
	"github.com/labstack/echo/v4"
)

type (
	CalculatorRequest struct {
		Numbers []float64 `json:"numbers" validate:"required,min=1"`
		Token   string    `json:"token" validate:"required,uuid4"`
	}

	CalculatorResponse struct {
		Result float64 `json:"result"`
		Key    string  `json:"key,omitempty"`
		Token  string  `json:"token,omitempty"`
	}

	CalculatorHandler struct {
		service *service.CalculatorService
	}
)

func NewCalculatorHandler(s *service.CalculatorService) *CalculatorHandler {
	return &CalculatorHandler{service: s}
}

func (h *CalculatorHandler) GenerateToken(c echo.Context) error {
	token := service.GenerateUserToken()
	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

func (h *CalculatorHandler) HandleSum(c echo.Context) error {
	var req CalculatorRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request format"})
	}

	result, key, err := h.service.CalculateSum(req.Token, req.Numbers)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, CalculatorResponse{
		Result: result,
		Key:    key,
		Token:  req.Token,
	})
}

func (h *CalculatorHandler) HandleMultiply(c echo.Context) error {
	var req CalculatorRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request format"})
	}

	result, key, err := h.service.CalculateMultiply(req.Token, req.Numbers)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, CalculatorResponse{
		Result: result,
		Key:    key,
		Token:  req.Token,
	})
}

func (h *CalculatorHandler) GetResult(c echo.Context) error {
	token := c.QueryParam("token")
	key := c.Param("key")

	result, exists, err := h.service.GetStoredResult(token, key)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if !exists {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Result not found"})
	}

	return c.JSON(http.StatusOK, CalculatorResponse{
		Result: result,
		Key:    key,
		Token:  token,
	})
}

func (h *CalculatorHandler) ClearData(c echo.Context) error {
	token := c.QueryParam("token")
	if err := h.service.ClearUserData(token); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"status": "data cleared"})
}
