package storage

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateDatabase(t *testing.T) {
	tests := map[string]struct {
		name      string
		tables    map[string]*Table
		expectErr error
	}{
		"create a database": {
			name:   "database1",
			tables: map[string]*Table{},
		},
		"error without name param": {
			name:      "",
			expectErr: ErrEmptyDatabaseName,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			dirName := t.TempDir()

			storage, err := NewEngine(dirName)
			require.NoError(t, err)

			database, err := storage.CreateDatabase(test.name)
			if test.expectErr != nil {
				require.ErrorIs(t, err, test.expectErr)
				require.Nil(t, database)
				return
			}

			require.NoError(t, err)

			require.Equal(t, filepath.Join(dirName, test.name), database.path)
			require.Equal(t, test.name, database.Name)
			require.Equal(t, test.tables, database.Tables)
		})
	}
}

func TestOpenDatabase(t *testing.T) {
	tests := map[string]struct {
		name           string
		createDatabase string
		tables         map[string][]string
		rows           map[string][]map[string]string
		files          map[string]string
		expectErr      error
	}{
		"successfully open a database with tables": {
			name:           "db1",
			createDatabase: "db1",
			tables:         map[string][]string{"users": {"id", "name"}},
			rows:           map[string][]map[string]string{"users": {{"id": "1", "name": "luna"}}},
		},
		"successfully open an empty database": {
			name:           "db1",
			createDatabase: "db1",
		},
		"ignores files that are not tables": {
			name:           "db1",
			createDatabase: "db1",
			tables:         map[string][]string{"users": {"id"}},
			files: map[string]string{
				"db1/notes.txt":  "hi",
				"db1/users.json": "{}",
			},
		},
		"error database does not exist": {
			name:      "db1",
			expectErr: fs.ErrNotExist,
		},
		"error database is a file": {
			name:      "db1",
			files:     map[string]string{"db1": "x"},
			expectErr: ErrNotDatabase,
		},
		"error table file is corrupt": {
			name:           "db1",
			createDatabase: "db1",
			files:          map[string]string{"db1/bad.table.json": "{not json"},
			expectErr:      ErrCorruptTable,
		},
		"error without name param": {
			name:      "",
			expectErr: ErrEmptyDatabaseName,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			dirName := t.TempDir()

			storage, err := NewEngine(dirName)
			require.NoError(t, err)

			if test.createDatabase != "" {
				database, err := storage.CreateDatabase(test.createDatabase)
				require.NoError(t, err)

				for tableName, columns := range test.tables {
					table, err := database.CreateTable(tableName, columns)
					require.NoError(t, err)

					for _, row := range test.rows[tableName] {
						require.NoError(t, table.Insert(row))
					}
				}
			}

			for path, content := range test.files {
				require.NoError(t, os.WriteFile(filepath.Join(dirName, path), []byte(content), 0o644))
			}

			opened, err := storage.OpenDatabase(test.name)
			if test.expectErr != nil {
				require.ErrorIs(t, err, test.expectErr)
				require.Nil(t, opened)
				return
			}

			require.NoError(t, err)

			require.Equal(t, filepath.Join(dirName, test.name), opened.path)
			require.Equal(t, test.name, opened.Name)
			require.NotNil(t, opened.Tables)
			require.Len(t, opened.Tables, len(test.tables))

			for tableName, columns := range test.tables {
				table, ok := opened.Tables[tableName]
				require.True(t, ok, "table %q not loaded", tableName)
				require.Equal(t, tableName, table.Name)
				require.Equal(t, columns, table.Columns)
				require.Equal(t, filepath.Join(opened.path, tableName+".table.json"), table.path)

				expectRows := test.rows[tableName]
				if expectRows == nil {
					expectRows = []map[string]string{}
				}
				require.Equal(t, expectRows, table.Rows)
			}
		})
	}
}

func TestCreateTable(t *testing.T) {
	tests := map[string]struct {
		name               string
		columns            []string
		existColumn        []string
		expectErr          error
		expectSameTableErr error
	}{
		"sucessfully create a table": {
			name:    "table1",
			columns: []string{"column1"},
		},
		"error table name not provided": {
			name:      "",
			expectErr: ErrEmptyTableName,
		},
		"error not enough columns provided": {
			name:      "table1",
			columns:   []string{},
			expectErr: ErrEmptyColumns,
		},
		"error when creating column with the same name": {
			name:      "table1",
			columns:   []string{"column1", "column1"},
			expectErr: ErrDuplicateColumn,
		},
		"error table already exists": {
			name:               "table1",
			columns:            []string{"column1"},
			expectSameTableErr: ErrTableExists,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			dirName := t.TempDir()

			storage, err := NewEngine(dirName)
			require.NoError(t, err)

			database, err := storage.CreateDatabase("database")
			require.NoError(t, err)

			tables, err := database.CreateTable(test.name, test.columns)
			if test.expectErr != nil {
				require.ErrorIs(t, err, test.expectErr)
				require.Nil(t, tables)
				return
			}

			if test.expectSameTableErr != nil {
				tables, err := database.CreateTable(test.name, test.columns)
				require.ErrorIs(t, err, test.expectSameTableErr)
				require.Nil(t, tables)
				return
			}

			require.NoError(t, err)

			require.Equal(t, filepath.Join(database.path, test.name+".table.json"), tables.path)
			require.Equal(t, test.name, tables.Name)
			require.Equal(t, test.columns, tables.Columns)
		})
	}
}
