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
}

func Storage(n ast.Stmt) {
	var executor Executor
	engine, err := NewEngine("./database")
	if err != nil {
		panic(err)
	}

	executor.engine = engine
	generateStatement(n, &executor)
}

func generateStatement(stmt ast.Stmt, exec *Executor) {
	switch n := stmt.(type) {
	case ast.BlockStmt:
		generateBlockStmt(n, exec)
	case ast.CreateDatabaseStmt:
		db, err := generateCreateDatabaseStmt(n, exec.engine)
		if err != nil {
			panic(err)
		}

		exec.database = db
	case ast.CreateTableStmt:
		if exec.database == nil {
			panic("Database not set")
		}

		err := generateCreateTableStmt(n, exec.database)
		if err != nil {
			panic(err)
		}
	default:
		fmt.Printf("// Unsupported statement type: %T\n", stmt)
	}
}

// NewEngine creates a storage engine rooted at root. The folder is created if needed.
func NewEngine(root string) (*Engine, error) {
	if err := os.MkdirAll(root, 0o755); err != nil {
		return nil, err
	}
	return &Engine{Root: root}, nil
}
