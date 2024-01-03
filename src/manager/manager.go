package manager

import (
	"fmt"

	"github.com/fiuskyws/thoth/src/storage"
)

type (
	Manager struct {
		dbs map[string]storage.API
	}
)

func NewManager() *Manager {
	return &Manager{
		dbs: map[string]storage.API{},
	}
}

func (m *Manager) CreateTable(name string) error {
	if _, ok := m.dbs[name]; ok {
		return fmt.Errorf("database '%s' already exists", name)
	}

	m.dbs[name] = storage.NewMapStorage()

	return nil
}

func (m *Manager) Set(dbName, key, value string) error {
	db, ok := m.dbs[dbName]
	if !ok {
		return fmt.Errorf("database '%s' does not exists", dbName)
	}

	return db.Set(key, value)
}

func (m *Manager) Get(dbName, key string) (string, error) {
	db, ok := m.dbs[dbName]
	if !ok {
		return "", fmt.Errorf("database '%s' does not exists", dbName)
	}

	return db.Get(key)
}
