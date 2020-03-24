package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumericColumn_SQLikeFieldExpr(t *testing.T) {
	tbl := NewTable("tbl")
	tblAlias := NewTable("tbl").As("tbl_alias")

	tests := []struct {
		Expect string
		Column ColumnField
	}{
		{
			Expect: "`tbl`.`col`",
			Column: &NumberColumn{Table: tbl, Name: "col"},
		},
		{
			Expect: "`tbl`.`col` AS `col_alias`",
			Column: (&NumberColumn{Table: tbl, Name: "col"}).As("col_alias"),
		},
		{
			Expect: "`tbl_alias`.`col`",
			Column: &NumberColumn{Table: tblAlias, Name: "col"},
		},
		{
			Expect: "`tbl_alias`.`col` AS `col_alias`",
			Column: (&NumberColumn{Table: tblAlias, Name: "col"}).As("col_alias"),
		},
		{
			Expect: "`tbl`.`col` + 123",
			Column: (&NumberColumn{Table: tbl, Name: "col"}).PlusInt(123),
		},
		{
			Expect: "`tbl`.`col` + 123.4",
			Column: (&NumberColumn{Table: tbl, Name: "col"}).PlusFloat(123.4),
		},
		{
			Expect: "`tbl`.`col` - 123",
			Column: (&NumberColumn{Table: tbl, Name: "col"}).MinusInt(123),
		},
		{
			Expect: "`tbl`.`col` - 123.4",
			Column: (&NumberColumn{Table: tbl, Name: "col"}).MinusFloat(123.4),
		},
		{
			Expect: "`tbl`.`col` * 123",
			Column: (&NumberColumn{Table: tbl, Name: "col"}).MultipleInt(123),
		},
		{
			Expect: "`tbl`.`col` * 123.4",
			Column: (&NumberColumn{Table: tbl, Name: "col"}).MultipleFloat(123.4),
		},
		{
			Expect: "`tbl`.`col` / 123",
			Column: (&NumberColumn{Table: tbl, Name: "col"}).DivideInt(123),
		},
		{
			Expect: "`tbl`.`col` / 123.4",
			Column: (&NumberColumn{Table: tbl, Name: "col"}).DivideFloat(123.4),
		},
		{
			Expect: "(`tbl`.`col` + 123) * 456",
			Column: (&NumberColumn{Table: tbl, Name: "col"}).PlusInt(123).MultipleInt(456),
		},
		{
			Expect: "`tbl`.`col` + 123 AS `col_plus`",
			Column: (&NumberColumn{Table: tbl, Name: "col"}).PlusInt(123).As("col_plus"),
		},
	}

	for _, test := range tests {
		t.Run(test.Expect, func(t *testing.T) {
			asserts := assert.New(t)

			asserts.Equal(test.Expect, test.Column.SQLikeFieldExpr())
		})
	}
}

func TestInt8Column_SetAndColumnValue(t *testing.T) {

	tests := []struct {
		ExpectExpr string
		Column     *Int8Column
	}{
		{
			ExpectExpr: "`tbl`.`col`",
			Column:     NewInt8Column(NewTable("tbl"), "col"),
		},
		{
			ExpectExpr: "`tbl_alias`.`col`",
			Column:     NewInt8Column(NewTable("tbl").As("tbl_alias"), "col"),
		},
	}

	for _, test := range tests {
		t.Run(test.ExpectExpr, func(t *testing.T) {
			asserts := assert.New(t)

			colV := test.Column.SQLikeSet(1)
			asserts.Equal(int8(1), colV.SQLikeColumnValue())

			test.Column.SQLikeSet(2)
			asserts.Equal(int8(2), colV.SQLikeColumnValue())
		})
	}

}

func TestInt16Column_SetAndColumnValue(t *testing.T) {

	tests := []struct {
		ExpectExpr string
		Column     *Int16Column
	}{
		{
			ExpectExpr: "`tbl`.`col`",
			Column:     NewInt16Column(NewTable("tbl"), "col"),
		},
		{
			ExpectExpr: "`tbl_alias`.`col`",
			Column:     NewInt16Column(NewTable("tbl").As("tbl_alias"), "col"),
		},
	}

	for _, test := range tests {
		t.Run(test.ExpectExpr, func(t *testing.T) {
			asserts := assert.New(t)

			colV := test.Column.SQLikeSet(1)
			asserts.Equal(int16(1), colV.SQLikeColumnValue())

			test.Column.SQLikeSet(2)
			asserts.Equal(int16(2), colV.SQLikeColumnValue())
		})
	}

}

func TestInt32Column_SetAndColumnValue(t *testing.T) {

	tests := []struct {
		ExpectExpr string
		Column     *Int32Column
	}{
		{
			ExpectExpr: "`tbl`.`col`",
			Column:     NewInt32Column(NewTable("tbl"), "col"),
		},
		{
			ExpectExpr: "`tbl_alias`.`col`",
			Column:     NewInt32Column(NewTable("tbl").As("tbl_alias"), "col"),
		},
	}

	for _, test := range tests {
		t.Run(test.ExpectExpr, func(t *testing.T) {
			asserts := assert.New(t)

			colV := test.Column.SQLikeSet(1)
			asserts.Equal(int32(1), colV.SQLikeColumnValue())

			test.Column.SQLikeSet(2)
			asserts.Equal(int32(2), colV.SQLikeColumnValue())
		})
	}

}

func TestInt64Column_SetAndColumnValue(t *testing.T) {

	tests := []struct {
		ExpectExpr string
		Column     *Int64Column
	}{
		{
			ExpectExpr: "`tbl`.`col`",
			Column:     NewInt64Column(NewTable("tbl"), "col"),
		},
		{
			ExpectExpr: "`tbl_alias`.`col`",
			Column:     NewInt64Column(NewTable("tbl").As("tbl_alias"), "col"),
		},
	}

	for _, test := range tests {
		t.Run(test.ExpectExpr, func(t *testing.T) {
			asserts := assert.New(t)

			colV := test.Column.SQLikeSet(1)
			asserts.Equal(int64(1), colV.SQLikeColumnValue())

			test.Column.SQLikeSet(2)
			asserts.Equal(int64(2), colV.SQLikeColumnValue())
		})
	}

}

func TestFloat32Column_SetAndColumnValue(t *testing.T) {

	tests := []struct {
		ExpectExpr string
		Column     *Float32Column
	}{
		{
			ExpectExpr: "`tbl`.`col`",
			Column:     NewFloat32Column(NewTable("tbl"), "col"),
		},
		{
			ExpectExpr: "`tbl_alias`.`col`",
			Column:     NewFloat32Column(NewTable("tbl").As("tbl_alias"), "col"),
		},
	}

	for _, test := range tests {
		t.Run(test.ExpectExpr, func(t *testing.T) {
			asserts := assert.New(t)

			colV := test.Column.SQLikeSet(1.1)
			asserts.Equal(float32(1.1), colV.SQLikeColumnValue())

			test.Column.SQLikeSet(2.2)
			asserts.Equal(float32(2.2), colV.SQLikeColumnValue())
		})
	}

}

func TestFloat64Column_SetAndColumnValue(t *testing.T) {

	tests := []struct {
		ExpectExpr string
		Column     *Float64Column
	}{
		{
			ExpectExpr: "`tbl`.`col`",
			Column:     NewFloat64Column(NewTable("tbl"), "col"),
		},
		{
			ExpectExpr: "`tbl_alias`.`col`",
			Column:     NewFloat64Column(NewTable("tbl").As("tbl_alias"), "col"),
		},
	}

	for _, test := range tests {
		t.Run(test.ExpectExpr, func(t *testing.T) {
			asserts := assert.New(t)

			colV := test.Column.SQLikeSet(1.1)
			asserts.Equal(1.1, colV.SQLikeColumnValue())

			test.Column.SQLikeSet(2.2)
			asserts.Equal(2.2, colV.SQLikeColumnValue())
		})
	}

}

func TestInt8Column_Cond(t *testing.T) {
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
			Cond: NewInt8Column(t1, "c1").Eq(123),
			Stmt: "`t1`.`c1` = ?",
			Bind: []interface{}{int8(123)},
		},
		{
			Name: "CondNotEq",
			Cond: NewInt8Column(t1, "c1").NotEq(123),
			Stmt: "`t1`.`c1` != ?",
			Bind: []interface{}{int8(123)},
		},
		{
			Name: "CondGt",
			Cond: NewInt8Column(t1, "c1").Gt(123),
			Stmt: "`t1`.`c1` > ?",
			Bind: []interface{}{int8(123)},
		},
		{
			Name: "CondGtOrEq",
			Cond: NewInt8Column(t1, "c1").GtOrEq(123),
			Stmt: "`t1`.`c1` >= ?",
			Bind: []interface{}{int8(123)},
		},
		{
			Name: "CondLt",
			Cond: NewInt8Column(t1, "c1").Lt(123),
			Stmt: "`t1`.`c1` < ?",
			Bind: []interface{}{int8(123)},
		},
		{
			Name: "CondLtOrEq",
			Cond: NewInt8Column(t1, "c1").LtOrEq(123),
			Stmt: "`t1`.`c1` <= ?",
			Bind: []interface{}{int8(123)},
		},
		{
			Name: "CondIsNull",
			Cond: NewInt8Column(t1, "c1").IsNull(),
			Stmt: "`t1`.`c1` IS NULL",
			Bind: []interface{}{},
		},
		{
			Name: "CondIsNotNull",
			Cond: NewInt8Column(t1, "c1").IsNotNull(),
			Stmt: "`t1`.`c1` IS NOT NULL",
			Bind: []interface{}{},
		},
		{
			Name: "CondEqCol",
			Cond: NewInt8Column(t1, "c1").EqCol(NewInt8Column(t2, "c2")),
			Stmt: "`t1`.`c1` = `t2`.`c2`",
			Bind: []interface{}{},
		},
		{
			Name: "CondIn/One",
			Cond: NewInt8Column(t1, "c1").In(1),
			Stmt: "`t1`.`c1` IN (?)",
			Bind: []interface{}{int8(1)},
		},
		{
			Name: "CondIn/Two",
			Cond: NewInt8Column(t1, "c1").In(1, 2),
			Stmt: "`t1`.`c1` IN (?, ?)",
			Bind: []interface{}{int8(1), int8(2)},
		},
		{
			Name: "CondNotIn/One",
			Cond: NewInt8Column(t1, "c1").NotIn(1),
			Stmt: "`t1`.`c1` NOT IN (?)",
			Bind: []interface{}{int8(1)},
		},
		{
			Name: "CondNotIn/Two",
			Cond: NewInt8Column(t1, "c1").NotIn(1, 2),
			Stmt: "`t1`.`c1` NOT IN (?, ?)",
			Bind: []interface{}{int8(1), int8(2)},
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

func TestInt16Column_Cond(t *testing.T) {
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
			Cond: NewInt16Column(t1, "c1").Eq(123),
			Stmt: "`t1`.`c1` = ?",
			Bind: []interface{}{int16(123)},
		},
		{
			Name: "CondNotEq",
			Cond: NewInt16Column(t1, "c1").NotEq(123),
			Stmt: "`t1`.`c1` != ?",
			Bind: []interface{}{int16(123)},
		},
		{
			Name: "CondGt",
			Cond: NewInt16Column(t1, "c1").Gt(123),
			Stmt: "`t1`.`c1` > ?",
			Bind: []interface{}{int16(123)},
		},
		{
			Name: "CondGtOrEq",
			Cond: NewInt16Column(t1, "c1").GtOrEq(123),
			Stmt: "`t1`.`c1` >= ?",
			Bind: []interface{}{int16(123)},
		},
		{
			Name: "CondLt",
			Cond: NewInt16Column(t1, "c1").Lt(123),
			Stmt: "`t1`.`c1` < ?",
			Bind: []interface{}{int16(123)},
		},
		{
			Name: "CondLtOrEq",
			Cond: NewInt16Column(t1, "c1").LtOrEq(123),
			Stmt: "`t1`.`c1` <= ?",
			Bind: []interface{}{int16(123)},
		},
		{
			Name: "CondIsNull",
			Cond: NewInt16Column(t1, "c1").IsNull(),
			Stmt: "`t1`.`c1` IS NULL",
			Bind: []interface{}{},
		},
		{
			Name: "CondIsNotNull",
			Cond: NewInt16Column(t1, "c1").IsNotNull(),
			Stmt: "`t1`.`c1` IS NOT NULL",
			Bind: []interface{}{},
		},
		{
			Name: "CondEqCol",
			Cond: NewInt16Column(t1, "c1").EqCol(NewInt8Column(t2, "c2")),
			Stmt: "`t1`.`c1` = `t2`.`c2`",
			Bind: []interface{}{},
		},
		{
			Name: "CondIn/One",
			Cond: NewInt16Column(t1, "c1").In(1),
			Stmt: "`t1`.`c1` IN (?)",
			Bind: []interface{}{int16(1)},
		},
		{
			Name: "CondIn/Two",
			Cond: NewInt16Column(t1, "c1").In(1, 2),
			Stmt: "`t1`.`c1` IN (?, ?)",
			Bind: []interface{}{int16(1), int16(2)},
		},
		{
			Name: "CondNotIn/One",
			Cond: NewInt16Column(t1, "c1").NotIn(1),
			Stmt: "`t1`.`c1` NOT IN (?)",
			Bind: []interface{}{int16(1)},
		},
		{
			Name: "CondNotIn/Two",
			Cond: NewInt16Column(t1, "c1").NotIn(1, 2),
			Stmt: "`t1`.`c1` NOT IN (?, ?)",
			Bind: []interface{}{int16(1), int16(2)},
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

func TestInt32Column_Cond(t *testing.T) {
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
			Cond: NewInt32Column(t1, "c1").Eq(123),
			Stmt: "`t1`.`c1` = ?",
			Bind: []interface{}{int32(123)},
		},
		{
			Name: "CondNotEq",
			Cond: NewInt32Column(t1, "c1").NotEq(123),
			Stmt: "`t1`.`c1` != ?",
			Bind: []interface{}{int32(123)},
		},
		{
			Name: "CondGt",
			Cond: NewInt32Column(t1, "c1").Gt(123),
			Stmt: "`t1`.`c1` > ?",
			Bind: []interface{}{int32(123)},
		},
		{
			Name: "CondGtOrEq",
			Cond: NewInt32Column(t1, "c1").GtOrEq(123),
			Stmt: "`t1`.`c1` >= ?",
			Bind: []interface{}{int32(123)},
		},
		{
			Name: "CondLt",
			Cond: NewInt32Column(t1, "c1").Lt(123),
			Stmt: "`t1`.`c1` < ?",
			Bind: []interface{}{int32(123)},
		},
		{
			Name: "CondLtOrEq",
			Cond: NewInt32Column(t1, "c1").LtOrEq(123),
			Stmt: "`t1`.`c1` <= ?",
			Bind: []interface{}{int32(123)},
		},
		{
			Name: "CondIsNull",
			Cond: NewInt32Column(t1, "c1").IsNull(),
			Stmt: "`t1`.`c1` IS NULL",
			Bind: []interface{}{},
		},
		{
			Name: "CondIsNotNull",
			Cond: NewInt32Column(t1, "c1").IsNotNull(),
			Stmt: "`t1`.`c1` IS NOT NULL",
			Bind: []interface{}{},
		},
		{
			Name: "CondEqCol",
			Cond: NewInt32Column(t1, "c1").EqCol(NewInt8Column(t2, "c2")),
			Stmt: "`t1`.`c1` = `t2`.`c2`",
			Bind: []interface{}{},
		},
		{
			Name: "CondIn/One",
			Cond: NewInt32Column(t1, "c1").In(1),
			Stmt: "`t1`.`c1` IN (?)",
			Bind: []interface{}{int32(1)},
		},
		{
			Name: "CondIn/Two",
			Cond: NewInt32Column(t1, "c1").In(1, 2),
			Stmt: "`t1`.`c1` IN (?, ?)",
			Bind: []interface{}{int32(1), int32(2)},
		},
		{
			Name: "CondNotIn/One",
			Cond: NewInt32Column(t1, "c1").NotIn(1),
			Stmt: "`t1`.`c1` NOT IN (?)",
			Bind: []interface{}{int32(1)},
		},
		{
			Name: "CondNotIn/Two",
			Cond: NewInt32Column(t1, "c1").NotIn(1, 2),
			Stmt: "`t1`.`c1` NOT IN (?, ?)",
			Bind: []interface{}{int32(1), int32(2)},
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

func TestInt64Column_Cond(t *testing.T) {
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
			Cond: NewInt64Column(t1, "c1").Eq(123),
			Stmt: "`t1`.`c1` = ?",
			Bind: []interface{}{int64(123)},
		},
		{
			Name: "CondNotEq",
			Cond: NewInt64Column(t1, "c1").NotEq(123),
			Stmt: "`t1`.`c1` != ?",
			Bind: []interface{}{int64(123)},
		},
		{
			Name: "CondGt",
			Cond: NewInt64Column(t1, "c1").Gt(123),
			Stmt: "`t1`.`c1` > ?",
			Bind: []interface{}{int64(123)},
		},
		{
			Name: "CondGtOrEq",
			Cond: NewInt64Column(t1, "c1").GtOrEq(123),
			Stmt: "`t1`.`c1` >= ?",
			Bind: []interface{}{int64(123)},
		},
		{
			Name: "CondLt",
			Cond: NewInt64Column(t1, "c1").Lt(123),
			Stmt: "`t1`.`c1` < ?",
			Bind: []interface{}{int64(123)},
		},
		{
			Name: "CondLtOrEq",
			Cond: NewInt64Column(t1, "c1").LtOrEq(123),
			Stmt: "`t1`.`c1` <= ?",
			Bind: []interface{}{int64(123)},
		},
		{
			Name: "CondIsNull",
			Cond: NewInt64Column(t1, "c1").IsNull(),
			Stmt: "`t1`.`c1` IS NULL",
			Bind: []interface{}{},
		},
		{
			Name: "CondIsNotNull",
			Cond: NewInt64Column(t1, "c1").IsNotNull(),
			Stmt: "`t1`.`c1` IS NOT NULL",
			Bind: []interface{}{},
		},
		{
			Name: "CondEqCol",
			Cond: NewInt64Column(t1, "c1").EqCol(NewInt8Column(t2, "c2")),
			Stmt: "`t1`.`c1` = `t2`.`c2`",
			Bind: []interface{}{},
		},
		{
			Name: "CondIn/One",
			Cond: NewInt64Column(t1, "c1").In(1),
			Stmt: "`t1`.`c1` IN (?)",
			Bind: []interface{}{int64(1)},
		},
		{
			Name: "CondIn/Two",
			Cond: NewInt64Column(t1, "c1").In(1, 2),
			Stmt: "`t1`.`c1` IN (?, ?)",
			Bind: []interface{}{int64(1), int64(2)},
		},
		{
			Name: "CondNotIn/One",
			Cond: NewInt64Column(t1, "c1").NotIn(1),
			Stmt: "`t1`.`c1` NOT IN (?)",
			Bind: []interface{}{int64(1)},
		},
		{
			Name: "CondNotIn/Two",
			Cond: NewInt64Column(t1, "c1").NotIn(1, 2),
			Stmt: "`t1`.`c1` NOT IN (?, ?)",
			Bind: []interface{}{int64(1), int64(2)},
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

func TestFloat32Column_Cond(t *testing.T) {
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
			Cond: NewFloat32Column(t1, "c1").Eq(123),
			Stmt: "`t1`.`c1` = ?",
			Bind: []interface{}{float32(123)},
		},
		{
			Name: "CondNotEq",
			Cond: NewFloat32Column(t1, "c1").NotEq(123),
			Stmt: "`t1`.`c1` != ?",
			Bind: []interface{}{float32(123)},
		},
		{
			Name: "CondGt",
			Cond: NewFloat32Column(t1, "c1").Gt(123),
			Stmt: "`t1`.`c1` > ?",
			Bind: []interface{}{float32(123)},
		},
		{
			Name: "CondGtOrEq",
			Cond: NewFloat32Column(t1, "c1").GtOrEq(123),
			Stmt: "`t1`.`c1` >= ?",
			Bind: []interface{}{float32(123)},
		},
		{
			Name: "CondLt",
			Cond: NewFloat32Column(t1, "c1").Lt(123),
			Stmt: "`t1`.`c1` < ?",
			Bind: []interface{}{float32(123)},
		},
		{
			Name: "CondLtOrEq",
			Cond: NewFloat32Column(t1, "c1").LtOrEq(123),
			Stmt: "`t1`.`c1` <= ?",
			Bind: []interface{}{float32(123)},
		},
		{
			Name: "CondIsNull",
			Cond: NewFloat32Column(t1, "c1").IsNull(),
			Stmt: "`t1`.`c1` IS NULL",
			Bind: []interface{}{},
		},
		{
			Name: "CondIsNotNull",
			Cond: NewFloat32Column(t1, "c1").IsNotNull(),
			Stmt: "`t1`.`c1` IS NOT NULL",
			Bind: []interface{}{},
		},
		{
			Name: "CondEqCol",
			Cond: NewFloat32Column(t1, "c1").EqCol(NewInt8Column(t2, "c2")),
			Stmt: "`t1`.`c1` = `t2`.`c2`",
			Bind: []interface{}{},
		},
		{
			Name: "CondIn/One",
			Cond: NewFloat32Column(t1, "c1").In(1),
			Stmt: "`t1`.`c1` IN (?)",
			Bind: []interface{}{float32(1)},
		},
		{
			Name: "CondIn/Two",
			Cond: NewFloat32Column(t1, "c1").In(1, 2),
			Stmt: "`t1`.`c1` IN (?, ?)",
			Bind: []interface{}{float32(1), float32(2)},
		},
		{
			Name: "CondNotIn/One",
			Cond: NewFloat32Column(t1, "c1").NotIn(1),
			Stmt: "`t1`.`c1` NOT IN (?)",
			Bind: []interface{}{float32(1)},
		},
		{
			Name: "CondNotIn/Two",
			Cond: NewFloat32Column(t1, "c1").NotIn(1, 2),
			Stmt: "`t1`.`c1` NOT IN (?, ?)",
			Bind: []interface{}{float32(1), float32(2)},
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

func TestFloat64Column_Cond(t *testing.T) {
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
			Cond: NewFloat64Column(t1, "c1").Eq(123),
			Stmt: "`t1`.`c1` = ?",
			Bind: []interface{}{float64(123)},
		},
		{
			Name: "CondNotEq",
			Cond: NewFloat64Column(t1, "c1").NotEq(123),
			Stmt: "`t1`.`c1` != ?",
			Bind: []interface{}{float64(123)},
		},
		{
			Name: "CondGt",
			Cond: NewFloat64Column(t1, "c1").Gt(123),
			Stmt: "`t1`.`c1` > ?",
			Bind: []interface{}{float64(123)},
		},
		{
			Name: "CondGtOrEq",
			Cond: NewFloat64Column(t1, "c1").GtOrEq(123),
			Stmt: "`t1`.`c1` >= ?",
			Bind: []interface{}{float64(123)},
		},
		{
			Name: "CondLt",
			Cond: NewFloat64Column(t1, "c1").Lt(123),
			Stmt: "`t1`.`c1` < ?",
			Bind: []interface{}{float64(123)},
		},
		{
			Name: "CondLtOrEq",
			Cond: NewFloat64Column(t1, "c1").LtOrEq(123),
			Stmt: "`t1`.`c1` <= ?",
			Bind: []interface{}{float64(123)},
		},
		{
			Name: "CondIsNull",
			Cond: NewFloat64Column(t1, "c1").IsNull(),
			Stmt: "`t1`.`c1` IS NULL",
			Bind: []interface{}{},
		},
		{
			Name: "CondIsNotNull",
			Cond: NewFloat64Column(t1, "c1").IsNotNull(),
			Stmt: "`t1`.`c1` IS NOT NULL",
			Bind: []interface{}{},
		},
		{
			Name: "CondEqCol",
			Cond: NewFloat64Column(t1, "c1").EqCol(NewInt8Column(t2, "c2")),
			Stmt: "`t1`.`c1` = `t2`.`c2`",
			Bind: []interface{}{},
		},
		{
			Name: "CondIn/One",
			Cond: NewFloat64Column(t1, "c1").In(1),
			Stmt: "`t1`.`c1` IN (?)",
			Bind: []interface{}{float64(1)},
		},
		{
			Name: "CondIn/Two",
			Cond: NewFloat64Column(t1, "c1").In(1, 2),
			Stmt: "`t1`.`c1` IN (?, ?)",
			Bind: []interface{}{float64(1), float64(2)},
		},
		{
			Name: "CondNotIn/One",
			Cond: NewFloat64Column(t1, "c1").NotIn(1),
			Stmt: "`t1`.`c1` NOT IN (?)",
			Bind: []interface{}{float64(1)},
		},
		{
			Name: "CondNotIn/Two",
			Cond: NewFloat64Column(t1, "c1").NotIn(1, 2),
			Stmt: "`t1`.`c1` NOT IN (?, ?)",
			Bind: []interface{}{float64(1), float64(2)},
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
