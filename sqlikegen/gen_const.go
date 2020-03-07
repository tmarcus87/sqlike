package main

import (
	"github.com/iancoleman/strcase"
)

func NewConstGenerator(w Writer) Generator {
	return &ConstGenerator{
		w: w,
	}
}

type ConstGenerator struct {
	w Writer
}

func (g *ConstGenerator) Generate(pkg string, schema *Schema) error {
	g.w.Writeln("package %s", pkg).Ln()

	g.w.Writeln("const (")

	for table, columns := range schema.Schema {
		g.w.Writeln(`TableName%s = "%s"`, strcase.ToCamel(table.Name), table.Name)
		for _, column := range columns {
			g.w.Writeln(`ColumnName%sTable%s = "%s"`, strcase.ToCamel(table.Name), strcase.ToCamel(column.Name), column.Name)
		}
	}

	g.w.Writeln(")")

	return g.w.Close()
}
