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

func formatData(value Value, vType ValueType) (string, error) {
	switch vType {
	case TextType:
		text, ok := value.Data.(string)
		if !ok {
			return "", fmt.Errorf("expected TEXT data to be string, got %T", value.Data)
		}
		return text, nil
	case IntType:
		switch n := value.Data.(type) {
		case int:
			return fmt.Sprintf("%d", n), nil
		case int64:
			return fmt.Sprintf("%d", n), nil
		case float64: // JSON unmarshals numbers into float64 when the target is any.
			return fmt.Sprintf("%.0f", n), nil
		default:
			return "", fmt.Errorf("expected INT data to be numeric, got %T", value.Data)
		}
	case BoolType:
		boolean, ok := value.Data.(bool)
		if !ok {
			return "", fmt.Errorf("expected BOOL data to be bool, got %T", value.Data)
		}
		return fmt.Sprintf("%t", boolean), nil
	default:
		return "", fmt.Errorf("this type: %s is not supported", vType)
	}
}
