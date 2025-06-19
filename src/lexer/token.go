package lexer

import "fmt"

type TokenKind int

const (
	EOF TokenKind = iota

	STRING
	NUMBER
	IDENTIFIER

	SELECT
	FROM
	WHERE
	INSERT
	UPDATE
	DELETE
	CREATE
	ALTER
	DROP
	TABLE
	COLUMN
	Value

	SEMICOLON

	STAR
)

type Token struct {
	Kind  TokenKind
	Value string
}

func (token Token) oneOfMany(expectedTokens ...TokenKind) bool {
	for _, expected := range expectedTokens {
		if expected == token.Kind {
			return true
		}
	}

	return false
}

// Helper method
func (token Token) Debug() {
	if token.oneOfMany(IDENTIFIER, NUMBER, STRING) {
		fmt.Printf("%s (%s)\n", TokenKindString(token.Kind), token.Value)
	} else {
		fmt.Printf("%s ()\n", TokenKindString(token.Kind))
	}
}

// Creates a token
func NewToken(kind TokenKind, value string) Token {
	return Token{
		kind, value,
	}
}

func TokenKindString(kind TokenKind) string {
	switch kind {
	case EOF:
		return "EOf"
	case SELECT:
		return "SELECT"
	case FROM:
		return "FROM"
	case WHERE:
		return "WHERE"
	case INSERT:
		return "INSERT"
	case UPDATE:
		return "UPDATE"
	case DELETE:
		return "DELETE"
	case CREATE:
		return "CREATE"
	case ALTER:
		return "ALTER"
	case DROP:
		return "DROP"
	case TABLE:
		return "TABLE"
	case COLUMN:
		return "COLUMN"
	case Value:
		return "Value"
	case STRING:
		return "STRING"
	case NUMBER:
		return "NUMBER"
	case IDENTIFIER:
		return "IDENTIFIER"
	case SEMICOLON:
		return "SEMICOLON"
	case STAR:
		return "STAR"
	default:
		return fmt.Sprintf("unknown(%d)", kind)
	}
}
