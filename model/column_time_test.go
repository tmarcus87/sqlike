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
