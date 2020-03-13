package model

import "fmt"

func Count(field ColumnField) ColumnField {
	return &CountColumnModifier{
		column: field,
	}
}

type CountColumnModifier struct {
	column ColumnField
	alias  string
}

func (c *CountColumnModifier) SQLikeTable() Table {
	return c.column.SQLikeTable()
}

func (c *CountColumnModifier) SQLikeColumnName() string {
	return c.column.SQLikeColumnName()
}

func (c *CountColumnModifier) SQLikeAliasOrName() string {
	return c.column.SQLikeAliasOrName()
}

func (c *CountColumnModifier) SQLikeAs(alias string) ColumnField {
	c.alias = alias
	return c
}

func (c *CountColumnModifier) SQLikeFieldExpr() string {
	expr := fmt.Sprintf("COUNT(%s)", c.column.SQLikeFieldExpr())
	if c.alias != "" {
		expr += fmt.Sprintf(" AS `%s`", c.alias)
	}
	return expr
}

//
//import "fmt"
//
//// TableAsStatement Aliasが設定されている場合はAs句をつけたテーブル名を返します
//func TableAsStatement(t Table) string {
//	if alias := t.SQLikeTableAlias(); alias != "" {
//		return fmt.Sprintf("`%s` AS `%s`", t.SQLikeTableName(), alias)
//	}
//	return fmt.Sprintf("`%s`", t.SQLikeTableName())
//}
//
//// ColumnAsStatement Aliasが設定されている場合はAs句をつけたテーブル名を返します
//// テーブル名のAliasが設定されている場合はAlias名で表記されます
//func ColumnAsStatement(c Column) string {
//	if alias := c.SQLikeColumnAlias(); alias != "" {
//		return fmt.Sprintf("`%s`.`%s` AS `%s`",
//			TableName(c.SQLikeTable()), c.SQLikeColumnName(), alias)
//	}
//	return fmt.Sprintf("`%s`.`%s`", TableName(c.SQLikeTable()), c.SQLikeColumnName())
//}
//
//func Count(column *BasicColumn) Column {
//	newCol := copyBasicColumn(*column)
//	newCol.selectModFmt = "COUNT(%s)"
//	return newCol
//}
//
//func CountAs(column *BasicColumn, alias string) Column {
//	newCol := copyBasicColumn(*column)
//	newCol.selectModFmt = "COUNT(%s) AS `" + alias + "`"
//	return newCol
//}
//
//func copyBasicColumn(column BasicColumn) *BasicColumn {
//	return &column
//}
