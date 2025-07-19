package handler

import (
	"net/http"

	"github.com/Pavel-Casp/MyGoProject1/internal/service"
	"github.com/labstack/echo/v4"
)

type (
	// CalculatorRequest represents the request structure for calculator operations
	// @Description Request payload for calculator operations
	CalculatorRequest struct {
		Numbers []float64 `json:"numbers" validate:"required,min=1"`
		Token   string    `json:"token" validate:"required,uuid4"`
	}

	// CalculatorResponse represents the response structure for calculator operations
	// @Description Response payload for calculator operations
	CalculatorResponse struct {
		Result float64 `json:"result"`
		Key    string  `json:"key,omitempty"`
		Token  string  `json:"token,omitempty"`
	}

	// CalculatorHandler handles calculator HTTP requests
	CalculatorHandler struct {
		service *service.CalculatorService
	}
)

// NewCalculatorHandler creates a new CalculatorHandler instance
func NewCalculatorHandler(s *service.CalculatorService) *CalculatorHandler {
	return &CalculatorHandler{service: s}
}

// GenerateToken godoc
// @Summary Generate a new user token
// @Description Generates a new UUID token for user identification
// @Tags token
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "Generated token"
// @Router /token [get]
func (h *CalculatorHandler) GenerateToken(c echo.Context) error {
	token := service.GenerateUserToken()
	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

// HandleSum godoc
// @Summary Calculate sum of numbers
// @Description Accepts an array of numbers and returns their sum
// @Tags calculator
// @Accept json
// @Produce json
// @Param request body CalculatorRequest true "Numbers to sum and user token"
// @Success 200 {object} CalculatorResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /sum [post]
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

// HandleMultiply godoc
// @Summary Calculate product of numbers
// @Description Accepts an array of numbers and returns their product
// @Tags calculator
// @Accept json
// @Produce json
// @Param request body CalculatorRequest true "Numbers to multiply and user token"
// @Success 200 {object} CalculatorResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /multiply [post]
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

// GetResult godoc
// @Summary Get stored calculation result
// @Description Retrieves a previously stored calculation result by key
// @Tags results
// @Accept json
// @Produce json
// @Param key path string true "Result key"
// @Param token query string true "User token"
// @Success 200 {object} CalculatorResponse
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /results/{key} [get]
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

// ClearData godoc
// @Summary Clear user data
// @Description Clears all stored calculation results for a user
// @Tags results
// @Accept json
// @Produce json
// @Param token query string true "User token"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /clear [delete]
func (h *CalculatorHandler) ClearData(c echo.Context) error {
	token := c.QueryParam("token")
	if err := h.service.ClearUserData(token); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"status": "data cleared"})
}
