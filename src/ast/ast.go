// Package ast provides tools for constructing, traversing, and analyzing
// abstract syntax trees (ASTs) in Go.
//
// This package defines a set of data structures that represent the syntax
// and provide mechanisms for introspection and manipulation of AST nodes.
package ast

type Stmt interface {
	stmt()
}

type Expr interface {
	expr()
}

type Type interface {
	_type()
}
