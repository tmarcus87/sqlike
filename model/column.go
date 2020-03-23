package model

import (
	"fmt"
	"strings"
)

type Column interface {
	// カラムを保持しているテーブル情報を返します
	SQLikeTable() Table

	// SQLikeColumnName カラム名を返します
	SQLikeColumnName() string

	SQLikeAliasOrName() string
}

type ColumnField interface {
	Column

	// SQLikeAs フィールド句にエイリアス名を指定します
	As(alias string) ColumnField

	// SQLikeFieldExpr フィールド句を返します
	SQLikeFieldExpr() string
}

// ColumnValue column & its value
type ColumnValue interface {
	Column

	SQLikeFieldExpr() string

	SQLikeColumnValue() interface{}
}

func calcExpr(c Column, cExpr, nExpr string) string {
	fn := fmt.Sprintf("`%s`.`%s`", c.SQLikeTable().SQLikeAliasOrName(), c.SQLikeColumnName())
	if cExpr == "" {
		return strings.ReplaceAll(nExpr, "$$", fn)
	}
	return strings.ReplaceAll(nExpr, "$$", "("+cExpr+")")
}

func fieldExpr(c Column, alias, expr string) string {
	fn := fmt.Sprintf("`%s`.`%s`", c.SQLikeTable().SQLikeAliasOrName(), c.SQLikeColumnName())
	if expr == "" {
		expr = "$$"
	}
	expr = strings.ReplaceAll(expr, "$$", fn)
	if alias != "" {
		return fmt.Sprintf("%s AS `%s`", expr, alias)
	}
	return expr
}

func NewAllColumnField() ColumnField {
	return &AllColumn{}
}

func NewAllTableColumnField(table Table) ColumnField {
	return &AllColumn{
		table: table,
	}
}

type AllColumn struct {
	table Table
}

func (a *AllColumn) SQLikeTable() Table {
	return a.table
}

func (a *AllColumn) SQLikeColumnName() string {
	return ""
}

func (a *AllColumn) SQLikeAliasOrName() string {
	return ""
}

func (a *AllColumn) As(string) ColumnField {
	return a
}

func (a *AllColumn) SQLikeFieldExpr() string {
	if a.table == nil {
		return "*"
	}
	return fmt.Sprintf("`%s`.*", a.table.SQLikeAliasOrName())
}
