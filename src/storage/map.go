package storage

import (
	"fmt"
	"sync"
)

type (
	mapDB struct {
		m sync.Map
	}
)

func NewMapDB() API {
	return &mapDB{}
}

func (m *mapDB) Set(key, value string) error {
	m.m.Store(key, value)
	return nil
}
func (m *mapDB) Get(key string) (string, error) {
	v, ok := m.m.Load(key)
	if !ok {
		return "", fmt.Errorf("no value found for '%s'", key)
	}
	str, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("invalid type '%T'", v)
	}

	return str, nil
}
