package parser

import (
	"fmt"

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
	expression, err := parseExpr(p, defalt_bp)
	if err != nil {
		return nil, err
	}

	return ast.ExprStmt{
		Expression: expression,
	}, nil
}

func parseCreateStmt(p *parser) (ast.Stmt, error) {
	if _, err := p.expect(lexer.CREATE); err != nil {
		return nil, err
	}
	tokenKind := p.advance().Kind

	if tokenKind == lexer.DATABASE {
		databaseName, err := p.expect(lexer.IDENTIFIER)
		if err != nil {
			return nil, err
		}

		if _, err := p.expect(lexer.SEMICOLON); err != nil {
			return nil, err
		}

		return ast.CreateDatabaseStmt{
			DatabaseName: databaseName.Value,
		}, nil
	}

	if tokenKind == lexer.TABLE {
		tableName, err := p.expect(lexer.IDENTIFIER)
		if err != nil {
			return nil, err
		}

		if _, err := p.expect(lexer.OPEN_PAREN); err != nil {
			return nil, err
		}
		var columns []string

		for p.currentTokenKind() != lexer.CLOSE_PAREN {
			columnName, err := p.expect(lexer.IDENTIFIER)
			if err != nil {
				return nil, err
			}
			columns = append(columns, columnName.Value)

			if p.currentTokenKind() == lexer.COMMA {
				p.advance()
			}
		}

		if _, err := p.expect(lexer.CLOSE_PAREN); err != nil {
			return nil, err
		}
		if len(columns) == 0 {
			return nil, fmt.Errorf("CREATE TABLE requires at least one column")
		}
		if _, err := p.expect(lexer.SEMICOLON); err != nil {
			return nil, err
		}

		return ast.CreateTableStmt{
			TableName: tableName.Value,
			Column:    columns,
		}, nil
	}

	return nil, fmt.Errorf("Create expects DATABASE or TABLE token after CREATE")
}

func parseInsertStmt(p *parser) (ast.Stmt, error) {
	p.advance()
	if _, err := p.expect(lexer.INTO); err != nil {
		return nil, err
	}

	tableName, err := p.expect(lexer.IDENTIFIER)
	if err != nil {
		return nil, err
	}

	var columns []string

	if _, err := p.expect(lexer.OPEN_PAREN); err != nil {
		return nil, err
	}

	for p.currentTokenKind() != lexer.CLOSE_PAREN {

		columnName, err := p.expect(lexer.IDENTIFIER)
		if err != nil {
			return nil, err
		}
		columns = append(columns, columnName.Value)

		if p.currentTokenKind() == lexer.COMMA {
			p.advance()
		}
	}

	if _, err := p.expect(lexer.CLOSE_PAREN); err != nil {
		return nil, err
	}
	if _, err := p.expect(lexer.VALUES); err != nil {
		return nil, err
	}
	if _, err := p.expect(lexer.OPEN_PAREN); err != nil {
		return nil, err
	}

	var values []string

	for p.currentTokenKind() != lexer.CLOSE_PAREN {

		value, err := p.expect(lexer.IDENTIFIER)
		if err != nil {
			return nil, err
		}
		values = append(values, value.Value)

		if p.currentTokenKind() == lexer.COMMA {
			p.advance()
		}
	}

	row := map[string]string{}

	if len(columns) != len(values) {
		return nil, fmt.Errorf("INSERT column count (%d) does not match value count (%d)", len(columns), len(values))
	}

	for i, column := range columns {
		row[column] = values[i]
	}

	if _, err := p.expect(lexer.CLOSE_PAREN); err != nil {
		return nil, err
	}
	if _, err := p.expect(lexer.SEMICOLON); err != nil {
		return nil, err
	}

	return ast.InsertIntoTable{
		TableName: tableName.Value,
		Insert:    row,
	}, nil
}

func parseSelectStmt(p *parser) (ast.Stmt, error) {
	p.advance()
	var columns []string

	if p.currentTokenKind() == lexer.STAR {
		columns = append(columns, lexer.TokenKindString(lexer.STAR))
		p.advance()
	}

	if _, err := p.expect(lexer.FROM); err != nil {
		return nil, err
	}
	tableName, err := p.expect(lexer.IDENTIFIER)
	if err != nil {
		return nil, err
	}
	if _, err := p.expect(lexer.SEMICOLON); err != nil {
		return nil, err
	}

	return ast.SelectFromTable{
		TableName: tableName.Value,
		Columns:   columns,
	}, nil
}
