package model

import "fmt"

type BasicTable struct {
	Name  string
	alias string
}

func (t *BasicTable) SQLikeTableName() string {
	return t.Name
}

func (t *BasicTable) SQLikeTableAlias() string {
	return t.alias
}

func (t *BasicTable) SQLikeAs(alias string) Table {
	t.alias = alias
	return t
}

type BasicColumn struct {
	Table        Table
	Name         string
	alias        string
	selectModFmt string
}

func (c *BasicColumn) SQLikeTable() Table {
	return c.Table
}

func (c *BasicColumn) SQLikeColumnName() string {
	return c.Name
}

func (c *BasicColumn) SQLikeColumnAlias() string {
	return c.alias
}

func (c *BasicColumn) SQLikeAs(alias string) Column {
	c.alias = alias
	return c
}

func (c *BasicColumn) SQLikeSelectModFmt() string {
	return c.selectModFmt
}

func (c *BasicColumn) Eq(v interface{}) Condition {
	return &BasicEqCondition{
		column: c,
		v:      v,
	}
}

func (c *BasicColumn) EqCol(column Column) Condition {
	return &BasicEqColCondition{
		left:  c,
		right: column,
	}
}

type BasicEqCondition struct {
	column Column
	v      interface{}
}

func (c *BasicEqCondition) Apply(stmt *string, bindings *[]interface{}) {
	*stmt += fmt.Sprintf("`%s`.`%s` = ?", TableName(c.column.SQLikeTable()), c.column.SQLikeColumnName())
	*bindings = append(*bindings, c.v)
}

type BasicEqColCondition struct {
	left  Column
	right Column
}

func (c *BasicEqColCondition) Apply(stmt *string, bindings *[]interface{}) {
	*stmt +=
		fmt.Sprintf("`%s`.`%s` = `%s`.`%s`",
			TableName(c.left.SQLikeTable()), c.left.SQLikeColumnName(),
			TableName(c.right.SQLikeTable()), c.right.SQLikeColumnName())
}
