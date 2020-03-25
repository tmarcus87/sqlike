package model

import "database/sql"

func NewBoolColumn(table Table, name string) *BoolColumn {
	return &BoolColumn{table: table, name: name}
}

type BoolColumn struct {
	table Table
	name  string
	alias string
	value sql.NullBool
}

func (c *BoolColumn) Table() Table {
	return c.table
}

func (c *BoolColumn) ColumnName() string {
	return c.name
}

func (c *BoolColumn) AliasOrName() string {
	if c.alias != "" {
		return c.alias
	}
	return c.name
}

func (c *BoolColumn) As(alias string) ColumnField {
	c.alias = alias
	return c
}

func (c *BoolColumn) FieldExpr() string {
	return fieldExpr(c, c.alias, "")
}

func (c *BoolColumn) ColumnValue() interface{} {
	if c.value.Valid {
		return c.value.Bool
	}
	return c.value
}

func (c *BoolColumn) NullValue() ColumnValue {
	return c
}

func (c *BoolColumn) Value(v bool) ColumnValue {
	c.value = sql.NullBool{Bool: v, Valid: true}
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
