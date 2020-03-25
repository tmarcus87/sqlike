package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoolColumn_SQLikeFieldExpr(t *testing.T) {

	tests := []struct {
		Expect string
		Column ColumnField
	}{
		{
			Expect: "`tbl`.`col`",
			Column: NewBoolColumn(NewTable("tbl"), "col"),
		},
		{
			Expect: "`tbl`.`col` AS `col_alias`",
			Column: NewBoolColumn(NewTable("tbl"), "col").As("col_alias"),
		},
		{
			Expect: "`tbl_alias`.`col`",
			Column: NewBoolColumn(NewTable("tbl").As("tbl_alias"), "col"),
		},
		{
			Expect: "`tbl_alias`.`col` AS `col_alias`",
			Column: NewBoolColumn(NewTable("tbl").As("tbl_alias"), "col").As("col_alias"),
		},
	}

	for _, test := range tests {
		t.Run(test.Expect, func(t *testing.T) {
			asserts := assert.New(t)

			asserts.Equal(test.Expect, test.Column.FieldExpr())
		})
	}

}

func TestBoolColumn_SetAndColumnValue(t *testing.T) {

	tests := []struct {
		ExpectExpr string
		Column     *BoolColumn
	}{
		{
			ExpectExpr: "`tbl`.`col`",
			Column:     NewBoolColumn(NewTable("tbl"), "col"),
		},
		{
			ExpectExpr: "`tbl_alias`.`col`",
			Column:     NewBoolColumn(NewTable("tbl").As("tbl_alias"), "col"),
		},
	}

	for _, test := range tests {
		t.Run(test.ExpectExpr, func(t *testing.T) {
			asserts := assert.New(t)

			colV := test.Column.Value(false)
			asserts.Equal(false, colV.ColumnValue())

			test.Column.Value(true)
			asserts.Equal(true, colV.ColumnValue())
		})
	}

}
