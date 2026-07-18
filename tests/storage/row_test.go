package storage_test

import (
	"strings"
	"testing"

	"github.com/ZeBartosz/LunaSQL/src/storage"
)

func TestParseDataTypeReturnsExpectedValues(t *testing.T) {
	tests := map[string]struct {
		raw  string
		typ  storage.ValueType
		want storage.Value
	}{
		"text value": {
			raw:  "hello luna",
			typ:  storage.TextType,
			want: storage.Value{Type: storage.TextType, Data: "hello luna"},
		},
		"empty text value": {
			raw:  "",
			typ:  storage.TextType,
			want: storage.Value{Type: storage.TextType, Data: ""},
		},
		"integer value": {
			raw:  "42",
			typ:  storage.IntType,
			want: storage.Value{Type: storage.IntType, Data: int64(42)},
		},
		"negative integer value": {
			raw:  "-7",
			typ:  storage.IntType,
			want: storage.Value{Type: storage.IntType, Data: int64(-7)},
		},
		"boolean true value": {
			raw:  "true",
			typ:  storage.BoolType,
			want: storage.Value{Type: storage.BoolType, Data: true},
		},
		"boolean false value": {
			raw:  "false",
			typ:  storage.BoolType,
			want: storage.Value{Type: storage.BoolType, Data: false},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := storage.ParseDataType(tt.raw, tt.typ)
			if err != nil {
				t.Fatalf("ParseDataType(%q, %q) returned unexpected error: %v", tt.raw, tt.typ, err)
			}

			if got != tt.want {
				t.Fatalf("ParseDataType(%q, %q) = %#v, want %#v", tt.raw, tt.typ, got, tt.want)
			}
		})
	}
}

func TestParseDataTypeReturnsErrors(t *testing.T) {
	tests := []struct {
		name       string
		raw        string
		typ        storage.ValueType
		wantErrSub string
	}{
		{
			name:       "invalid integer",
			raw:        "not-an-int",
			typ:        storage.IntType,
			wantErrSub: "invalid syntax",
		},
		{
			name:       "integer overflow",
			raw:        "2147483648",
			typ:        storage.IntType,
			wantErrSub: "value out of range",
		},
		{
			name:       "invalid boolean",
			raw:        "not-a-bool",
			typ:        storage.BoolType,
			wantErrSub: "invalid syntax",
		},
		{
			name:       "unsupported value type",
			raw:        "anything",
			typ:        storage.ValueType("FLOAT"),
			wantErrSub: "not supported",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := storage.ParseDataType(tt.raw, tt.typ)
			if err == nil {
				t.Fatalf("ParseDataType(%q, %q) = %#v, want error containing %q", tt.raw, tt.typ, got, tt.wantErrSub)
			}

			if !strings.Contains(err.Error(), tt.wantErrSub) {
				t.Fatalf("ParseDataType(%q, %q) error = %q, want substring %q", tt.raw, tt.typ, err.Error(), tt.wantErrSub)
			}
		})
	}
}
