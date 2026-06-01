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

func Storage(n ast.Stmt) {
	engine, err := NewEngine("./database")
	if err != nil {
		panic(err)
	}

	generateStatement(n, engine)
}

func generateStatement(stmt ast.Stmt, eng *Engine) {
	switch n := stmt.(type) {
	case ast.BlockStmt:
		generateBlockStmt(n, eng)
	case ast.CreateStmt:
		generateCreateStmt(n, eng)
	default:
		fmt.Sprintf("// Unsupported statement type: %T\n", stmt)
	}
}

// NewEngine creates a storage engine rooted at root. The folder is created if needed.
func NewEngine(root string) (*Engine, error) {
	if err := os.MkdirAll(root, 0o755); err != nil {
		return nil, err
	}
	return &Engine{Root: root}, nil
}
