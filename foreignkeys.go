package main

import (
	"database/sql"
	"fmt"
	"regexp"
)

var (
	regFK = regexp.MustCompile(
		`FOREIGN KEY \(([^\)]+)\) REFERENCES ([^\(]+)\(([^\)]+)\)`,
	)
)

// ForeignKeysQuery is used to get all foreign key constraints for a given table.
const ForeignKeysQuery = `
SELECT pg_catalog.pg_get_constraintdef(r.oid, true)
FROM pg_catalog.pg_constraint r
WHERE r.conrelid = $1::regclass
  AND r.contype = 'f'
`

// ForeignKey represents a foreign key constraint for a specific table.
type ForeignKey struct {
	Field      string
	OtherTable string
	OtherField string
}

// Scan scans a ForeignKey from a string.
func (fk *ForeignKey) Scan(src interface{}) error {
	str, ok := src.(string)
	if !ok {
		return fmt.Errorf("expected string")
	}
	matches := regFK.FindStringSubmatch(str)
	if len(matches) != 4 {
		return fmt.Errorf("expected 4 matches, got %v", matches)
	}
	fk.Field = matches[1]
	fk.OtherTable = matches[2]
	fk.OtherField = matches[3]
	return nil
}

// ListForeignKeys fetches all foreign keys for a given table.
func ListForeignKeys(db *sql.DB, tableName string) ([]ForeignKey, error) {
	rows, err := db.Query(ForeignKeysQuery, tableName)
	if err != nil {
		return nil, err
	}
	// In case we get an error during scanning.
	// TODO(kdungs): Figure out if this is needed.
	defer rows.Close()
	keys := make([]ForeignKey, 0)
	for rows.Next() {
		var key ForeignKey
		if err := rows.Scan(&key); err != nil {
			return nil, err
		}
		keys = append(keys, key)
	}
	return keys, nil
}
