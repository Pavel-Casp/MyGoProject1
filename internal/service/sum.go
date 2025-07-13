package service

import (
	"fmt"
	"strings"

	"github.com/Pavel-Casp/MyGoProject1/internal/storage"
)

type SumService struct {
	storage *storage.MemoryStorage
}

func NewSumService(storage *storage.MemoryStorage) *SumService {
	return &SumService{storage: storage}
}

func (s *SumService) CalculateAndStore(numbers []float64) (float64, error) {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}

	key := generateKey(numbers)
	if err := s.storage.AddResult(key, sum); err != nil {
		return 0, fmt.Errorf("failed to store result: %w", err)
	}

	return sum, nil
}

func (s *SumService) GetStoredResult(key string) (float64, bool, error) {
	value, exists, err := s.storage.GetResult(key)
	if err != nil {
		return 0, false, fmt.Errorf("storage error: %w", err)
	}
	return value, exists, nil
}

func generateKey(numbers []float64) string {
	var parts []string
	for _, num := range numbers {
		parts = append(parts, fmt.Sprintf("%.2f", num))
	}
	return strings.Join(parts, "-")
}
