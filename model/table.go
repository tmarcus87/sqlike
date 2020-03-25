package model

import "fmt"

type Table interface {
	// テーブル名を返します
	SQLikeTableName() string
	// テーブルのエイリアス名もしくはテーブル名を返します
	SQLikeAliasOrName() string
	// SQLikeTableExpr
	SQLikeTableExpr() string
}

type BasicTable struct {
	Name  string
	alias string
}

func (t *BasicTable) SQLikeTableName() string {
	return t.Name
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

func (t *BasicTable) As(alias string) *BasicTable {
	t.alias = alias
	return t
}

func NewTable(name string) *BasicTable {
	return &BasicTable{Name: name}
}
