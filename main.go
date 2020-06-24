package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var (
	pgHost        = flag.String("host", "localhost", "Postgres hostname")
	pgPort        = flag.Int("port", 5432, "Postgres port")
	pgUser        = flag.String("user", "postgres", "Postgres username")
	pgPass        = flag.String("pass", "", "Postgres password")
	pgDB          = flag.String("db", "", "Postgres database")
	pgNoSSL       = flag.Bool("no-ssl", false, "Disable ssl")
	showColumns   = flag.Bool("show-columns", false, "whether to show columns for each table")
	showRelations = flag.Bool("show-relations", true, "whether to show relationships between tables (based on foreign keys)")
)

func connect() (*sql.DB, error) {
	connstr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s",
		*pgHost,
		*pgPort,
		*pgUser,
		*pgPass,
		*pgDB,
	)
	if *pgNoSSL {
		connstr += " sslmode=disable"
	}
	return sql.Open("postgres", connstr)
}

func main() {
	flag.Parse()

	if *pgDB == "" {
		panic("You need to provide the database name using -db <name>")
	}

	db, err := connect()
	if err != nil {
		log.Fatalf("%v", err)
	}

	tables, err := ListTables(db, ListTablesOptions{
		LoadColumns:     *showColumns,
		LoadForeignKeys: *showRelations,
	})
	if err != nil {
		log.Fatalf("%v", err)
	}

	if err := RenderGraphviz(os.Stdout, tables); err != nil {
		log.Fatalf("%v", err)
	}
}
