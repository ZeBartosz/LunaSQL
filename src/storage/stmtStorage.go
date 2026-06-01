package storage

import (
	"github.com/ZeBartosz/LunaSQL/src/ast"
)

//	func generateExpressionStmt(stmt ast.ExpressionStmt, gen *Generator) {
//		gen.writeln(generateExpression(stmt.Expression) + ";\n")
//	}

func generateBlockStmt(block ast.BlockStmt, eng *Engine) {
	for _, stmt := range block.Body {
		generateStatement(stmt, eng)
	}
}

func generateCreateStmt(createStmt ast.CreateStmt, eng *Engine) (*Database, error) {
	return eng.CreateDatabase(createStmt.TableName)
}
