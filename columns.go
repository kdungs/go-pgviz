package main

import "database/sql"

// ColumnsQuery is used to get all columns for a specific table.
const ColumnsQuery = `
SELECT column_name, data_type, is_nullable
FROM information_schema.columns
WHERE table_name = $1
`

// Column represents a column within a table.
type Column struct {
	Name       string
	Type       string
	IsNullable string
}

// ListColumns fetches all columns for a given table.
func ListColumns(db *sql.DB, tableName string) ([]Column, error) {
	rows, err := db.Query(ColumnsQuery, tableName)
	if err != nil {
		return nil, err
	}
	// In case we get an error during scanning.
	// TODO(kdungs): Figure out if this is needed.
	defer rows.Close()

	cs := make([]Column, 0)
	for rows.Next() {
		var c Column
		if err := rows.Scan(&c.Name, &c.Type, &c.IsNullable); err != nil {
			return nil, err
		}
		cs = append(cs, c)
	}
	return cs, nil
}
