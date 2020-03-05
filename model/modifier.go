package model

import "fmt"

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
