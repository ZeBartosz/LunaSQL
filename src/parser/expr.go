package parser

import (
	"fmt"

	"github.com/ZeBartosz/LunaSQL/src/ast"
	"github.com/ZeBartosz/LunaSQL/src/lexer"
)

func parseExpr(p *parser, bp binding_power) (ast.Expr, error) {
	tokenKind := p.currentTokenKind()

	nud_fn, exists := nud_lu[tokenKind]

	if !exists {
		return nil, fmt.Errorf("NUD handler expected for token %s at position %d", lexer.TokenKindString(tokenKind), p.pos)
	}

	left, err := nud_fn(p)
	if err != nil {
		return nil, err
	}

	for bp_lu[p.currentTokenKind()] > bp {
		tokenKind = p.currentTokenKind()
		led_fn, exists := led_lu[tokenKind]

		if !exists {
			return nil, fmt.Errorf("LED handler expected for token %s at position %d", lexer.TokenKindString(tokenKind), p.pos)
		}

		left, err = led_fn(p, left, bp_lu[p.currentTokenKind()])
		if err != nil {
			return nil, err
		}
	}

	return left, nil
}
