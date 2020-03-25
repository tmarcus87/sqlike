package model

import "database/sql"

func NewTextColumn(table Table, name string) *TextColumn {
	return &TextColumn{table: table, name: name}
}

type TextField interface {
	ColumnField
}

type TextColumn struct {
	table Table
	name  string
	alias string
	expr  string
	value sql.NullString
}

func (c *TextColumn) Table() Table {
	return c.table
}

func (c *TextColumn) ColumnName() string {
	return c.name
}

func (c *TextColumn) AliasOrName() string {
	if c.alias != "" {
		return c.alias
	}
	return c.name
}

func (c *TextColumn) As(alias string) ColumnField {
	c.alias = alias
	return c
}

func (c *TextColumn) FieldExpr() string {
	return fieldExpr(c, c.alias, c.expr)
}

func (c *TextColumn) NullValue() ColumnValue {
	return c
}

func (c *TextColumn) Value(v string) ColumnValue {
	c.value = sql.NullString{String: v, Valid: true}
	return c
}

func (c *TextColumn) ColumnValue() interface{} {
	if c.value.Valid {
		return c.value.String
	}
	return c.value
}

func (c *TextColumn) Eq(v string) Condition {
	return &SingleValueCondition{Column: c, Operator: "=", Value: v}
}

func (c *TextColumn) NotEq(v string) Condition {
	return &SingleValueCondition{Column: c, Operator: "!=", Value: v}
}

func (c *TextColumn) Gt(v string) Condition {
	return &SingleValueCondition{Column: c, Operator: ">", Value: v}
}

func (c *TextColumn) GtOrEq(v string) Condition {
	return &SingleValueCondition{Column: c, Operator: ">=", Value: v}
}

func (c *TextColumn) Lt(v string) Condition {
	return &SingleValueCondition{Column: c, Operator: "<", Value: v}
}

func (c *TextColumn) LtOrEq(v string) Condition {
	return &SingleValueCondition{Column: c, Operator: "<=", Value: v}
}

func (c *TextColumn) Like(v string) Condition {
	return &SingleValueCondition{Column: c, Operator: "LIKE", Value: v}
}

func (c *TextColumn) IsNull() Condition {
	return &NoValueCondition{Column: c, Operator: "IS NULL"}
}

func (c *TextColumn) IsNotNull() Condition {
	return &NoValueCondition{Column: c, Operator: "IS NOT NULL"}
}

func (c *TextColumn) EqCol(field ColumnField) Condition {
	return &SingleColumnCondition{Column: c, Operator: "=", Value: field}
}

func (c *TextColumn) In(vs ...string) Condition {
	return &MultiValueCondition{
		Column:   c,
		Operator: "IN",
		Values:   StringSliceToInterfaceSlice(vs),
	}
}

func (c *TextColumn) NotIn(vs ...string) Condition {
	return &MultiValueCondition{
		Column:   c,
		Operator: "NOT IN",
		Values:   StringSliceToInterfaceSlice(vs),
	}
}

func (c *TextColumn) Asc() *SortOrder {
	return &SortOrder{
		Column: c,
		Order:  OrderAsc,
	}
}

func (c *TextColumn) Desc() *SortOrder {
	return &SortOrder{
		Column: c,
		Order:  OrderDesc,
	}
}

func StringSliceToInterfaceSlice(in []string) []interface{} {
	out := make([]interface{}, 0)
	for _, v := range in {
		out = append(out, v)
	}
	return out
}
