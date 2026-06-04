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
	tokenKind := p.advance().Kind

	if tokenKind == lexer.DATABASE {
		databaseName := p.expect(lexer.IDENTIFIER)
		p.expect(lexer.SEMICOLON)

		return ast.CreateDatabaseStmt{
			DatabaseName: databaseName.Value,
		}, nil
	}

	if tokenKind == lexer.TABLE {
		tableName := p.expect(lexer.IDENTIFIER)
		p.expect(lexer.OPEN_PAREN)
		var columns []string

		for p.currentTokenKind() != lexer.CLOSE_PAREN {
			columnName := p.expect(lexer.IDENTIFIER).Value
			columns = append(columns, columnName)

			if p.currentTokenKind() == lexer.COMMA {
				p.advance()
			}
		}

		p.expect(lexer.CLOSE_PAREN)
		p.expect(lexer.SEMICOLON)

		return ast.CreateTableStmt{
			TableName: tableName.Value,
			Column:    columns,
		}, nil
	}

	panic("Create expects DATABASE or TABLE token after CREATE")
}

func parseInsertStmt(p *parser) (ast.Stmt, error) {
	p.advance()
	p.expect(lexer.INTO)

	tableName := p.expect(lexer.IDENTIFIER).Value
	insertValues := make(map[string]string)

	p.expect(lexer.OPEN_PAREN)
	for p.currentTokenKind() != lexer.CLOSE_PAREN {

		columnName := p.expect(lexer.IDENTIFIER).Value
		insertValues[columnName] = ""

		if p.currentTokenKind() == lexer.COMMA {
			p.advance()
		}
	}

	p.expect(lexer.CLOSE_PAREN)
	p.expect(lexer.VALUES)
	p.expect(lexer.OPEN_PAREN)

	for k, _ := range insertValues {
		valueToInsert := p.expect(lexer.IDENTIFIER).Value
		insertValues[k] = valueToInsert

		if p.currentTokenKind() == lexer.COMMA {
			p.advance()
		}
	}

	p.expect(lexer.CLOSE_PAREN)
	p.expect(lexer.SEMICOLON)

	return ast.InsertIntoTable{
		TableName: tableName,
		Insert:    insertValues,
	}, nil
}
func parseSelectStmt(p *parser) (ast.Stmt, error) {
	p.advance()
	var colums []string

	if p.currentTokenKind() == lexer.STAR {
		colums = append(colums, lexer.TokenKindString(lexer.STAR))
		p.advance()
	}

	p.expect(lexer.FROM)
	tableName := p.expect(lexer.IDENTIFIER).Value

	p.expect(lexer.SEMICOLON)

	return ast.SelectFromTable{
		TableName: tableName,
		Columns:   colums,
	}, nil
}
