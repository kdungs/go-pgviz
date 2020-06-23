package main

import (
	"fmt"
	"sort"
	"strings"
)

// DependencyGraph is a directed graph where the vertices are tables
// (or rather table names) and the edges to from a table with a
// foreign key constraint to the table that constraint references.
type DependencyGraph map[string][]string

// BuildDependencies builds the DependencyGraph from a slice of Table.
func BuildDependencies(tables []Table) DependencyGraph {
	graph := make(DependencyGraph)
	for _, table := range tables {
		refs := make([]string, len(table.ForeignKeys))
		for i, fk := range table.ForeignKeys {
			refs[i] = fk.OtherTable
		}
		graph[table.Name] = refs
	}
	return graph
}

// DrawGraphviz produces a string that contains the DependencyGraph's graphviz
// representation.
// This function is guaranteed to produce a representation
// that doesn't depend on the incidental ordering of the graph's edges.
func DrawGraphviz(graph DependencyGraph) string {
	lines := make([]string, len(graph))
	for from, tos := range graph {
		for _, to := range tos {
			line := fmt.Sprintf("  %s -> %s\n", from, to)
			lines = append(lines, line)
		}
	}
	sort.Strings(lines)
	return fmt.Sprintf("digraph DependencyGraph {\n%s}\n", strings.Join(lines, "\n"))
}
