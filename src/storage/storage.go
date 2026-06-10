// Package storage contains a deliberately small, file-backed storage layer.
//
// It is scaffolding for learning: tables are JSON files, rows are stored in
// memory while the engine runs, and every mutation rewrites the full table.
// This is simple (not fast), which makes it easy to replace piece by piece as
// you learn how real databases work.
package storage

import (
	"fmt"
	"os"

	"github.com/ZeBartosz/LunaSQL/src/ast"
)

// Engine owns the root folder where databases are stored.
type Engine struct {
	Root string
}

type Executor struct {
	engine   *Engine
	database *Database
	Table    *Table
}

func Storage(n ast.Stmt) error {
	var executor Executor
	engine, err := NewEngine("./database")
	if err != nil {
		return err
	}

	executor.engine = engine
	return generateStatement(n, &executor)
}

func generateStatement(stmt ast.Stmt, exec *Executor) error {
	switch n := stmt.(type) {
	case ast.BlockStmt:
		return generateBlockStmt(n, exec)
	case ast.CreateDatabaseStmt:
		return generateCreateDatabaseStmt(n, exec)
	case ast.CreateTableStmt:
		return generateCreateTableStmt(n, exec.database)
	case ast.InsertIntoTable:
		return insertToTableStmt(n, exec.database)
	case ast.SelectFromTable:
		return selectFromTableStmt(n, exec.database)
	default:
		return fmt.Errorf("unsupported statement type: %T", stmt)
	}
}

// NewEngine creates a storage engine rooted at root. The folder is created if needed.
func NewEngine(root string) (*Engine, error) {
	if err := os.MkdirAll(root, 0o755); err != nil {
		return nil, err
	}
	return &Engine{Root: root}, nil
}
