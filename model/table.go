package model

import "fmt"

type Table interface {
	// SQLikeTableName テーブル名を返します
	SQLikeTableName() string
	// SQLikeAliasOrName テーブルのエイリアス名もしくはテーブル名を返します
	SQLikeAliasOrName() string
	// SQLikeTableExpr
	SQLikeTableExpr() string
	// SQLikeAs テーブルにエイリアスを設定します
	As(alias string) Table
}

type BasicTable struct {
	Name  string
	alias string
}

func (t *BasicTable) SQLikeTableName() string {
	return t.Name
}

func (t *BasicTable) As(alias string) Table {
	t.alias = alias
	return t
}

func (t *BasicTable) SQLikeAliasOrName() string {
	if t.alias != "" {
		return t.alias
	}
	return t.Name
}

func (t *BasicTable) SQLikeTableExpr() string {
	expr := fmt.Sprintf("`%s`", t.Name)
	if t.alias != "" {
		expr = fmt.Sprintf("%s AS `%s`", expr, t.alias)
	}
	return expr
}

func NewTable(name string) Table {
	return &BasicTable{Name: name}
}
