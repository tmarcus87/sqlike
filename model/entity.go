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

// TableName テーブル名を返します
// Aliasが設定されている場合はエイリアス名を返します
func TableName(t Table) string {
	tname := t.SQLikeTableName()
	if alias := t.SQLikeTableAlias(); alias != "" {
		tname = alias
	}
	return tname
}

// ColumnName カラム名を返します
// Aliasが設定されている場合はエイリアス名を返します
func ColumnName(c Column) string {
	cname := c.SQLikeColumnName()
	if alias := c.SQLikeColumnAlias(); alias != "" {
		cname = alias
	}
	return cname
}

// TableAsStatement Aliasが設定されている場合はAs句をつけたテーブル名を返します
func TableAsStatement(t Table) string {
	if alias := t.SQLikeTableAlias(); alias != "" {
		return fmt.Sprintf("`%s` AS `%s`", t.SQLikeTableName(), alias)
	}
	return fmt.Sprintf("`%s`", t.SQLikeTableName())
}

// ColumnAsStatement Aliasが設定されている場合はAs句をつけたテーブル名を返します
// テーブル名のAliasが設定されている場合はAlias名で表記されます
func ColumnAsStatement(c Column) string {
	if alias := c.SQLikeColumnAlias(); alias != "" {
		return fmt.Sprintf("`%s`.`%s` AS `%s`",
			TableName(c.SQLikeTable()), c.SQLikeColumnName(), alias)
	}
	return fmt.Sprintf("`%s`.`%s`", TableName(c.SQLikeTable()), c.SQLikeColumnName())
}

func Count(column *BasicColumn) Column {
	newCol := copyBasicColumn(*column)
	newCol.selectModFmt = "COUNT(%s)"
	return newCol
}

func CountAs(column *BasicColumn, alias string) Column {
	newCol := copyBasicColumn(*column)
	newCol.selectModFmt = "COUNT(%s) AS `" + alias + "`"
	return newCol
}

func copyBasicColumn(column BasicColumn) *BasicColumn {
	return &column
}
