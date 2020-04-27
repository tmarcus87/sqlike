package main

import (
	"github.com/iancoleman/strcase"
)

func NewValueEntityGenerator(w Writer) Generator {
	return &ValueEntityGenerator{
		w: w,
	}
}

type ValueEntityGenerator struct {
	w Writer
}

func (g *ValueEntityGenerator) Generate(pkg string, schema *Schema) error {
	// Package
	g.w.Writeln("package model").Ln()

	// Import
	imports := make(map[string]struct{})
	for _, table := range schema.Schema {
		for _, column := range table.Columns {
			it, err := column.ValueImport()
			if err != nil {
				return err
			}
			if it != "" {
				imports[it] = struct{}{}
			}
		}
	}

	if len(imports) > 0 {
		g.w.Writeln("import (")
		for impt := range imports {
			g.w.Writeln(`"%s"`, impt)
		}
		g.w.Writeln(")").Ln()
	}

	// Struct
	for _, table := range schema.Schema {
		g.w.Writeln("type %s struct {", strcase.ToCamel(table.Name))

		for _, column := range table.Columns {
			ft, err := column.ValueFieldType()
			if err != nil {
				return err
			}

			g.w.Writeln("%s %s `sqlike:\"%s\"`",
				strcase.ToCamel(column.Name),
				ft,
				column.Name)
		}

		g.w.Writeln("}")
		g.w.Writeln("")

	}

	return g.w.Close()
}
