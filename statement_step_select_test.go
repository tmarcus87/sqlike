package sqlike

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildExplain(t *testing.T) {
	asserts := assert.New(t)

	s := &basicSession{dialect: DialectMySQL, db: &sql.DB{}}

	{
		stmt, _ := s.Explain().SelectOne().Build().StatementAndBindings()
		asserts.Equal("EXPLAIN SELECT 1 FROM dual", stmt)
	}

	{
		t1 := &BasicTable{name: "t1"}

		c1 := &BasicColumn{table: t1, name: "c1"}
		c2 := &BasicColumn{table: t1, name: "c2"}

		stmt, _ := s.Explain().Select(c1, c2).From(t1).Build().StatementAndBindings()
		asserts.Equal("EXPLAIN SELECT `t1`.`c1`, `t1`.`c2` FROM `t1`", stmt)
	}
}

func TestBuildSelectOne(t *testing.T) {
	tests := []struct {
		dialect string
		expect  string
	}{
		{
			dialect: DialectMySQL,
			expect:  "SELECT 1 FROM dual",
		},
		{
			dialect: DialectSqlite3,
			expect:  "SELECT 1",
		},
	}

	for _, test := range tests {
		t.Run(test.dialect, func(t *testing.T) {
			asserts := assert.New(t)

			s := &basicSession{dialect: test.dialect, db: &sql.DB{}}
			stmt, _ := s.SelectOne().Build().StatementAndBindings()
			asserts.Equal(test.expect, stmt)
		})
	}
}

func TestBuildSelectFrom(t *testing.T) {
	asserts := assert.New(t)

	t.Run("WithoutAs", func(t *testing.T) {
		t1 := &BasicTable{name: "t1"}

		c1 := &BasicColumn{table: t1, name: "c1"}
		c2 := &BasicColumn{table: t1, name: "c2"}

		s := &basicSession{db: &sql.DB{}}
		stmt, bindings := s.Select(c1, c2).From(t1).Build().StatementAndBindings()
		asserts.Equal("SELECT `t1`.`c1`, `t1`.`c2` FROM `t1`", stmt)
		asserts.Empty(bindings)
	})

	t.Run("WithAs", func(t *testing.T) {
		asserts := assert.New(t)
		asserts.Fail("not yet tested")
	})
}

func TestBuildSelectFromWithOneWhere(t *testing.T) {
	asserts := assert.New(t)

	t1 := &BasicTable{name: "t1"}

	c1 := &BasicColumn{table: t1, name: "c1"}
	c2 := &BasicColumn{table: t1, name: "c2"}

	s := &basicSession{db: &sql.DB{}}
	stmt, bindings := s.Select(c1, c2).From(t1).Where(c1.Eq(1)).Build().StatementAndBindings()
	asserts.Equal("SELECT `t1`.`c1`, `t1`.`c2` FROM `t1` WHERE `t1`.`c1` = ?", stmt)
	asserts.Len(bindings, 1)
	asserts.Equal(1, bindings[0])
}

func TestBuildSelectFromWithTwoWhere(t *testing.T) {
	asserts := assert.New(t)

	t1 := &BasicTable{name: "t1"}

	c1 := &BasicColumn{table: t1, name: "c1"}
	c2 := &BasicColumn{table: t1, name: "c2"}

	t.Run("And", func(t *testing.T) {
		s := &basicSession{db: &sql.DB{}}
		stmt, bindings := s.Select(c1, c2).From(t1).Where(And(c1.Eq(1), c2.Eq(2))).Build().StatementAndBindings()
		asserts.Equal("SELECT `t1`.`c1`, `t1`.`c2` FROM `t1` WHERE (`t1`.`c1` = ? AND `t1`.`c2` = ?)", stmt)
		asserts.Len(bindings, 2)
		asserts.Equal(1, bindings[0])
		asserts.Equal(2, bindings[1])
	})

	t.Run("Or", func(t *testing.T) {
		s := &basicSession{db: &sql.DB{}}
		stmt, bindings := s.Select(c1, c2).From(t1).Where(Or(c1.Eq(1), c2.Eq(2))).Build().StatementAndBindings()
		asserts.Equal("SELECT `t1`.`c1`, `t1`.`c2` FROM `t1` WHERE (`t1`.`c1` = ? OR `t1`.`c2` = ?)", stmt)
		asserts.Len(bindings, 2)
		asserts.Equal(1, bindings[0])
		asserts.Equal(2, bindings[1])
	})
}

func TestBuildSelectFromJoin(t *testing.T) {
	t.Run("WithoutAs", func(t *testing.T) {
		asserts := assert.New(t)

		t1 := &BasicTable{name: "t1"}
		t2 := &BasicTable{name: "t2"}

		c1 := &BasicColumn{table: t1, name: "c1"}
		c2 := &BasicColumn{table: t1, name: "c2"}
		c3 := &BasicColumn{table: t2, name: "c3"}
		c4 := &BasicColumn{table: t2, name: "c4"}

		s := &basicSession{db: &sql.DB{}}

		stmt, bindings := s.Select(c1, c2).From(t1).LeftOuterJoin(t2, c1.EqCol(c3)).Where(c4.Eq(1)).Build().StatementAndBindings()
		asserts.Equal(
			"SELECT `t1`.`c1`, `t1`.`c2` FROM `t1` LEFT OUTER JOIN `t2` ON `t1`.`c1` = `t2`.`c3` WHERE `t2`.`c4` = ?",
			stmt)
		asserts.Len(bindings, 1)
		asserts.Equal(1, bindings[0])
	})

	t.Run("WithAs", func(t *testing.T) {
		asserts := assert.New(t)
		asserts.Fail("not yet tested")
	})
}

func TestBuildSelectFromGroupBy(t *testing.T) {
	t.Run("WithoutAs", func(t *testing.T) {
		asserts := assert.New(t)

		t1 := &BasicTable{name: "t1"}

		c1 := &BasicColumn{table: t1, name: "c1"}
		c2 := &BasicColumn{table: t1, name: "c2"}

		s := &basicSession{db: &sql.DB{}}

		stmt, _ := s.Select(c1, Count(c2)).From(t1).GroupBy(c1).Build().StatementAndBindings()
		asserts.Equal("SELECT `t1`.`c1`, COUNT(`t1`.`c2`) FROM `t1` GROUP BY `t1`.`c1`", stmt)
	})

	t.Run("WithAs", func(t *testing.T) {
		asserts := assert.New(t)
		asserts.Fail("not yet tested")
	})
}

func TestBuildSelectFromOrderBy(t *testing.T) {
	t.Run("WithoutAs", func(t *testing.T) {
		asserts := assert.New(t)

		t1 := &BasicTable{name: "t1"}

		c1 := &BasicColumn{table: t1, name: "c1"}
		c2 := &BasicColumn{table: t1, name: "c2"}

		s := &basicSession{db: &sql.DB{}}

		stmt, _ := s.Select(c1, c2).From(t1).OrderBy(Order(c2, OrderDesc)).Build().StatementAndBindings()
		asserts.Equal("SELECT `t1`.`c1`, `t1`.`c2` FROM `t1` ORDER BY `t1`.`c2` DESC", stmt)
	})

	t.Run("WithAs", func(t *testing.T) {
		asserts := assert.New(t)
		asserts.Fail("not yet tested")
	})

}

func TestBuildSelectFromLimitAndOffset(t *testing.T) {

	t.Run("Limit", func(t *testing.T) {
		asserts := assert.New(t)

		t1 := &BasicTable{name: "t1"}

		c1 := &BasicColumn{table: t1, name: "c1"}

		s := &basicSession{db: &sql.DB{}}

		stmt, _ := s.Select(c1).From(t1).LimitAndOffset(10, 0).Build().StatementAndBindings()
		asserts.Equal("SELECT `t1`.`c1` FROM `t1` LIMIT 10", stmt)
	})

	t.Run("LimitAndOffset", func(t *testing.T) {
		asserts := assert.New(t)

		t1 := &BasicTable{name: "t1"}

		c1 := &BasicColumn{table: t1, name: "c1"}

		s := &basicSession{db: &sql.DB{}}

		stmt, _ := s.Select(c1).From(t1).LimitAndOffset(10, 1).Build().StatementAndBindings()
		asserts.Equal("SELECT `t1`.`c1` FROM `t1` LIMIT 10 OFFSET 1", stmt)
	})
}
