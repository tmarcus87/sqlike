package main

import (
	"fmt"
	"github.com/iancoleman/strcase"
)

func NewSchemaSourceGenerator(w Writer) Generator {
	return &FluentSyntaxSourceGenerator{
		w: w,
	}
}

type FluentSyntaxSourceGenerator struct {
	w Writer
}

func (g *FluentSyntaxSourceGenerator) Generate(pkg string, schema *Schema) error {
	g.w.Writeln("package %s", pkg).Ln()

	// Import
	imports := make(map[string]struct{})
	imports["github.com/tmarcus87/sqlike/model"] = struct{}{}
	for _, columns := range schema.Schema {
		for _, column := range columns {
			if impt, err := column.Import(); err != nil {
				return err
			} else if impt != "" {
				imports[impt] = struct{}{}
			}
		}
	}
	g.w.Writeln("import (")
	for impt := range imports {
		g.w.Writeln(`"%s"`, impt)
	}
	g.w.Writeln(")").Ln()

	// Generate accessor
	for table, columns := range schema.Schema {
		tableName := strcase.ToCamel(table.Name)
		tableStructName := fmt.Sprintf("%sTable", tableName)

		// Table
		g.w.Writeln("func %s() *%s {", tableName, tableStructName)
		g.w.Writeln("    return &%s{", tableStructName)
		g.w.Writeln("        Table: model.NewTable(\"%s\"),", tableName)
		g.w.Writeln("    }")
		g.w.Writeln("}").Ln()

		// Struct of table
		g.w.Writeln("type %s struct {", tableStructName)
		g.w.Writeln("    model.Table")
		g.w.Writeln("}").Ln()

		for _, column := range columns {
			columnName := strcase.ToCamel(column.Name)
			columnStructName := fmt.Sprintf("%s%sColumn", tableName, columnName)

			baseStructType, err := column.StructType()
			if err != nil {
				return err
			}

			// Column func in table
			g.w.Writeln("func (t *%s) %s() *%s {", tableStructName, columnName, columnStructName)
			g.w.Writeln("    return &%s{", columnStructName)
			g.w.Writeln("        %s: model.New%s(t, ColumnName%sTable%s),", baseStructType, baseStructType, tableName, columnName)
			g.w.Writeln("    }")
			g.w.Writeln("}").Ln()

			// Column struct
			g.w.Writeln("type %s struct {", columnStructName)
			g.w.Writeln("    *model.%s", baseStructType)
			g.w.Writeln("}")
		}
	}

	return g.w.Close()
}
