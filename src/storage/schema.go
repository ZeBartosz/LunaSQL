package storage

import (
	"fmt"

	"github.com/ZeBartosz/LunaSQL/src/ast"
)

type ValueType string

const (
	IntType  ValueType = "INT"
	TextType ValueType = "TEXT"
	BoolType ValueType = "BOOL"
)

type Column struct {
	Name string    `json:"name"`
	Type ValueType `json:"type"`
}

func toStorageType(t ast.Type) (ValueType, error) {
	sym, ok := t.(ast.SymbolType)
	if !ok {
		return "", fmt.Errorf("unsupported type %T", t)
	}

	switch sym.Name {
	case "INT":
		return IntType, nil
	case "TEXT":
		return TextType, nil
	case "BOOL":
		return BoolType, nil
	default:
		return "", fmt.Errorf("unknown type %q", sym.Name)
	}
}
