package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// Database is a collection of tables stored under Root/Name.
type Database struct {
	Name   string
	path   string
	Tables map[string]*Table
}

// Table is a tiny row-store. Each row is a map from column name to string value.
type Table struct {
	Name    string              `json:"name"`
	Columns []Column            `json:"columns"`
	Rows    []map[string]string `json:"rows"`
	path    string              `json:"-"`
}

// CreateDatabase creates a database directory and returns it.
func (e *Engine) CreateDatabase(name string) (*Database, error) {
	if name == "" {
		return nil, errors.New("database name cannot be empty")
	}

	path := filepath.Join(e.Root, name)
	if err := os.MkdirAll(path, 0o755); err != nil {
		return nil, err
	}

	fmt.Printf("Database Created: %s\n", name)
	return &Database{Name: name, path: path, Tables: map[string]*Table{}}, nil
}

// OpenDatabase loads an existing database directory.
func (e *Engine) OpenDatabase(name string) (*Database, error) {
	path := filepath.Join(e.Root, name)
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("%s is not a database directory", path)
	}

	db := &Database{Name: name, path: path, Tables: map[string]*Table{}}
	files, err := filepath.Glob(filepath.Join(path, "*.table.json"))
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		table, err := loadTable(file)
		if err != nil {
			return nil, err
		}
		db.Tables[table.Name] = table
	}
	return db, nil
}

// CreateTable creates an empty table with fixed column names.
func (db *Database) CreateTable(name string, columns []Column) (*Table, error) {
	if name == "" {
		return nil, errors.New("table name cannot be empty")
	}
	if len(columns) == 0 {
		return nil, errors.New("table must have at least one column")
	}
	if _, exists := db.Tables[name]; exists {
		return nil, fmt.Errorf("table %q already exists", name)
	}

	table := &Table{
		Name:    name,
		Columns: columns,
		Rows:    []map[string]string{},
		path:    filepath.Join(db.path, name+".table.json"),
	}
	if err := table.save(); err != nil {
		return nil, err
	}
	db.Tables[name] = table
	fmt.Printf("Table %s created\n", name)

	return table, nil
}

// Insert appends one row. The row must contain exactly the table columns.
func (t *Table) Insert(row map[string]string) error {
	clean := map[string]string{}
	for _, column := range t.Columns {
		value, ok := row[column.Name]
		if !ok {
			return fmt.Errorf("missing value for column %q", column.Name)
		}
		clean[column.Name] = value
	}
	for column := range row {
		if !containsColumn(t.Columns, column) {
			return fmt.Errorf("unknown column %q", column)
		}
	}

	t.Rows = append(t.Rows, clean)
	return t.save()
}

// SelectAll returns a copy of all rows in insertion order.
func (t *Table) SelectAll() []map[string]string {
	rows := make([]map[string]string, 0, len(t.Rows))
	for _, row := range t.Rows {
		copyRow := map[string]string{}
		for key, value := range row {
			copyRow[key] = value
		}
		rows = append(rows, copyRow)
	}
	return rows
}

func loadTable(path string) (*Table, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var table Table
	if err := json.Unmarshal(bytes, &table); err != nil {
		return nil, err
	}
	table.path = path
	return &table, nil
}

func (t *Table) save() error {
	bytes, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(t.path, bytes, 0o644)
}

func containsColumn(columns []Column, target string) bool {
	for _, column := range columns {
		if column.Name == target {
			return true
		}
	}
	return false
}
