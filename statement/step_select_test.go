package statement

import (
	"github.com/stretchr/testify/assert"
	"github.com/tmarcus87/sqlike/dialect"
	"github.com/tmarcus87/sqlike/model"
	"testing"
)

func root(d string) *RootStep {
	return &RootStep{
		dialectStatement: dialect.GetDialectStatements(d),
	}
}

func TestExplain_Accept(t *testing.T) {
	asserts := assert.New(t)

	{

		stmt, _, err := NewExplainSelectBranchStep(root(dialect.MySQL)).SelectOne().Build().StatementAndBindings()
		asserts.Nil(err)
		asserts.Equal("EXPLAIN SELECT 1 FROM dual", stmt)
	}

	{
		t1 := model.NewTable("t1")

		c1 := model.NewBoolColumn(t1, "c1")
		c2 := model.NewBoolColumn(t1, "c2")

		stmt, _, err := NewExplainSelectBranchStep(root(dialect.MySQL)).Select(c1, c2).From(t1).Build().StatementAndBindings()
		asserts.Nil(err)
		asserts.Equal("EXPLAIN SELECT `t1`.`c1`, `t1`.`c2` FROM `t1`", stmt)
	}
}

func TestSelectOne_Accept(t *testing.T) {
	tests := []struct {
		dialect string
		expect  string
	}{
		{
			dialect: dialect.MySQL,
			expect:  "SELECT 1 FROM dual",
		},
		{
			dialect: dialect.Sqlite3,
			expect:  "SELECT 1",
		},
	}

	for _, test := range tests {
		t.Run(test.dialect, func(t *testing.T) {
			asserts := assert.New(t)

			stmt, _, err := NewSelectOneBranchStep(root(test.dialect)).Build().StatementAndBindings()
			asserts.Nil(err)
			asserts.Equal(test.expect, stmt)
		})
	}
}

func TestSelectFrom_Accept(t *testing.T) {
	asserts := assert.New(t)

	t.Run("WithoutAs", func(t *testing.T) {
		t1 := model.NewTable("t1")

		c1 := model.NewBoolColumn(t1, "c1")
		c2 := model.NewBoolColumn(t1, "c2")

		stmt, bindings, err := NewSelectColumnBranchStep(root(dialect.MySQL), c1, c2).From(t1).Build().StatementAndBindings()
		asserts.Nil(err)
		asserts.Equal("SELECT `t1`.`c1`, `t1`.`c2` FROM `t1`", stmt)
		asserts.Empty(bindings)
	})

	t.Run("WithAs", func(t *testing.T) {
		asserts := assert.New(t)

		t1 := model.NewTable("t1")

		c1 := model.NewBoolColumn(t1, "c1")
		c2 := model.NewBoolColumn(t1, "c2")

		stmt, bindings, err := NewSelectColumnBranchStep(root(dialect.MySQL), c1.SQLikeAs("c1alt"), c2.SQLikeAs("c2alt")).From(t1.SQLikeAs("t1alt")).Build().StatementAndBindings()
		asserts.Nil(err)
		asserts.Equal("SELECT `t1alt`.`c1` AS `c1alt`, `t1alt`.`c2` AS `c2alt` FROM `t1` AS `t1alt`", stmt)
		asserts.Empty(bindings)

	})
}

func TestSelectFromWithOneWhere_Accept(t *testing.T) {
	asserts := assert.New(t)

	t1 := model.NewTable("t1")

	c1 := model.NewBoolColumn(t1, "c1")
	c2 := model.NewBoolColumn(t1, "c2")

	stmt, bindings, err := NewSelectColumnBranchStep(root(dialect.MySQL), c1, c2).From(t1).Where(c1.CondEq(true)).Build().StatementAndBindings()
	asserts.Nil(err)
	asserts.Equal("SELECT `t1`.`c1`, `t1`.`c2` FROM `t1` WHERE `t1`.`c1` = ?", stmt)
	asserts.Len(bindings, 1)
	asserts.Equal(true, bindings[0])
}

func TestSelectFromWithTwoWhere_Accept(t *testing.T) {
	asserts := assert.New(t)

	t1 := model.NewTable("t1")

	c1 := model.NewBoolColumn(t1, "c1")
	c2 := model.NewBoolColumn(t1, "c2")

	t.Run("And", func(t *testing.T) {
		stmt, bindings, err := NewSelectColumnBranchStep(root(dialect.MySQL), c1, c2).From(t1).Where(And(c1.CondEq(true), c2.CondEq(false))).Build().StatementAndBindings()
		asserts.Nil(err)
		asserts.Equal("SELECT `t1`.`c1`, `t1`.`c2` FROM `t1` WHERE (`t1`.`c1` = ? AND `t1`.`c2` = ?)", stmt)
		asserts.Len(bindings, 2)
		asserts.Equal(true, bindings[0])
		asserts.Equal(false, bindings[1])
	})

	t.Run("Or", func(t *testing.T) {
		stmt, bindings, err := NewSelectColumnBranchStep(root(dialect.MySQL), c1, c2).From(t1).Where(Or(c1.CondEq(true), c2.CondEq(false))).Build().StatementAndBindings()
		asserts.Nil(err)
		asserts.Equal("SELECT `t1`.`c1`, `t1`.`c2` FROM `t1` WHERE (`t1`.`c1` = ? OR `t1`.`c2` = ?)", stmt)
		asserts.Len(bindings, 2)
		asserts.Equal(true, bindings[0])
		asserts.Equal(false, bindings[1])
	})
}

func TestSelectFromJoin_Accept(t *testing.T) {
	t.Run("WithoutAs", func(t *testing.T) {
		asserts := assert.New(t)

		t1 := model.NewTable("t1")
		t2 := model.NewTable("t2")

		c1 := model.NewBoolColumn(t1, "c1")
		c2 := model.NewBoolColumn(t1, "c2")
		c3 := model.NewBoolColumn(t2, "c3")
		c4 := model.NewBoolColumn(t2, "c4")

		stmt, bindings, err := NewSelectColumnBranchStep(root(dialect.MySQL), c1, c2).From(t1).LeftOuterJoin(t2, c1.CondEqCol(c3)).Where(c4.CondEq(true)).Build().StatementAndBindings()
		asserts.Nil(err)
		asserts.Equal(
			"SELECT `t1`.`c1`, `t1`.`c2` FROM `t1` LEFT OUTER JOIN `t2` ON `t1`.`c1` = `t2`.`c3` WHERE `t2`.`c4` = ?",
			stmt)
		asserts.Len(bindings, 1)
		asserts.Equal(true, bindings[0])
	})

	t.Run("WithAs", func(t *testing.T) {
		asserts := assert.New(t)

		t1 := model.NewTable("t1")
		t2 := model.NewTable("t2")

		c1 := model.NewBoolColumn(t1, "c1")
		c2 := model.NewBoolColumn(t1, "c2")
		c3 := model.NewBoolColumn(t2, "c3")
		c4 := model.NewBoolColumn(t2, "c4")

		stmt, bindings, err :=
			NewSelectColumnBranchStep(root(dialect.MySQL),
				c1.SQLikeAs("c1alt"),
				c2.SQLikeAs("c2alt")).
				From(t1.SQLikeAs("t1alt")).
				LeftOuterJoin(t2.SQLikeAs("t2alt"), c1.CondEqCol(c3)).
				Where(c4.CondEq(true)).
				Build().
				StatementAndBindings()
		asserts.Nil(err)
		asserts.Equal(
			"SELECT `t1alt`.`c1` AS `c1alt`, `t1alt`.`c2` AS `c2alt` FROM `t1` AS `t1alt` LEFT OUTER JOIN `t2` AS `t2alt` ON `t1alt`.`c1` = `t2alt`.`c3` WHERE `t2alt`.`c4` = ?",
			stmt)
		asserts.Len(bindings, 1)
		asserts.Equal(true, bindings[0])
	})
}

func TestSelectFromGroupBy_Accept(t *testing.T) {
	t.Run("WithoutAs", func(t *testing.T) {
		asserts := assert.New(t)

		t1 := model.NewTable("t1")

		c1 := model.NewBoolColumn(t1, "c1")
		c2 := model.NewBoolColumn(t1, "c2")

		stmt, _, err := NewSelectColumnBranchStep(root(dialect.MySQL), c1, model.Count(c2)).From(t1).GroupBy(c1).Build().StatementAndBindings()
		asserts.Nil(err)
		asserts.Equal("SELECT `t1`.`c1`, COUNT(`t1`.`c2`) FROM `t1` GROUP BY `t1`.`c1`", stmt)
	})

	t.Run("WithAs", func(t *testing.T) {
		asserts := assert.New(t)

		t1 := model.NewTable("t1")

		c1 := model.NewBoolColumn(t1, "c1")
		c2 := model.NewBoolColumn(t1, "c2")

		stmt, _, err := NewSelectColumnBranchStep(root(dialect.MySQL), c1.SQLikeAs("c1alt"), model.Count(c2).SQLikeAs("cnt")).From(t1.SQLikeAs("t1alt")).GroupBy(c1).Build().StatementAndBindings()
		asserts.Nil(err)
		asserts.Equal("SELECT `t1alt`.`c1` AS `c1alt`, COUNT(`t1alt`.`c2`) AS `cnt` FROM `t1` AS `t1alt` GROUP BY `t1alt`.`c1`", stmt)
	})
}

func TestSelectFromOrderBy_Accept(t *testing.T) {
	t.Run("WithoutAs", func(t *testing.T) {
		asserts := assert.New(t)

		t1 := model.NewTable("t1")

		c1 := model.NewBoolColumn(t1, "c1")
		c2 := model.NewBoolColumn(t1, "c2")

		stmt, _, err := NewSelectColumnBranchStep(root(dialect.MySQL), c1, c2).From(t1).OrderBy(&model.SortOrder{Column: c2, Order: model.OrderDesc}).Build().StatementAndBindings()
		asserts.Nil(err)
		asserts.Equal("SELECT `t1`.`c1`, `t1`.`c2` FROM `t1` ORDER BY `t1`.`c2` DESC", stmt)
	})

	t.Run("WithAs", func(t *testing.T) {
		asserts := assert.New(t)

		t1 := model.NewTable("t1")

		c1 := model.NewBoolColumn(t1, "c1")
		c2 := model.NewBoolColumn(t1, "c2")

		stmt, _, err := NewSelectColumnBranchStep(root(dialect.MySQL), c1.SQLikeAs("c1alt"), c2.SQLikeAs("c2alt")).From(t1.SQLikeAs("t1alt")).OrderBy(&model.SortOrder{Column: c2, Order: model.OrderDesc}).Build().StatementAndBindings()
		asserts.Nil(err)
		asserts.Equal("SELECT `t1alt`.`c1` AS `c1alt`, `t1alt`.`c2` AS `c2alt` FROM `t1` AS `t1alt` ORDER BY `t1alt`.`c2alt` DESC", stmt)
	})

}

func TestSelectFromLimitAndOffset_Accept(t *testing.T) {

	t.Run("Limit", func(t *testing.T) {
		asserts := assert.New(t)

		t1 := model.NewTable("t1")

		c1 := model.NewBoolColumn(t1, "c1")

		stmt, _, err := NewSelectColumnBranchStep(root(dialect.MySQL), c1).From(t1).LimitAndOffset(10, 0).Build().StatementAndBindings()
		asserts.Nil(err)
		asserts.Equal("SELECT `t1`.`c1` FROM `t1` LIMIT 10", stmt)
	})

	t.Run("LimitAndOffset", func(t *testing.T) {
		asserts := assert.New(t)

		t1 := model.NewTable("t1")

		c1 := model.NewBoolColumn(t1, "c1")

		stmt, _, err := NewSelectColumnBranchStep(root(dialect.MySQL), c1).From(t1).LimitAndOffset(10, 1).Build().StatementAndBindings()
		asserts.Nil(err)
		asserts.Equal("SELECT `t1`.`c1` FROM `t1` LIMIT 10 OFFSET 1", stmt)
	})
}
