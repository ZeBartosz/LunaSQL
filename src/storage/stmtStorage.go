package storage

import (
	"github.com/ZeBartosz/LunaSQL/src/ast"
)

//	func generateExpressionStmt(stmt ast.ExpressionStmt, gen *Generator) {
//		gen.writeln(generateExpression(stmt.Expression) + ";\n")
//	}

func generateBlockStmt(block ast.BlockStmt, exec Executor) {
	for _, stmt := range block.Body {
		generateStatement(stmt, exec)
	}
}

func generateCreateStmt(createStmt ast.CreateDatabaseStmt, eng *Engine) (*Database, error) {
	return eng.CreateDatabase(createStmt.TableName)
}
