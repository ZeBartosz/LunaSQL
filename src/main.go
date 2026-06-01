package main

import (
	"fmt"
	"os"

	"github.com/ZeBartosz/LunaSQL/src/lexer"
	"github.com/ZeBartosz/LunaSQL/src/parser"
	"github.com/ZeBartosz/LunaSQL/src/storage"
	"github.com/sanity-io/litter"
)

func main() {
	filePath := "./tests/01.sql"
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	tokens := lexer.Tokenize(string(bytes))

	fmt.Println("\n--- Tokens ---")
	for _, i := range tokens {
		i.Debug()
	}

	ast, err := parser.Parse(tokens)
	if err != nil {
		fmt.Printf("Error parsing ast %v\n", err)
	}

	fmt.Println("\n--- Abstract Syntax Tree ---")
	litter.Dump(ast)

	fmt.Println("\n--- StorageAction ---")
	storage.Storage(ast)

	fmt.Println("\nFile read successfully, no errors found.")
}
