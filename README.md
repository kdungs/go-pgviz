# go-pgviz

Postgres database schema visualizations.

Install via

```zsh
go get github.com/kdungs/go-pgviz
go install github.com/kdungs/go-pgviz
```

The program produces output in `graphviz` format. You'll need to install it in
order to be able to render the output, e.g. via Homebrew on Mac OS using `brew
install graphviz`. Then, you can simply use

```zsh
go-pgviz | dot -Tpdf -o deps.pdf
```

For information on command line parameters, use

```zsh
go-pgviz -help
```
