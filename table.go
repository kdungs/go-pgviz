package main

import "database/sql"

// TableNamesQuery is used to get all table names for the public schema.
const TableNamesQuery = `
SELECT table_name
FROM information_schema.tables
WHERE table_type = 'BASE TABLE'
AND table_schema = 'public'
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

// Load loads all additional information for the given table.
func (t *Table) Load(db *sql.DB) error {
	if err := t.LoadColumns(db); err != nil {
		return err
	}
	if err := t.LoadForeignKeys(db); err != nil {
		return err
	}
	return nil
}

// ListTables lists all tables in the database (public schema).
// If full is set to true, additional information like columns and foreign keys
// is loaded as well.
func ListTables(db *sql.DB, full bool) ([]Table, error) {
	rows, err := db.Query(TableNamesQuery)
	if err != nil {
		return nil, err
	}
	// In case we get an error during scanning.
	// TODO(kdungs): Figure out if this is needed.
	defer rows.Close()

	tables := make([]Table, 0)
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		t := Table{Name: name}
		if full {
			if err := t.Load(db); err != nil {
				return nil, err
			}
		}
		tables = append(tables, t)
	}
	return tables, nil
}
