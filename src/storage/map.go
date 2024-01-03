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

func NewMapStorage() API {
	return &mapDB{}
}

const (
	invalidTypeErr  = "invalid type '%T'"
	noValueFoundErr = "no value found for '%s'"
	fieldNotSetErr  = "field '%s' is not set"
)

func (m *mapDB) Set(key, value string) error {
	if key == "" {
		return fmt.Errorf(fieldNotSetErr, "key")
	}
	if value == "" {
		return fmt.Errorf(fieldNotSetErr, "value")
	}
	m.m.Store(key, value)
	return nil
}

func (m *mapDB) Get(key string) (string, error) {
	if key == "" {
		return "", fmt.Errorf(fieldNotSetErr, "key")
	}
	v, ok := m.m.Load(key)
	if !ok {
		return "", fmt.Errorf(noValueFoundErr, key)
	}
	str, ok := v.(string)
	if !ok {
		return "", fmt.Errorf(invalidTypeErr, v)
	}

	return str, nil
}

func (m *mapDB) Delete(key string) error {
	if _, err := m.Get(key); err != nil {
		return err
	}
	m.m.Delete(key)
	return nil
}
