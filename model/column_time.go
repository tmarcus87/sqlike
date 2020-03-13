package model

import "time"

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

func (c *TimeColumn) SQLikeAs(alias string) ColumnField {
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
