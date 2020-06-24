# go-pgviz

Postgres database schema visualizations.

_Warning: This is very much a work in progress. It just does what I need it to
do right now and makes some very strong assumptions about the databases (e.g.
doesn't support passwordless auth and assumes that FOREIGN KEY is used properly
throughout). Nevertheless, feel free to send PRs._

Install via

```zsh
go get github.com/kdungs/go-pgviz
go install github.com/kdungs/go-pgviz
```

The program produces output in `graphviz` format. You'll need to install that
in order to be able to render the output, e.g. via Homebrew on Mac OS using
`brew install graphviz`. Then, you can simply use

```zsh
go-pgviz | dot -Tpdf -o deps.pdf
```

By default, the program just displays the relationships between the tables. If
you want it to also list all columns, use the `-show-columns` option. For more
information on command line parameters, use

```zsh
go-pgviz -help
```

```
Usage of go-pgviz:
  -db string
        Postgres database
  -host string
        Postgres hostname (default "localhost")
  -pass string
        Postgres password
  -port int
        Postgres port (default 5432)
  -show-columns
        whether to show columns for each table
  -show-relations
        whether to show relationships between tables (based on foreign keys) (default true)
  -user string
        Postgres username (default "postgres")
```

