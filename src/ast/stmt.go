package ast

type BlockStmt struct {
	Body []Stmt
}

func (n BlockStmt) stmt() {}

type ExprStmt struct {
	Expression Expr
}

func (n ExprStmt) stmt() {}

type CreateDatabaseStmt struct {
	TableName string
}

func (n CreateDatabaseStmt) stmt() {}
