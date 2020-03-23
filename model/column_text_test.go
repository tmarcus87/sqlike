package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTextColumn_SQLikeFieldExpr(t *testing.T) {

	tests := []struct {
		Expect string
		Column ColumnField
	}{
		{
			Expect: "`tbl`.`col`",
			Column: NewTextColumn(NewTable("tbl"), "col"),
		},
		{
			Expect: "`tbl`.`col` AS `col_alias`",
			Column: NewTextColumn(NewTable("tbl"), "col").As("col_alias"),
		},
		{
			Expect: "`tbl_alias`.`col`",
			Column: NewTextColumn(NewTable("tbl").As("tbl_alias"), "col"),
		},
		{
			Expect: "`tbl_alias`.`col` AS `col_alias`",
			Column: NewTextColumn(NewTable("tbl").As("tbl_alias"), "col").As("col_alias"),
		},
	}

	for _, test := range tests {
		t.Run(test.Expect, func(t *testing.T) {
			asserts := assert.New(t)

			asserts.Equal(test.Expect, test.Column.SQLikeFieldExpr())
		})
	}

}

func TestTextColumn_SetAndColumnValue(t *testing.T) {

	tests := []struct {
		ExpectExpr string
		Column     *TextColumn
	}{
		{
			ExpectExpr: "`tbl`.`col`",
			Column:     NewTextColumn(NewTable("tbl"), "col"),
		},
		{
			ExpectExpr: "`tbl_alias`.`col`",
			Column:     NewTextColumn(NewTable("tbl").As("tbl_alias"), "col"),
		},
	}

	for _, test := range tests {
		t.Run(test.ExpectExpr, func(t *testing.T) {
			asserts := assert.New(t)

			colV := test.Column.SQLikeSet("hogehoge")
			asserts.Equal("hogehoge", colV.SQLikeColumnValue())

			test.Column.SQLikeSet("fugafuga")
			asserts.Equal("fugafuga", colV.SQLikeColumnValue())
		})
	}

}

func TestTextColumn_Cond(t *testing.T) {

	t1 := NewTable("t1")
	t2 := NewTable("t2")

	tests := []struct {
		Name string
		Cond Condition
		Stmt string
		Bind []interface{}
	}{
		{
			Name: "CondEq",
			Cond: NewTextColumn(t1, "c1").CondEq("hoge"),
			Stmt: "`t1`.`c1` = ?",
			Bind: []interface{}{"hoge"},
		},
		{
			Name: "CondNotEq",
			Cond: NewTextColumn(t1, "c1").CondNotEq("hoge"),
			Stmt: "`t1`.`c1` != ?",
			Bind: []interface{}{"hoge"},
		},
		{
			Name: "CondGt",
			Cond: NewTextColumn(t1, "c1").CondGt("hoge"),
			Stmt: "`t1`.`c1` > ?",
			Bind: []interface{}{"hoge"},
		},
		{
			Name: "CondGtOrEq",
			Cond: NewTextColumn(t1, "c1").CondGtOrEq("hoge"),
			Stmt: "`t1`.`c1` >= ?",
			Bind: []interface{}{"hoge"},
		},
		{
			Name: "CondLt",
			Cond: NewTextColumn(t1, "c1").CondLt("hoge"),
			Stmt: "`t1`.`c1` < ?",
			Bind: []interface{}{"hoge"},
		},
		{
			Name: "CondLtOrEq",
			Cond: NewTextColumn(t1, "c1").CondLtOrEq("hoge"),
			Stmt: "`t1`.`c1` <= ?",
			Bind: []interface{}{"hoge"},
		},
		{
			Name: "CondLike",
			Cond: NewTextColumn(t1, "c1").CondLike("%hoge%"),
			Stmt: "`t1`.`c1` LIKE ?",
			Bind: []interface{}{"%hoge%"},
		},
		{
			Name: "CondIsNull",
			Cond: NewTextColumn(t1, "c1").CondIsNull(),
			Stmt: "`t1`.`c1` IS NULL",
			Bind: []interface{}{},
		},
		{
			Name: "CondIsNotNull",
			Cond: NewTextColumn(t1, "c1").CondIsNotNull(),
			Stmt: "`t1`.`c1` IS NOT NULL",
			Bind: []interface{}{},
		},
		{
			Name: "CondEqCol",
			Cond: NewTextColumn(t1, "c1").CondEqCol(NewTextColumn(t2, "c2")),
			Stmt: "`t1`.`c1` = `t2`.`c2`",
			Bind: []interface{}{},
		},
		{
			Name: "CondIn/One",
			Cond: NewTextColumn(t1, "c1").CondIn("hoge"),
			Stmt: "`t1`.`c1` IN (?)",
			Bind: []interface{}{"hoge"},
		},
		{
			Name: "CondIn/Two",
			Cond: NewTextColumn(t1, "c1").CondIn("hoge", "fuga"),
			Stmt: "`t1`.`c1` IN (?, ?)",
			Bind: []interface{}{"hoge", "fuga"},
		},
		{
			Name: "CondNotIn/One",
			Cond: NewTextColumn(t1, "c1").CondNotIn("hoge"),
			Stmt: "`t1`.`c1` NOT IN (?)",
			Bind: []interface{}{"hoge"},
		},
		{
			Name: "CondNotIn/Two",
			Cond: NewTextColumn(t1, "c1").CondNotIn("hoge", "fuga"),
			Stmt: "`t1`.`c1` NOT IN (?, ?)",
			Bind: []interface{}{"hoge", "fuga"},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			stmt := ""
			bindings := make([]interface{}, 0)
			test.Cond.Apply(&stmt, &bindings)

			asserts := assert.New(t)

			asserts.Equal(test.Stmt, stmt)
			asserts.Len(bindings, len(test.Bind))
			asserts.EqualValues(test.Bind, bindings)
		})
	}

}
