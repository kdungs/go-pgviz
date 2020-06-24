package main

import (
	"fmt"
	"io"
	"text/template"
)

var (
	header = "digraph DependencyGraph {\n"
	footer = "}\n"

	tableTpl = template.Must(template.New("table").Parse(
		`{{ .Name }} [
  shape = "record"
  label = "{ {{ .Name }}{{ if .Columns }} | {{ range .Columns }}{{ .Name }}: {{ if .IsNullable }}nullable {{ end }}{{ .Type }}\l{{ end }}{{ end }}}"
]
`,
	))
)

// RenderGraphviz renders a given list of `Table`s to an `io.Writer`.
// Depending on which information is available on the tables, different outputs
// are produced.
func RenderGraphviz(w io.Writer, tables []Table) error {
	_, err := io.WriteString(w, header)
	if err != nil {
		return err
	}
	if err := renderTables(w, tables); err != nil {
		return err
	}
	if err := renderEdges(w, tables); err != nil {
		return err
	}
	_, err = io.WriteString(w, footer)
	return err
}

func renderTables(w io.Writer, tables []Table) error {
	for _, t := range tables {
		if err := renderTable(w, t); err != nil {
			return err
		}
	}
	return nil
}

func renderTable(w io.Writer, t Table) error {
	return tableTpl.Execute(w, t)
}

func renderEdges(w io.Writer, tables []Table) error {
	for _, t := range tables {
		vfrom := t.Name
		fks := t.ForeignKeys
		for _, fk := range fks {
			vto := fk.OtherTable
			if err := renderEdge(w, vfrom, vto); err != nil {
				return err
			}
		}
	}
	return nil
}

func renderEdge(w io.Writer, vfrom, vto string) error {
	_, err := fmt.Fprintf(w, "%s -> %s;\n", vfrom, vto)
	return err
}
