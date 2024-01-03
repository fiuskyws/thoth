package manager

import (
	"fmt"

	"github.com/fiuskyws/thoth/src/storage"
)

type (
	Manager struct {
		tables map[string]storage.API
	}
)

func NewManager() *Manager {
	return &Manager{
		tables: map[string]storage.API{},
	}
}

func (m *Manager) CreateTable(name string) error {
	if _, ok := m.tables[name]; ok {
		return fmt.Errorf("database '%s' already exists", name)
	}

	m.tables[name] = storage.NewMapStorage()

	return nil
}

func (m *Manager) GetTables() []string {
	tables := []string{}

	for k := range m.tables {
		tables = append(tables, k)
	}

	return tables
}

func (m *Manager) Set(tableName, key, value string) error {
	table, ok := m.tables[tableName]
	if !ok {
		return fmt.Errorf("database '%s' does not exists", tableName)
	}

	return table.Set(key, value)
}

func (m *Manager) Get(tableName, key string) (string, error) {
	table, ok := m.tables[tableName]
	if !ok {
		return "", fmt.Errorf("database '%s' does not exists", tableName)
	}

	return table.Get(key)
}

func (m *Manager) Delete(tableName, key string) error {
	table, ok := m.tables[tableName]
	if !ok {
		return fmt.Errorf("database '%s' does not exists", tableName)
	}

	return table.Delete(key)
}
