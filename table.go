package main

import (
	"database/sql"
	"log"
)

// TableNamesQuery is used to get all table names for the public schema.
const TableNamesQuery = `
SELECT table_name
FROM information_schema.tables
WHERE table_type = 'BASE TABLE'
AND table_schema = 'public'
ORDER BY table_name ASC
`

// Table represents a database table.
type Table struct {
	Name        string
	Columns     []Column
	ForeignKeys []ForeignKey
}

// LoadColumns loads all columns for the given table.
func (t *Table) LoadColumns(db *sql.DB) error {
	cols, err := ListColumns(db, t.Name)
	if err != nil {
		return err
	}
	t.Columns = cols
	return nil
}

// LoadForeignKeys loads all foreign key constraints for the given table.
func (t *Table) LoadForeignKeys(db *sql.DB) error {
	keys, err := ListForeignKeys(db, t.Name)
	if err != nil {
		return err
	}
	t.ForeignKeys = keys
	return nil
}

// ListTablesOptions contains options for loading table information from the
// database.
type ListTablesOptions struct {
	LoadColumns     bool
	LoadForeignKeys bool
}

// ListTables lists all tables in the database (public schema).
// Depending on the ListTableOptions, additional information (columns, foreign
// keys) is loaded.
func ListTables(db *sql.DB, options ListTablesOptions) ([]Table, error) {
	rows, err := db.Query(TableNamesQuery)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatalf("%v", err)
		}
	}()

	tables := make([]Table, 0)
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		t := Table{Name: name}
		if options.LoadColumns {
			if err := t.LoadColumns(db); err != nil {
				return nil, err
			}
		}
		if options.LoadForeignKeys {
			if err := t.LoadForeignKeys(db); err != nil {
				return nil, err
			}
		}
		tables = append(tables, t)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tables, nil
}
