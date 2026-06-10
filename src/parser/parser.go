package parser

import (
	"github.com/ZeBartosz/LunaSQL/src/ast"
	"github.com/ZeBartosz/LunaSQL/src/lexer"
)

type parser struct {
	tokens []lexer.Token
	pos    int
}

// create a parser instance
func createParser(tokens []lexer.Token) *parser {
	createTokenLookups()

	return &parser{
		tokens: tokens,
		pos:    0,
	}
}

func Parse(tokens []lexer.Token) (ast.Stmt, error) {
	// Create the parser instance
	p := createParser(tokens)

	Body := make([]ast.Stmt, 0)

	// Iterate until we reach the end of the file
	for p.hasToken() {
		stmt, err := parseStmt(p)
		if err != nil {
			return nil, err
		}
		Body = append(Body, stmt)
	}

	return ast.BlockStmt{
		Body: Body,
	}, nil
}
