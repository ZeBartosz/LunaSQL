package storage

import (
	"fmt"
	"path/filepath"

	"github.com/ZeBartosz/LunaSQL/src/ast"
)

//	func generateExpressionStmt(stmt ast.ExpressionStmt, gen *Generator) {
//		gen.writeln(generateExpression(stmt.Expression) + ";\n")
//	}

func generateBlockStmt(block ast.BlockStmt, exec *Executor) error {
	for _, stmt := range block.Body {
		if err := generateStatement(stmt, exec); err != nil {
			return err
		}
	}
	return nil
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
	table, ok := db.Tables[insertStmt.TableName]
	if !ok {
		return fmt.Errorf("table %q does not exist", insertStmt.TableName)
	}

	return table.Insert(insertStmt.Insert)
}
func selectFromTableStmt(selectStmt ast.SelectFromTable, db *Database) error {
	if table, ok := db.Tables[selectStmt.TableName]; ok {
		rows := table.SelectAll()

		fmt.Printf("\nTable: %s\n", selectStmt.TableName)
		for _, r := range rows {
			for k, v := range r {
				fmt.Println("Key:", k, "Value:", v)
			}
		}
	} else {
		tablePath := filepath.Join(db.path, selectStmt.TableName+".table.json")
		table, err := loadTable(tablePath)
		if err != nil {
			return err
		}
		rows := table.SelectAll()

		fmt.Printf("Table: %s\n", selectStmt.TableName)
		for _, r := range rows {
			for k, v := range r {
				fmt.Println("Key:", k, "Value:", v)
			}
		}
	}
	return nil
}
