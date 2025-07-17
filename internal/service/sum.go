package service

import (
	"fmt"
	"strings"

	"github.com/Pavel-Casp/MyGoProject1/internal/storage"
	"github.com/google/uuid"
)

type CalculatorService struct {
	storage *storage.MemoryStorage
}

func NewCalculatorService(storage *storage.MemoryStorage) *CalculatorService {
	return &CalculatorService{storage: storage}
}

func (s *CalculatorService) CalculateSum(userToken string, numbers []float64) (float64, string, error) {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}

	key := generateKey("sum", numbers)
	if err := s.storage.AddResult(userToken, key, sum); err != nil {
		return 0, "", fmt.Errorf("failed to store result: %w", err)
	}

	return sum, key, nil
}

func (s *CalculatorService) CalculateMultiply(userToken string, numbers []float64) (float64, string, error) {
	product := 1.0
	for _, num := range numbers {
		product *= num
	}

	key := generateKey("multiply", numbers)
	if err := s.storage.AddResult(userToken, key, product); err != nil {
		return 0, "", fmt.Errorf("failed to store result: %w", err)
	}

	return product, key, nil
}

func (s *CalculatorService) GetStoredResult(userToken, key string) (float64, bool, error) {
	return s.storage.GetResult(userToken, key)
}

func (s *CalculatorService) ClearUserData(userToken string) error {
	return s.storage.ClearUserData(userToken)
}

func generateKey(operation string, numbers []float64) string {
	var parts []string
	for _, num := range numbers {
		parts = append(parts, fmt.Sprintf("%.2f", num))
	}
	return operation + ":" + strings.Join(parts, "-")
}

func GenerateUserToken() string {
	return uuid.NewString()
}
