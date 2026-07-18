package storage

import (
	"fmt"
	"strconv"
)

type Value struct {
	Type ValueType
	Data any
}

type Row struct {
	Values map[string]Value
}

func ParseDataType(rawData string, vType ValueType) (Value, error) {
	switch vType {
	case TextType:
		return Value{Type: vType, Data: rawData}, nil
	case IntType:
		parsedInt, err := strconv.ParseInt(rawData, 10, 32)
		if err != nil {
			return Value{}, err
		}
		return Value{Type: vType, Data: parsedInt}, nil

	case BoolType:
		parsedBool, err := strconv.ParseBool(rawData)
		if err != nil {
			return Value{}, err
		}
		return Value{Type: vType, Data: parsedBool}, nil
	default:
		return Value{}, fmt.Errorf("this type: %s is not supported", vType)
	}
}
