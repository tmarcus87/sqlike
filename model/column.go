package model

import (
	"fmt"
	"strings"
)

type Column interface {
	// カラムを保持しているテーブル情報を返します
	Table() Table

	// カラム名を返します
	ColumnName() string

	// エイリアス名が設定されている場合はエイリアス名、設定されていない場合はカラム名を返します
	AliasOrName() string
}

type ColumnField interface {
	Column

	// フィールド句にエイリアス名を指定します
	As(alias string) ColumnField

	// フィールド句を返します
	FieldExpr() string
}

// ColumnValue column & its value
type ColumnValue interface {
	Column

	// フィールド句を返します
	FieldExpr() string

	// カラム値を返します
	ColumnValue() interface{}
}

func calcExpr(c Column, cExpr, nExpr string) string {
	fn := fmt.Sprintf("`%s`.`%s`", c.Table().SQLikeAliasOrName(), c.ColumnName())
	if cExpr == "" {
		return strings.ReplaceAll(nExpr, "$$", fn)
	}
	return strings.ReplaceAll(nExpr, "$$", "("+cExpr+")")
}

func fieldExpr(c Column, alias, expr string) string {
	fn := fmt.Sprintf("`%s`.`%s`", c.Table().SQLikeAliasOrName(), c.ColumnName())
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

func (a *AllColumn) Table() Table {
	return a.table
}

func (a *AllColumn) ColumnName() string {
	return ""
}

func (a *AllColumn) AliasOrName() string {
	return ""
}

func (a *AllColumn) As(string) ColumnField {
	return a
}

func (a *AllColumn) FieldExpr() string {
	if a.table == nil {
		return "*"
	}
	return fmt.Sprintf("`%s`.*", a.table.SQLikeAliasOrName())
}
