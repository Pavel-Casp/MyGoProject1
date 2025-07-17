package storage

import (
	"sync"
)

type MemoryStorage struct {
	mu      sync.RWMutex
	storage map[string]map[string]float64
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		storage: make(map[string]map[string]float64),
	}
}

func (s *MemoryStorage) AddResult(userToken, key string, value float64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.storage[userToken]; !exists {
		s.storage[userToken] = make(map[string]float64)
	}

	s.storage[userToken][key] = value
	return nil
}

func (s *MemoryStorage) GetResult(userToken, key string) (float64, bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	userResults, exists := s.storage[userToken]
	if !exists {
		return 0, false, nil
	}

	val, ok := userResults[key]
	return val, ok, nil
}

func (s *MemoryStorage) ClearUserData(userToken string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.storage, userToken)
	return nil
}
