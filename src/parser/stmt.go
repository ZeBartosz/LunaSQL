package parser

import (
	"github.com/ZeBartosz/LunaSQL/src/ast"
	"github.com/ZeBartosz/LunaSQL/src/lexer"
)

func parseStmt(p *parser) (ast.Stmt, error) {
	stmt_fn, exists := stmt_lu[p.currentTokenKind()]

	if exists {
		return stmt_fn(p)
	}

	expression, err := parseExprStmt(p)
	if err != nil {
		return nil, err
	}

	return expression, nil
}

func parseExprStmt(p *parser) (ast.Stmt, error) {
	expression := parseExpr(p, defalt_bp)

	return ast.ExprStmt{
		Expression: expression,
	}, nil
}

func parseCreateStmt(p *parser) (ast.Stmt, error) {
	p.expect(lexer.CREATE)
	p.expect(lexer.DATABASE)

	databaseName := p.expect(lexer.IDENTIFIER)

	p.expect(lexer.SEMICOLON)

	return ast.CreateStmt{
		TableName: databaseName.Value,
	}, nil
}
