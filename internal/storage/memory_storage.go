package storage

import (
	"sync"
)

type MemoryStorage struct {
	mu   sync.Mutex
	data map[string]float64
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[string]float64),
	}
}

func (s *MemoryStorage) AddResult(key string, value float64) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
	return nil
}

func (s *MemoryStorage) GetResult(key string) (float64, bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	val, ok := s.data[key]
	return val, ok, nil
}
