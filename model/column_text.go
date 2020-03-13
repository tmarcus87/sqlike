package model

func NewTextColumn(table Table, name string) *TextColumn {
	return &TextColumn{Table: table, Name: name}
}

type TextField interface {
	ColumnField
}

type TextColumn struct {
	Table Table
	Name  string
	alias string
	expr  string
	value string
}

func (c *TextColumn) SQLikeTable() Table {
	return c.Table
}

func (c *TextColumn) SQLikeColumnName() string {
	return c.Name
}

func (c *TextColumn) SQLikeAliasOrName() string {
	if c.alias != "" {
		return c.alias
	}
	return c.Name
}

func (c *TextColumn) SQLikeAs(alias string) ColumnField {
	c.alias = alias
	return c
}

func (c *TextColumn) SQLikeFieldExpr() string {
	return fieldExpr(c, c.alias, c.expr)
}

func (c *TextColumn) SQLikeSet(v string) ColumnValue {
	c.value = v
	return c
}

func (c *TextColumn) SQLikeColumnValue() interface{} {
	return c.value
}
