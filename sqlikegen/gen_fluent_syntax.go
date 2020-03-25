package main

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"strings"
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

	for table, columns := range schema.Schema {
		tableName := strcase.ToCamel(table.Name)
		tableStructName := fmt.Sprintf("%sTable", tableName)

		g.w.Writeln("func %s() *%s {", tableName, tableStructName)
		g.w.Writeln("    return &%s{", tableStructName)
		g.w.Writeln("        sqlikeName: TableName%s,", tableName)
		g.w.Writeln("    }")
		g.w.Writeln("}").Ln()

		g.w.Writeln("type %s struct {", tableStructName)
		g.w.Writeln("    sqlikeName  string")
		g.w.Writeln("    sqlikeAlias string")
		g.w.Writeln("}")

		g.w.Writeln("func (t *%s) SQLikeTableName() string {", tableStructName)
		g.w.Writeln("    return t.sqlikeName")
		g.w.Writeln("}")

		g.w.Writeln("")

		g.w.Writeln("func (t *%s) SQLikeTableAlias() string {", tableStructName)
		g.w.Writeln("    return t.sqlikeAlias")
		g.w.Writeln("}")

		g.w.Writeln("")

		g.w.Writeln("func (t *%s) SQLikeAs(alias string) model.Table {", tableStructName)
		g.w.Writeln("    t.sqlikeAlias = alias")
		g.w.Writeln("    return t")
		g.w.Writeln("}")

		g.w.Writeln("")

		for _, column := range columns {
			columnName := strcase.ToCamel(column.Name)
			columnStructName := fmt.Sprintf("%s%sColumn", tableName, columnName)

			g.w.Writeln("func (t *%s) %s() *%s {", tableStructName, columnName, columnStructName)
			g.w.Writeln("    return &%s{", columnStructName)
			g.w.Writeln("        sqlikeTable: t,")
			g.w.Writeln("        sqlikeName: ColumnName%sTable%s,", tableName, columnName)
			g.w.Writeln("    }")
			g.w.Writeln("}")

			g.w.Writeln("")

			g.w.Writeln("type %s struct {", columnStructName)
			g.w.Writeln("    sqlikeTable        model.Table")
			g.w.Writeln("    sqlikeName         string")
			g.w.Writeln("    sqlikeAlias        string")
			g.w.Writeln("    sqlikeSelectModFmt string")
			g.w.Writeln("}")

			//g.w.Writeln("func (c *%s) Table() model.Table {", columnStructName)
			//g.w.Writeln("    return c.sqlikeTable")
			//g.w.Writeln("}")

			g.w.Writeln("")

			g.w.Writeln("func (c *%s) ColumnName() string {", columnStructName)
			g.w.Writeln("    return c.sqlikeName")
			g.w.Writeln("}")

			g.w.Writeln("")

			g.w.Writeln("func (c *%s) SQLikeColumnAlias() string {", columnStructName)
			g.w.Writeln("    return c.sqlikeAlias")
			g.w.Writeln("}")

			g.w.Writeln("")

			g.w.Writeln("func (c *%s) SQLikeAs(alias string) model.Column {", columnStructName)
			g.w.Writeln("    c.sqlikeAlias = alias")
			g.w.Writeln("    return c")
			g.w.Writeln("}")

			g.w.Writeln("")

			g.w.Writeln("func (c *%s) SQLikeSelectModFmt() string {", columnStructName)
			g.w.Writeln("    return c.sqlikeSelectModFmt")
			g.w.Writeln("}")

			g.w.Writeln("")

			goType, err := column.GoType()
			if err != nil {
				return err
			}

			g.w.Writeln("func (c *%s) Eq(v %s) model.Condition {", columnStructName, goType)
			g.w.Writeln(`    return &model.SingleValueCondition{`)
			g.w.Writeln(`        Column: c,`)
			g.w.Writeln(`        Operator: "=",`)
			g.w.Writeln(`        Value: v,`)
			g.w.Writeln(`    }`)
			g.w.Writeln("}").Ln()

			g.w.Writeln("func (c *%s) NotEq(v %s) model.Condition {", columnStructName, goType)
			g.w.Writeln(`    return &model.SingleValueCondition{`)
			g.w.Writeln(`        Column: c,`)
			g.w.Writeln(`        Operator: "!=",`)
			g.w.Writeln(`        Value: v,`)
			g.w.Writeln(`    }`)
			g.w.Writeln("}").Ln()

			g.w.Writeln("func (c *%s) IsNull() model.Condition {", columnStructName)
			g.w.Writeln(`    return &model.NoValueCondition{`)
			g.w.Writeln(`        Column: c,`)
			g.w.Writeln(`        Operator: "IS NULL",`)
			g.w.Writeln(`    }`)
			g.w.Writeln("}").Ln()

			g.w.Writeln("func (c *%s) IsNotNull() model.Condition {", columnStructName)
			g.w.Writeln(`    return &model.NoValueCondition{`)
			g.w.Writeln(`        Column: c,`)
			g.w.Writeln(`        Operator: "IS NOT NULL",`)
			g.w.Writeln(`    }`)
			g.w.Writeln("}").Ln()

			g.w.Writeln("func (c *%s) Gt(v %s) model.Condition {", columnStructName, goType)
			g.w.Writeln(`    return &model.SingleValueCondition{`)
			g.w.Writeln(`        Column: c,`)
			g.w.Writeln(`        Operator: ">",`)
			g.w.Writeln(`        Value: v,`)
			g.w.Writeln(`    }`)
			g.w.Writeln("}").Ln()

			g.w.Writeln("func (c *%s) Ge(v %s) model.Condition {", columnStructName, goType)
			g.w.Writeln(`    return &model.SingleValueCondition{`)
			g.w.Writeln(`        Column: c,`)
			g.w.Writeln(`        Operator: ">=",`)
			g.w.Writeln(`        Value: v,`)
			g.w.Writeln(`    }`)
			g.w.Writeln("}").Ln()

			g.w.Writeln("func (c *%s) Lt(v %s) model.Condition {", columnStructName, goType)
			g.w.Writeln(`    return &model.SingleValueCondition{`)
			g.w.Writeln(`        Column: c,`)
			g.w.Writeln(`        Operator: "<",`)
			g.w.Writeln(`        Value: v,`)
			g.w.Writeln(`    }`)
			g.w.Writeln("}").Ln()

			g.w.Writeln("func (c *%s) Le(v %s) model.Condition {", columnStructName, goType)
			g.w.Writeln(`    return &model.SingleValueCondition{`)
			g.w.Writeln(`        Column: c,`)
			g.w.Writeln(`        Operator: "<=",`)
			g.w.Writeln(`        Value: v,`)
			g.w.Writeln(`    }`)
			g.w.Writeln("}").Ln()

			g.w.Writeln("func (c *%s) In(v ...%s) model.Condition {", columnStructName, goType)
			g.w.Writeln(`    return &model.MultiValueCondition{`)
			g.w.Writeln(`        Column: c,`)
			g.w.Writeln(`        Operator: "IN",`)
			g.w.Writeln(`        Values: %sSliceToInterfaceSlice(v),`, strings.ReplaceAll(goType, ".", ""))
			g.w.Writeln(`    }`)
			g.w.Writeln("}").Ln()

			g.w.Writeln("func (c *%s) NotIn(v ...%s) model.Condition {", columnStructName, goType)
			g.w.Writeln(`    return &model.MultiValueCondition{`)
			g.w.Writeln(`        Column: c,`)
			g.w.Writeln(`        Operator: "NOT IN ",`)
			g.w.Writeln(`        Values: %sSliceToInterfaceSlice(v),`, strings.ReplaceAll(goType, ".", ""))
			g.w.Writeln(`    }`)
			g.w.Writeln("}").Ln()

			if goType == typeString {
				g.w.Writeln("func (c *%s) Like(v %s) model.Condition {", columnStructName, goType)
				g.w.Writeln(`    return &model.SingleValueCondition{`)
				g.w.Writeln(`        Column: c,`)
				g.w.Writeln(`        Operator: "LIKE",`)
				g.w.Writeln(`        Value: v,`)
				g.w.Writeln(`    }`)
				g.w.Writeln("}").Ln()

				g.w.Writeln("func (c *%s) NotLike(v %s) model.Condition {", columnStructName, goType)
				g.w.Writeln(`    return &model.SingleValueCondition{`)
				g.w.Writeln(`        Column: c,`)
				g.w.Writeln(`        Operator: "NOT LIKE",`)
				g.w.Writeln(`        Value: v,`)
				g.w.Writeln(`    }`)
				g.w.Writeln("}").Ln()

			}

		}

	}

	goTypes := make(map[string]struct{})
	for _, dtd := range EngineDefs[schema.DBEngine].DataType {
		goTypes[dtd.GoType] = struct{}{}
	}

	for goType := range goTypes {
		g.w.Writeln("func %sSliceToInterfaceSlice(vs []%s) []interface{} {", strings.ReplaceAll(goType, ".", ""), goType)
		g.w.Writeln("    is := make([]interface{}, 0)")
		g.w.Writeln("    for _, v := range vs {")
		g.w.Writeln("        is = append(is, v)")
		g.w.Writeln("    }")
		g.w.Writeln("    return is")
		g.w.Writeln("}").Ln()
	}

	return g.w.Close()
}
