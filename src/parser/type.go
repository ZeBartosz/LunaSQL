package parser

import (
	"fmt"

	"github.com/ZeBartosz/LunaSQL/src/ast"
	"github.com/ZeBartosz/LunaSQL/src/lexer"
)

type typeNudHandler func(p *parser) (ast.Type, error)
type typeLedHandler func(p *parser, left ast.Type, bp binding_power) (ast.Type, error)

type typeNudLookup map[lexer.TokenKind]typeNudHandler
type typeLedLookup map[lexer.TokenKind]typeLedHandler
type typeBPLookup map[lexer.TokenKind]binding_power

var typeBPLu = typeBPLookup{}
var typeNudLu = typeNudLookup{}
var typeLedLu = typeLedLookup{}

func typeLed(kind lexer.TokenKind, bp binding_power, ledFn typeLedHandler) {
	typeBPLu[kind] = bp
	typeLedLu[kind] = ledFn
}

func typeNud(kind lexer.TokenKind, nudFn typeNudHandler) {
	typeNudLu[kind] = nudFn
}

func createTokenTypeLookups() {
	typeNud(lexer.IDENTIFIER, parseSymbolType)
	typeNud(lexer.OPEN_BRACKET, parseArrayType)
}

func parseSymbolType(p *parser) (ast.Type, error) {
	tk, err := p.expect(lexer.IDENTIFIER)
	if err != nil {
		return nil, err
	}

	return ast.SymbolType{
		Name: tk.Value,
	}, nil
}

func parseArrayType(p *parser) (ast.Type, error) {
	p.advance()
	if _, err := p.expect(lexer.CLOSE_BRACKET); err != nil {
		return nil, err
	}

	underlyingType, err := parseType(p, defaultBP)
	if err != nil {
		return nil, err
	}
	return ast.ArrayType{
		Underlying: underlyingType,
	}, nil
}

func parseType(p *parser, bp binding_power) (ast.Type, error) {
	tokenKind := p.currentTokenKind()

	nudFn, exists := typeNudLu[tokenKind]

	if !exists {
		return nil, fmt.Errorf("TYPE_NUD HANDLER EXPECTED FOR TOKEN %s\n", lexer.TokenKindString(tokenKind))
	}

	left, err := nudFn(p)
	if err != nil {
		return nil, err
	}

	for typeBPLu[p.currentTokenKind()] > bp {
		tokenKind := p.currentTokenKind()
		ledFn, exists := typeLedLu[tokenKind]

		if !exists {
			panic(fmt.Sprintf("TYPE_LED HANDLER EXPECTED FOR TOKEN %s\n", lexer.TokenKindString(tokenKind)))
		}

		left, err = ledFn(p, left, typeBPLu[tokenKind])
		if err != nil {
			return nil, err
		}
	}

	return left, nil
}
