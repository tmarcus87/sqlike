package model

import "time"

func NewTimeColumn(table Table, name string) *TimeColumn {
	return &TimeColumn{Table: table, Name: name}
}

type TimeField interface {
	ColumnField
}

type TimeColumn struct {
	Table Table
	Name  string
	alias string
	expr  string
	value time.Time
}

func (c *TimeColumn) SQLikeTable() Table {
	return c.Table
}

func (c *TimeColumn) SQLikeColumnName() string {
	return c.Name
}

func (c *TimeColumn) SQLikeAliasOrName() string {
	if c.alias != "" {
		return c.alias
	}
	return c.Name
}

func (c *TimeColumn) As(alias string) ColumnField {
	c.alias = alias
	return c
}

func (c *TimeColumn) SQLikeFieldExpr() string {
	return fieldExpr(c, c.alias, c.expr)
}

func (c *TimeColumn) SQLikeSet(v time.Time) ColumnValue {
	c.value = v
	return c
}

func (c *TimeColumn) SQLikeColumnValue() interface{} {
	return c.value
}

func (c *TimeColumn) Eq(v time.Time) Condition {
	return &SingleValueCondition{Column: c, Operator: "=", Value: v}
}

func (c *TimeColumn) NotEq(v time.Time) Condition {
	return &SingleValueCondition{Column: c, Operator: "!=", Value: v}
}

func (c *TimeColumn) Gt(v time.Time) Condition {
	return &SingleValueCondition{Column: c, Operator: ">", Value: v}
}

func (c *TimeColumn) GtOrEq(v time.Time) Condition {
	return &SingleValueCondition{Column: c, Operator: ">=", Value: v}
}

func (c *TimeColumn) Lt(v time.Time) Condition {
	return &SingleValueCondition{Column: c, Operator: "<", Value: v}
}

func (c *TimeColumn) LtOrEq(v time.Time) Condition {
	return &SingleValueCondition{Column: c, Operator: "<=", Value: v}
}

func (c *TimeColumn) IsNull() Condition {
	return &NoValueCondition{Column: c, Operator: "IS NULL"}
}

func (c *TimeColumn) IsNotNull() Condition {
	return &NoValueCondition{Column: c, Operator: "IS NOT NULL"}
}

func (c *TimeColumn) EqCol(field ColumnField) Condition {
	return &SingleColumnCondition{Column: c, Operator: "=", Value: field}
}

func (c *TimeColumn) In(vs ...time.Time) Condition {
	return &MultiValueCondition{
		Column:   c,
		Operator: "IN",
		Values:   TimeSliceToInterfaceSlice(vs),
	}
}

func (c *TimeColumn) NotIn(vs ...time.Time) Condition {
	return &MultiValueCondition{
		Column:   c,
		Operator: "NOT IN",
		Values:   TimeSliceToInterfaceSlice(vs),
	}
}

func (c *TimeColumn) Asc() *SortOrder {
	return &SortOrder{
		Column: c,
		Order:  OrderAsc,
	}
}

func (c *TimeColumn) Desc() *SortOrder {
	return &SortOrder{
		Column: c,
		Order:  OrderDesc,
	}
}

func TimeSliceToInterfaceSlice(in []time.Time) []interface{} {
	out := make([]interface{}, 0)
	for _, v := range in {
		out = append(out, v)
	}
	return out
}
