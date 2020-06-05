package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAndCondition(t *testing.T) {
	var (
		cond1 Condition = &SingleValueCondition{
			Column:   NewBoolColumn(NewTable("tbl"), "bool"),
			Operator: "=",
			Value:    true,
		}
		cond2 Condition = &SingleValueCondition{
			Column:   NewTextColumn(NewTable("tbl"), "text"),
			Operator: "=",
			Value:    "text",
		}
		cond3 Condition = &SingleValueCondition{
			Column:   NewInt32Column(NewTable("tbl"), "number"),
			Operator: "=",
			Value:    3,
		}
	)

	asserts := assert.New(t)

	t.Run("Simple", func(t *testing.T) {
		var (
			stmt     string
			bindings []interface{}
		)
		cond1.And(cond2).And(cond3).Apply(&stmt, &bindings)

		asserts.Equal("((`tbl`.`bool` = ? AND `tbl`.`text` = ?) AND `tbl`.`number` = ?)", stmt)
		if asserts.Len(bindings, 3) {
			asserts.Equal(true, bindings[0])
			asserts.Equal("text", bindings[1])
			asserts.Equal(3, bindings[2])
		}
	})

	t.Run("Group", func(t *testing.T) {
		var (
			stmt     string
			bindings []interface{}
		)
		cond1.And(cond2.And(cond3)).Apply(&stmt, &bindings)

		asserts.Equal("(`tbl`.`bool` = ? AND (`tbl`.`text` = ? AND `tbl`.`number` = ?))", stmt)
		if asserts.Len(bindings, 3) {
			asserts.Equal(true, bindings[0])
			asserts.Equal("text", bindings[1])
			asserts.Equal(3, bindings[2])
		}
	})
}

func TestOrCondition(t *testing.T) {
	var (
		cond1 Condition = &SingleValueCondition{
			Column:   NewBoolColumn(NewTable("tbl"), "bool"),
			Operator: "=",
			Value:    true,
		}
		cond2 Condition = &SingleValueCondition{
			Column:   NewTextColumn(NewTable("tbl"), "text"),
			Operator: "=",
			Value:    "text",
		}
		cond3 Condition = &SingleValueCondition{
			Column:   NewInt32Column(NewTable("tbl"), "number"),
			Operator: "=",
			Value:    3,
		}
	)

	asserts := assert.New(t)

	t.Run("Simple", func(t *testing.T) {
		var (
			stmt     string
			bindings []interface{}
		)
		cond1.Or(cond2).Or(cond3).Apply(&stmt, &bindings)

		asserts.Equal("((`tbl`.`bool` = ? OR `tbl`.`text` = ?) OR `tbl`.`number` = ?)", stmt)
		if asserts.Len(bindings, 3) {
			asserts.Equal(true, bindings[0])
			asserts.Equal("text", bindings[1])
			asserts.Equal(3, bindings[2])
		}
	})

	t.Run("Group", func(t *testing.T) {
		var (
			stmt     string
			bindings []interface{}
		)
		cond1.Or(cond2.Or(cond3)).Apply(&stmt, &bindings)

		asserts.Equal("(`tbl`.`bool` = ? OR (`tbl`.`text` = ? OR `tbl`.`number` = ?))", stmt)
		if asserts.Len(bindings, 3) {
			asserts.Equal(true, bindings[0])
			asserts.Equal("text", bindings[1])
			asserts.Equal(3, bindings[2])
		}
	})
}
