package parser

import (
	"github.com/ZeBartosz/LunaSQL/src/ast"
	"github.com/ZeBartosz/LunaSQL/src/lexer"
)

// binding_power represents the precedence of different operators.
// Higher values mean higher precedence.
type binding_power int

// Enum defining the binding power of various operators and expressions.
const (
	defalt_bp      binding_power = iota // Default (lowest precedence)
	comma                               // Comma operator (e.g., function arguments)
	assignment                          // Assignment operators (=, +=, etc.)
	logical                             // Logical operators (&&, ||)
	relational                          // Comparison operators (==, <, >, etc.)
	additive                            // Addition and subtraction (+, -)
	multiplicative                      // Multiplication, division, modulo (*, /, %)
	unary                               // Unary operators (!, -)
	call                                // Function calls (e.g., foo())
	member                              // Object property access (e.g., obj.prop)
	primary                             // Literals, variables (highest precedence)
)

// Function types used to handle parsing of different language constructs
type (
	stmt_handler func(p *parser) (ast.Stmt, error)
	nud_handler  func(p *parser) ast.Expr
	led_handler  func(p *parser, left ast.Expr, bp binding_power) ast.Expr
)

// Lookup tables that map token kinds to their respective handlers
type (
	stmt_lookup map[lexer.TokenKind]stmt_handler
	nud_lookup  map[lexer.TokenKind]nud_handler
	led_lookup  map[lexer.TokenKind]led_handler
	bp_lookup   map[lexer.TokenKind]binding_power
)

// Global lookup tables that will store parsing rules for different token kinds.
var (
	bp_lu   = bp_lookup{}   // Stores operator precedence values.
	nud_lu  = nud_lookup{}  // Stores handlers for prefix expressions and literals.
	led_lu  = led_lookup{}  // Stores handlers for infix and postfix expressions.
	stmt_lu = stmt_lookup{} // Stores handlers for full statements.
)

// Registers an infix or postfix operator in the lookup tables.
func led(kind lexer.TokenKind, bp binding_power, led_fn led_handler) {
	bp_lu[kind] = bp
	led_lu[kind] = led_fn
}

// Registers a prefix operator or literal in the lookup tables.
func nud(kind lexer.TokenKind, nud_fn nud_handler) {
	nud_lu[kind] = nud_fn
}

// Registers a statement type in the lookup tables.
func stmt(kind lexer.TokenKind, stmt_fn stmt_handler) {
	bp_lu[kind] = defalt_bp
	stmt_lu[kind] = stmt_fn
}

func createTokenLookups() {
	// led(lexer.ASSIGNMENT, assignment, parse_assignment_expr)
	// led(lexer.PLUS_EQUALS, assignment, parse_assignment_expr)
	// led(lexer.MINUS_EQUALS, assignment, parse_assignment_expr)
	// TODO: *= /= %=

	// Logical
	// led(lexer.AND, logical, parse_binary_expr)
	// led(lexer.OR, logical, parse_binary_expr)
	// led(lexer.DOT_DOT, logical, parse_binary_expr)

	// Relational
	// led(lexer.LESS, relational, parse_binary_expr)
	// led(lexer.LESS_EQUALS, relational, parse_binary_expr)
	// led(lexer.GREATER, relational, parse_binary_expr)
	// led(lexer.GREATER_EQUALS, relational, parse_binary_expr)
	// led(lexer.EQUALS, relational, parse_binary_expr)
	// led(lexer.NOT_EQUALS, relational, parse_binary_expr)

	// Additive & Multiplicative
	// led(lexer.PLUS, additive, parse_binary_expr)
	// led(lexer.DASH, additive, parse_binary_expr)
	// led(lexer.STAR, multiplicative, parse_binary_expr)
	// led(lexer.SLASH, multiplicative, parse_binary_expr)
	// led(lexer.PERCENT, multiplicative, parse_binary_expr)

	// Literals $ Symbols
	// nud(lexer.NUMBER, parse_primary_expr)
	// nud(lexer.STRING, parse_primary_expr)
	// nud(lexer.IDENTIFIER, parse_primary_expr)
	// nud(lexer.DASH, parse_prefix_expr)
	// nud(lexer.OPEN_PAREN, parse_grouping_expr)

	// Statements
	// stmt(lexer.LET, parse_var_decl_stmt)
	// stmt(lexer.VAR, parse_var_decl_stmt)
	// stmt(lexer.CONST, parse_var_decl_stmt)
	// stmt(lexer.IMPORT, parse_import_stmt)
	// stmt(lexer.FUNC, parse_func_stmt)
	// stmt(lexer.RETURN, parse_return_stmt)
	stmt(lexer.CREATE, parseCreateStmt)
}
