package parser

import (
	"fmt"

	"github.com/ZeBartosz/LunaSQL/src/lexer"
)

// returns current token
func (p *parser) currentToken() lexer.Token {
	if p.pos >= len(p.tokens) {
		return lexer.Token{Kind: lexer.EOF}
	}
	return p.tokens[p.pos]
}

// we get the current kind of the token
func (p *parser) currentTokenKind() lexer.TokenKind {
	return p.currentToken().Kind
}

// advance to the next token, returns current token
func (p *parser) advance() lexer.Token {
	tk := p.currentToken()
	if p.pos < len(p.tokens) {
		p.pos++
	}
	return tk
}

// Checks if there are more token to parse through
func (p *parser) hasToken() bool {
	return p.pos < len(p.tokens) && p.currentTokenKind() != lexer.EOF
}

// checks if the expected tokenKind is the same a the parse kind
func (p *parser) expectError(expectedKind lexer.TokenKind, errStr string) (lexer.Token, error) {
	token := p.currentToken()
	kind := token.Kind
	value := token.Value

	if kind != expectedKind {
		err := fmt.Errorf("expected: %s but received: %s, value: %s instead, at position %d", lexer.TokenKindString(expectedKind), lexer.TokenKindString(kind), value, p.pos)
		if errStr != "" {
			err = fmt.Errorf("%s: %w", errStr, err)
		}
		return token, err
	}

	return p.advance(), nil
}

// checks if the token is the one we expect
func (p *parser) expect(expectedKind lexer.TokenKind) lexer.Token {
	tok, err := p.expectError(expectedKind, "")
	if err != nil {
		panic(err)
	}
	return tok
}
