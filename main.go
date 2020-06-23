package main

import (
	"database/sql"
	"flag"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	host = flag.String("host", "localhost", "Postgres hostname")
	port = flag.Int("port", 5432, "Postgres port")
	user = flag.String("user", "postgres", "Postgres username")
	pass = flag.String("pass", "", "Postgres password")
	db   = flag.String("db", "", "Postgres database")
)

func main() {
	flag.Parse()
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s",
		*host,
		*port,
		*user,
		*pass,
		*db,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	tables, err := ListTables(db, true)
	if err != nil {
		panic(err)
	}

	graph := BuildDependencies(tables)
	fmt.Println(DrawGraphviz(graph))
}
