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

func (c *CountColumnModifier) SQLikeTable() Table {
	return c.column.SQLikeTable()
}

func (c *CountColumnModifier) SQLikeColumnName() string {
	return c.column.SQLikeColumnName()
}

func (c *CountColumnModifier) SQLikeAliasOrName() string {
	return c.column.SQLikeAliasOrName()
}

func (c *CountColumnModifier) SQLikeAs(alias string) ColumnField {
	c.alias = alias
	return c
}

func (c *CountColumnModifier) SQLikeFieldExpr() string {
	expr := fmt.Sprintf("COUNT(%s)", c.column.SQLikeFieldExpr())
	if c.alias != "" {
		expr += fmt.Sprintf(" AS `%s`", c.alias)
	}
	return expr
}
