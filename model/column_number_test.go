package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumericColumn_SQLikeFieldExpr(t *testing.T) {
	tbl := NewTable("tbl")
	tblAlias := NewTable("tbl").SQLikeAs("tbl_alias")

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
			Column: (&NumberColumn{Table: tbl, Name: "col"}).SQLikeAs("col_alias"),
		},
		{
			Expect: "`tbl_alias`.`col`",
			Column: &NumberColumn{Table: tblAlias, Name: "col"},
		},
		{
			Expect: "`tbl_alias`.`col` AS `col_alias`",
			Column: (&NumberColumn{Table: tblAlias, Name: "col"}).SQLikeAs("col_alias"),
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
			Column: (&NumberColumn{Table: tbl, Name: "col"}).PlusInt(123).SQLikeAs("col_plus"),
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
			Column:     NewInt8Column(NewTable("tbl").SQLikeAs("tbl_alias"), "col"),
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
			Column:     NewInt16Column(NewTable("tbl").SQLikeAs("tbl_alias"), "col"),
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
			Column:     NewInt32Column(NewTable("tbl").SQLikeAs("tbl_alias"), "col"),
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
			Column:     NewInt64Column(NewTable("tbl").SQLikeAs("tbl_alias"), "col"),
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
			Column:     NewFloat32Column(NewTable("tbl").SQLikeAs("tbl_alias"), "col"),
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
			Column:     NewFloat64Column(NewTable("tbl").SQLikeAs("tbl_alias"), "col"),
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
