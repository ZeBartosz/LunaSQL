package lexer

import (
	"fmt"
	"regexp"
)

type regexHandler func(lex *lexer, regex *regexp.Regexp)

type regexPattern struct {
	regex   *regexp.Regexp
	handler regexHandler
}

type lexer struct {
	patterns []regexPattern
	Tokens   []Token
	source   string
	pos      int
}

// updates the currect positiontokens[p.pos] in the source
func (lex *lexer) advanceN(n int) {
	lex.pos += n
}

// add new token to the array
func (lex *lexer) push(token Token) {
	lex.Tokens = append(lex.Tokens, token)
}

// checks what is at the current pos
func (lex *lexer) at() byte {
	return lex.source[lex.pos]
}

// checks if we are at the end of the source
func (lex *lexer) at_eof() bool {
	return lex.pos >= len(lex.source)
}

// checks how many bytes are left till the end
func (lex *lexer) remainder() string {
	return lex.source[lex.pos:]
}

func Tokenize(source string) []Token {
	lex := createLexer(source)

	// Iterate while we still have tokens
	for !lex.at_eof() {

		matched := false

		for _, pattern := range lex.patterns {
			loc := pattern.regex.FindStringIndex(lex.remainder())

			if loc != nil && loc[0] == 0 {
				pattern.handler(lex, pattern.regex)
				matched = true
				break
			}
		}

		// Can extend this print to show location and other stuff
		if !matched {
			panic(fmt.Sprintf("Lexer::Error -> unrecognized token near %s\n", lex.remainder()))
		}
	}

	lex.push(NewToken(EOF, "EOF"))
	return lex.Tokens
}

// default handling
func defaultHandler(kind TokenKind, value string) regexHandler {
	// pointer to lexer instance
	return func(lex *lexer, regex *regexp.Regexp) {
		lex.advanceN(len(value))
		lex.push(NewToken(kind, value))
	}
}

func createLexer(source string) *lexer {
	// & passing a pointer to a lexer instance
	return &lexer{
		pos:    0,
		source: source,
		Tokens: make([]Token, 0),
		patterns: []regexPattern{
			{regexp.MustCompile(`[a-zA-Z_][a-zA-Z0-9_]*`), symbolHandler},
			{regexp.MustCompile(`\d+(\.\d+)?`), numberHandler},
			{regexp.MustCompile(`'[^']*'`), stringHandler},
			{regexp.MustCompile(`\s+`), skipHandler},
			{regexp.MustCompile(`;`), defaultHandler(SEMICOLON, ";")},
			{regexp.MustCompile(`\*`), defaultHandler(STAR, "*")},
		},
	}
}

// handles numbers, pointer to a lexer instance, a compiled expression
func numberHandler(lex *lexer, regex *regexp.Regexp) {
	// finds first match of number pattern
	match := regex.FindString(lex.remainder())
	lex.push(NewToken(NUMBER, match))
	lex.advanceN(len(match))
}

// handles the stings
func stringHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remainder())
	stringLiteral := lex.remainder()[match[0]+1 : match[1]-1]

	lex.push(NewToken(STRING, stringLiteral))
	lex.advanceN(len(stringLiteral) + 2)
}

// handles reserved or indentifiersAdd commentMore actions
func symbolHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.remainder())
	kind := IDENTIFIER

	switch match {
	case "SELECT":
		kind = SELECT
	case "FROM":
		kind = FROM
	case "WHERE":
		kind = WHERE
	case "INSERT":
		kind = INSERT
	case "UPDATE":
		kind = UPDATE
	case "DELETE":
		kind = DELETE
	case "CREATE":
		kind = CREATE
	case "ALTER":
		kind = ALTER
	case "DROP":
		kind = DROP
	case "TABLE":
		kind = TABLE
	case "COLUMN":
		kind = COLUMN
	}

	lex.push(NewToken(kind, match))
	lex.advanceN(len(match))
}

// skips whitespace and comments
func skipHandler(lex *lexer, regex *regexp.Regexp) {

	match := regex.FindStringIndex(lex.remainder())

	lex.advanceN(match[1])

}
