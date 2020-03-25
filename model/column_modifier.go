package model

import "fmt"

func Count(field ColumnField) ColumnField {
	return &CountColumnModifier{
		column: field,
	}
}

type CountColumnModifier struct {
	column ColumnField
	alias  string
}

func (c *CountColumnModifier) Table() Table {
	return c.column.Table()
}

func (c *CountColumnModifier) ColumnName() string {
	return c.column.ColumnName()
}

func (c *CountColumnModifier) AliasOrName() string {
	return c.column.AliasOrName()
}

func (c *CountColumnModifier) As(alias string) ColumnField {
	c.alias = alias
	return c
}

func (c *CountColumnModifier) FieldExpr() string {
	expr := fmt.Sprintf("COUNT(%s)", c.column.FieldExpr())
	if c.alias != "" {
		expr += fmt.Sprintf(" AS `%s`", c.alias)
	}
	return expr
}
