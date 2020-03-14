package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTimeColumn_SQLikeFieldExpr(t *testing.T) {

	tests := []struct {
		Expect string
		Column ColumnField
	}{
		{
			Expect: "`tbl`.`col`",
			Column: NewTimeColumn(NewTable("tbl"), "col"),
		},
		{
			Expect: "`tbl`.`col` AS `col_alias`",
			Column: NewTimeColumn(NewTable("tbl"), "col").SQLikeAs("col_alias"),
		},
		{
			Expect: "`tbl_alias`.`col`",
			Column: NewTimeColumn(NewTable("tbl").SQLikeAs("tbl_alias"), "col"),
		},
		{
			Expect: "`tbl_alias`.`col` AS `col_alias`",
			Column: NewTimeColumn(NewTable("tbl").SQLikeAs("tbl_alias"), "col").SQLikeAs("col_alias"),
		},
	}

	for _, test := range tests {
		t.Run(test.Expect, func(t *testing.T) {
			asserts := assert.New(t)

			asserts.Equal(test.Expect, test.Column.SQLikeFieldExpr())
		})
	}

}

func TestTimeColumn_SetAndColumnValue(t *testing.T) {

	t1 := time.Now()
	t2 := time.Now().Add(time.Second)

	tests := []struct {
		ExpectExpr string
		Column     *TimeColumn
	}{
		{
			ExpectExpr: "`tbl`.`col`",
			Column:     NewTimeColumn(NewTable("tbl"), "col"),
		},
		{
			ExpectExpr: "`tbl_alias`.`col`",
			Column:     NewTimeColumn(NewTable("tbl").SQLikeAs("tbl_alias"), "col"),
		},
	}

	for _, test := range tests {
		t.Run(test.ExpectExpr, func(t *testing.T) {
			asserts := assert.New(t)

			colV := test.Column.SQLikeSet(t1)
			asserts.Equal(t1, colV.SQLikeColumnValue())

			test.Column.SQLikeSet(t2)
			asserts.Equal(t2, colV.SQLikeColumnValue())
		})
	}

}

func TestTimeColumn_Cond(t *testing.T) {

	t1 := NewTable("t1")
	t2 := NewTable("t2")

	tm1 := time.Now()
	tm2 := time.Now().Add(time.Second)

	tests := []struct {
		Name string
		Cond Condition
		Stmt string
		Bind []interface{}
	}{
		{
			Name: "CondEq",
			Cond: NewTimeColumn(t1, "c1").CondEq(tm1),
			Stmt: "`t1`.`c1` = ?",
			Bind: []interface{}{tm1},
		},
		{
			Name: "CondNotEq",
			Cond: NewTimeColumn(t1, "c1").CondNotEq(tm1),
			Stmt: "`t1`.`c1` != ?",
			Bind: []interface{}{tm1},
		},
		{
			Name: "CondGt",
			Cond: NewTimeColumn(t1, "c1").CondGt(tm1),
			Stmt: "`t1`.`c1` > ?",
			Bind: []interface{}{tm1},
		},
		{
			Name: "CondGtOrEq",
			Cond: NewTimeColumn(t1, "c1").CondGtOrEq(tm1),
			Stmt: "`t1`.`c1` >= ?",
			Bind: []interface{}{tm1},
		},
		{
			Name: "CondLt",
			Cond: NewTimeColumn(t1, "c1").CondLt(tm1),
			Stmt: "`t1`.`c1` < ?",
			Bind: []interface{}{tm1},
		},
		{
			Name: "CondLtOrEq",
			Cond: NewTimeColumn(t1, "c1").CondLtOrEq(tm1),
			Stmt: "`t1`.`c1` <= ?",
			Bind: []interface{}{tm1},
		},
		{
			Name: "CondIsNull",
			Cond: NewTimeColumn(t1, "c1").CondIsNull(),
			Stmt: "`t1`.`c1` IS NULL",
			Bind: []interface{}{},
		},
		{
			Name: "CondIsNotNull",
			Cond: NewTimeColumn(t1, "c1").CondIsNotNull(),
			Stmt: "`t1`.`c1` IS NOT NULL",
			Bind: []interface{}{},
		},
		{
			Name: "CondEqCol",
			Cond: NewTimeColumn(t1, "c1").CondEqCol(NewTextColumn(t2, "c2")),
			Stmt: "`t1`.`c1` = `t2`.`c2`",
			Bind: []interface{}{},
		},
		{
			Name: "CondIn/One",
			Cond: NewTimeColumn(t1, "c1").CondIn(tm1),
			Stmt: "`t1`.`c1` IN (?)",
			Bind: []interface{}{tm1},
		},
		{
			Name: "CondIn/Two",
			Cond: NewTimeColumn(t1, "c1").CondIn(tm1, tm2),
			Stmt: "`t1`.`c1` IN (?, ?)",
			Bind: []interface{}{tm1, tm2},
		},
		{
			Name: "CondNotIn/One",
			Cond: NewTimeColumn(t1, "c1").CondNotIn(tm1),
			Stmt: "`t1`.`c1` NOT IN (?)",
			Bind: []interface{}{tm1},
		},
		{
			Name: "CondNotIn/Two",
			Cond: NewTimeColumn(t1, "c1").CondNotIn(tm1, tm2),
			Stmt: "`t1`.`c1` NOT IN (?, ?)",
			Bind: []interface{}{tm1, tm2},
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
