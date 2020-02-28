package sqlike

import "fmt"

type Table interface {
	// SQLikeTableName テーブル名を返します
	SQLikeTableName() string
	// SQLikeTableAliasName エイリアス名を返します
	SQLikeTableAlias() string
	// SQLikeAs テーブルにエイリアスを設定します
	SQLikeAs(alias string) Table
}

type BasicTable struct {
	name  string
	alias string
}

func (t *BasicTable) SQLikeTableName() string {
	return t.name
}

func (t *BasicTable) SQLikeTableAlias() string {
	return t.alias
}

func (t *BasicTable) SQLikeAs(alias string) Table {
	t.alias = alias
	return t
}

type Column interface {
	// SQLikeTable
	SQLikeTable() Table
	// SQLikeColumnOriginalName
	SQLikeColumnName() string
	// SQLikeColumnAliasName
	SQLikeColumnAlias() string
	// SQLikeAs
	SQLikeAs(alias string) Column
}

type BasicColumn struct {
	table Table
	name  string
	alias string
}

func (c *BasicColumn) SQLikeTable() Table {
	return c.table
}

func (c *BasicColumn) SQLikeColumnName() string {
	return c.name
}

func (c *BasicColumn) SQLikeColumnAlias() string {
	return c.alias
}

func (c *BasicColumn) SQLikeAs(alias string) Column {
	c.alias = alias
	return c
}

func (c *BasicColumn) Eq(v interface{}) Condition {
	return &BasicEqCondition{
		column: c,
		v:      v,
	}
}

type Condition interface {
	Apply(stmt *string, bindings *[]interface{})
}

type BasicEqCondition struct {
	column Column
	v      interface{}
}

func (b *BasicEqCondition) Apply(stmt *string, bindings *[]interface{}) {
	*stmt += fmt.Sprintf("`%s`.`%s` = ?", TableName(b.column.SQLikeTable()), ColumnName(b.column))
	*bindings = append(*bindings, b.v)
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
