package ast

type BlockStmt struct {
	Body []Stmt
}

func (c BlockStmt) stmt() {}
