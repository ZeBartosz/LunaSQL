package storage

import (
	"github.com/ZeBartosz/LunaSQL/src/ast"
)

//	func generateExpressionStmt(stmt ast.ExpressionStmt, gen *Generator) {
//		gen.writeln(generateExpression(stmt.Expression) + ";\n")
//	}

func generateBlockStmt(block ast.BlockStmt, exec *Executor) {
	for _, stmt := range block.Body {
		generateStatement(stmt, exec)
	}
}

func generateCreateDatabaseStmt(createStmt ast.CreateDatabaseStmt, exec *Executor) error {
	db, err := exec.engine.CreateDatabase(createStmt.DatabaseName)
	if err != nil {
		return err
	}

	exec.database = db
	return nil
}

func generateCreateTableStmt(tableStmt ast.CreateTableStmt, db *Database) error {
	table, err := db.CreateTable(tableStmt.TableName, tableStmt.Column)
	if err != nil {
		return err
	}

	db.Tables[table.Name] = table
	return nil
}

func insertToTableStmt(insertStmt ast.InsertIntoTable, db *Database) error {
	return db.Tables[insertStmt.TableName].Insert(insertStmt.Insert)
}
