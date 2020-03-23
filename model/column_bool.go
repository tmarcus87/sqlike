package model

func NewBoolColumn(table Table, name string) *BoolColumn {
	return &BoolColumn{Table: table, Name: name}
}

type BoolColumn struct {
	Table Table
	Name  string
	alias string
	value bool
}

func (c *BoolColumn) SQLikeTable() Table {
	return c.Table
}

func (c *BoolColumn) SQLikeColumnName() string {
	return c.Name
}

func (c *BoolColumn) SQLikeAliasOrName() string {
	if c.alias != "" {
		return c.alias
	}
	return c.Name

}

func (c *BoolColumn) As(alias string) ColumnField {
	c.alias = alias
	return c
}

func (c *BoolColumn) SQLikeFieldExpr() string {
	return fieldExpr(c, c.alias, "")
}

func (c *BoolColumn) SQLikeColumnValue() interface{} {
	return c.value
}

func (c *BoolColumn) SetValue(v bool) ColumnValue {
	c.value = v
	return c
}

func (c *BoolColumn) Eq(v bool) Condition {
	return &SingleValueCondition{
		Column:   c,
		Operator: "=",
		Value:    v,
	}
}

func (c *BoolColumn) NotEq(v bool) Condition {
	return &SingleValueCondition{
		Column:   c,
		Operator: "!=",
		Value:    v,
	}
}

func (c *BoolColumn) IsNull() Condition {
	return &NoValueCondition{
		Column:   c,
		Operator: "IS NULL",
	}
}

func (c *BoolColumn) IsNotNull() Condition {
	return &NoValueCondition{
		Column:   c,
		Operator: "IS NOT NULL",
	}
}

func (c *BoolColumn) EqCol(field ColumnField) Condition {
	return &SingleColumnCondition{
		Column:   c,
		Operator: "=",
		Value:    field,
	}
}

func (c *BoolColumn) Asc() *SortOrder {
	return &SortOrder{
		Column: c,
		Order:  OrderAsc,
	}
}

func (c *BoolColumn) Desc() *SortOrder {
	return &SortOrder{
		Column: c,
		Order:  OrderDesc,
	}
}
