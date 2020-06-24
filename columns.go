package main

import (
	"database/sql"
	"log"
)

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
	IsNullable bool
}

// ListColumns fetches all columns for a given table.
func ListColumns(db *sql.DB, tableName string) ([]Column, error) {
	rows, err := db.Query(ColumnsQuery, tableName)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatalf("%v", err)
		}
	}()

	cs := make([]Column, 0)
	for rows.Next() {
		var c Column
		var isnull string
		if err := rows.Scan(&c.Name, &c.Type, &isnull); err != nil {
			return nil, err
		}
		c.IsNullable = (isnull == "YES")
		cs = append(cs, c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return cs, nil
}
