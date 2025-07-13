package handler

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/Pavel-Casp/MyGoProject1/internal/service"
	"github.com/labstack/echo/v4"
)

type (
	Request struct {
		Numbers []float64 `json:"numbers"`
	}

	Response struct {
		Result float64 `json:"result"`
		ID     string  `json:"id,omitempty"`
	}

	SumHandler struct {
		service *service.SumService
		cache   map[string]float64
		mu      sync.Mutex
	}
)

// Конструктор
func NewSumHandler(s *service.SumService) *SumHandler {
	return &SumHandler{
		service: s,
		cache:   make(map[string]float64),
	}
}

// Метод для обработки запроса
func (h *SumHandler) Handle(c echo.Context) error {
	var req Request
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}

	key := generateKey(req.Numbers)

	// Проверка кэша
	h.mu.Lock()
	if result, exists := h.cache[key]; exists {
		h.mu.Unlock()
		return c.JSON(http.StatusOK, Response{Result: result, ID: key})
	}
	h.mu.Unlock()

	// Вычисление и сохранение
	sum, err := h.service.CalculateAndStore(req.Numbers)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Calculation failed"})
	}

	// Обновление кэша
	h.mu.Lock()
	h.cache[key] = sum
	h.mu.Unlock()

	return c.JSON(http.StatusOK, Response{Result: sum, ID: key})
}

func generateKey(numbers []float64) string {
	var key string
	for _, num := range numbers {
		key += fmt.Sprintf("%.2f-", num)
	}
	return key
}
