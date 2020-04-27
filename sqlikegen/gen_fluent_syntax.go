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
	imports["fmt"] = struct{}{}
	imports["github.com/tmarcus87/sqlike/model"] = struct{}{}

	g.w.Writeln("import (")
	for impt := range imports {
		g.w.Writeln(`"%s"`, impt)
	}
	g.w.Writeln(")").Ln()

	// Generate accessor
	for _, table := range schema.Schema {
		tableName := strcase.ToCamel(table.Name)
		tableStructName := fmt.Sprintf("%sTable", tableName)

		// Table
		g.w.Writeln("func %s() *%s {", tableName, tableStructName)
		g.w.Writeln("    return &%s{", tableStructName)
		g.w.Writeln("        name: TableName%s,", tableName)
		g.w.Writeln("    }")
		g.w.Writeln("}").Ln()

		// Struct of table
		g.w.Writeln("type %s struct {", tableStructName)
		g.w.Writeln("    name  string")
		g.w.Writeln("    alias string")
		g.w.Writeln("}").Ln()

		g.w.Writeln("func (t *%s) SQLikeTableName() string {", tableStructName)
		g.w.Writeln("    return t.name")
		g.w.Writeln("}").Ln()

		g.w.Writeln("func (t *%s) SQLikeAliasOrName() string {", tableStructName)
		g.w.Writeln("    if t.alias != \"\" {")
		g.w.Writeln("        return t.alias")
		g.w.Writeln("    }")
		g.w.Writeln("    return t.name")
		g.w.Writeln("}").Ln()

		g.w.Writeln("func (t *%s) SQLikeTableExpr() string {", tableStructName)
		g.w.Writeln("    expr := fmt.Sprintf(\"`%%s`\", t.name)")
		g.w.Writeln("    if t.alias != \"\" {")
		g.w.Writeln("        expr = fmt.Sprintf(\"%%s AS `%%s`\", expr, t.alias)")
		g.w.Writeln("    }")
		g.w.Writeln("    return expr")
		g.w.Writeln("}").Ln()

		g.w.Writeln("func (t *%s) As(alias string) *%s {", tableStructName, tableStructName)
		g.w.Writeln("    t.alias = alias")
		g.w.Writeln("    return t")
		g.w.Writeln("}").Ln()

		g.w.Writeln("func (t *%s) SQLikeAllColumns() []model.ColumnField {", tableStructName)
		g.w.Writeln("    return []model.ColumnField {")
		for _, column := range table.Columns {
			g.w.Writeln("        t.%s(),", strcase.ToCamel(column.Name))
		}
		g.w.Writeln("    }")
		g.w.Writeln("}").Ln()

		for _, column := range table.Columns {
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
			g.w.Writeln("}").Ln()
		}
	}

	return g.w.Close()
}
