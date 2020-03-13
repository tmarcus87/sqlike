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
			Column: NewTextColumn(NewTable("tbl"), "col").SQLikeAs("col_alias"),
		},
		{
			Expect: "`tbl_alias`.`col`",
			Column: NewTextColumn(NewTable("tbl").SQLikeAs("tbl_alias"), "col"),
		},
		{
			Expect: "`tbl_alias`.`col` AS `col_alias`",
			Column: NewTextColumn(NewTable("tbl").SQLikeAs("tbl_alias"), "col").SQLikeAs("col_alias"),
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
			Column:     NewTextColumn(NewTable("tbl").SQLikeAs("tbl_alias"), "col"),
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
