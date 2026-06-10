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
	DatabaseName string
}

func (n CreateDatabaseStmt) stmt() {}

type CreateTableStmt struct {
	TableName string
	Column    []string
}

func (n CreateTableStmt) stmt() {}

type InsertIntoTable struct {
	TableName string
	Insert    map[string]string
}

func (n InsertIntoTable) stmt() {}

type SelectFromTable struct {
	TableName string
	Columns   []string
}

func (n SelectFromTable) stmt() {}
