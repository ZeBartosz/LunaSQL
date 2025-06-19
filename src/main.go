package main

import (
	"fmt"
	"os"

	"github.com/ZeBartosz/miniSQL/src/lexer"
)

func main() {
	filePath := "./tests/01.sql"
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	tokens := lexer.Tokenize(string(bytes))

	for _, i := range tokens {
		i.Debug()
	}

	fmt.Println("File read successfully, no errors found.")
}
