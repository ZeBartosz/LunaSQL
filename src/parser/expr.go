package parser

import (
	"fmt"

	"github.com/ZeBartosz/LunaSQL/src/ast"
	"github.com/ZeBartosz/LunaSQL/src/lexer"
)

func parseExpr(p *parser, bp binding_power) ast.Expr {
	tokenKind := p.currentTokenKind()

	nud_fn, exists := nud_lu[tokenKind]

	if !exists {
		panic(fmt.Sprintf("NUD HANDLER EXPECTED FOR TOKEN %s\n", lexer.TokenKindString(tokenKind)))
	}

	left := nud_fn(p)
	for bp_lu[p.currentTokenKind()] > bp {
		tokenKind = p.currentTokenKind()
		led_fn, exists := led_lu[tokenKind]

		if !exists {
			panic(fmt.Sprintf("LED HANDLER EXPECTED FOR TOKEN %s\n", lexer.TokenKindString(tokenKind)))
		}

		left = led_fn(p, left, bp_lu[p.currentTokenKind()])
	}

	return left
}
