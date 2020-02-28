package sqlike

import (
	"context"
	"database/sql"
	"fmt"
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

	t1 := &BasicTable{name: "t1"}

	c1 := &BasicColumn{table: t1, name: "c1"}
	c2 := &BasicColumn{table: t1, name: "c2"}

	s := &basicSession{db: &sql.DB{}}
	stmt, bindings := s.Select(c1, c2).From(t1).Build().StatementAndBindings()
	asserts.Equal("SELECT `t1`.`c1`, `t1`.`c2` FROM `t1`", stmt)
	asserts.Empty(bindings)
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

type Author struct {
	Id   int    `sqlike:"id"`
	Name string `sqlike:"name"`
}

func TestFetchMap(t *testing.T) {
	authorTable := &BasicTable{name: "author"}

	authorIdColumn := &BasicColumn{table: authorTable, name: "id"}
	authorNameColumn := &BasicColumn{table: authorTable, name: "name"}

	db, err := sql.Open(DialectMySQL, "user:password@tcp(127.0.0.1:3306)/sqlike")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	s := &basicSession{db: db, ctx: context.Background()}

	asserts := assert.New(t)

	mapslice, err := s.Select(authorIdColumn, authorNameColumn).From(authorTable).Build().FetchMap()
	asserts.Nil(err)
	asserts.Len(mapslice, 2)
	asserts.Equal("1", mapslice[0]["id"])
	asserts.Equal("William Shakespeare", mapslice[0]["name"])
	asserts.Equal("2", mapslice[1]["id"])
	asserts.Equal("J. K. Rowling", mapslice[1]["name"])
	fmt.Printf("[%v] %+v\n", err, mapslice)
}

func TestFetchInto(t *testing.T) {
	authorTable := &BasicTable{name: "author"}

	authorIdColumn := &BasicColumn{table: authorTable, name: "id"}
	authorNameColumn := &BasicColumn{table: authorTable, name: "name"}

	db, err := sql.Open(DialectMySQL, "user:password@tcp(127.0.0.1:3306)/sqlike")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	s := &basicSession{db: db, ctx: context.Background()}

	asserts := assert.New(t)

	t.Run("PtrValue", func(t *testing.T) {
		authors := make([]Author, 0)
		err := s.Select(authorIdColumn, authorNameColumn).From(authorTable).Build().FetchInto(&authors)
		asserts.Nil(err)
		asserts.Len(authors, 2)
		asserts.Equal(1, authors[0].Id)
		asserts.Equal("William Shakespeare", authors[0].Name)
		asserts.Equal(2, authors[1].Id)
		asserts.Equal("J. K. Rowling", authors[1].Name)
		fmt.Printf("[%v] %+v\n", err, authors)
	})

	t.Run("PtrPtr", func(t *testing.T) {
		authors := make([]*Author, 0)
		err := s.Select(authorIdColumn, authorNameColumn).From(authorTable).Build().FetchInto(&authors)
		asserts.Nil(err)
		asserts.Len(authors, 2)
		asserts.Equal(1, authors[0].Id)
		asserts.Equal("William Shakespeare", authors[0].Name)
		asserts.Equal(2, authors[1].Id)
		asserts.Equal("J. K. Rowling", authors[1].Name)
		fmt.Printf("[%v] %+v\n", err, authors)
	})

	t.Run("Value", func(t *testing.T) {
		authors := make([]Author, 0)
		err := s.Select(authorIdColumn, authorNameColumn).From(authorTable).Build().FetchInto(authors)
		asserts.NotNil(err)
		fmt.Printf("[%v] %+v\n", err, authors)
	})

}

func TestFetchOneInto(t *testing.T) {
	authorTable := &BasicTable{name: "author"}

	authorIdColumn := &BasicColumn{table: authorTable, name: "id"}
	authorNameColumn := &BasicColumn{table: authorTable, name: "name"}

	db, err := sql.Open(DialectMySQL, "user:password@tcp(127.0.0.1:3306)/sqlike")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	s := &basicSession{db: db, ctx: context.Background()}

	asserts := assert.New(t)

	t.Run("Found", func(t *testing.T) {
		author := Author{}
		stmt := s.Select(authorIdColumn, authorNameColumn).From(authorTable).Where(authorIdColumn.Eq(1)).Build()
		ok, err := stmt.FetchOneInto(&author)
		asserts.Nil(err)
		asserts.True(ok)
		asserts.Equal(1, author.Id)
		asserts.Equal("William Shakespeare", author.Name)
	})

	t.Run("NotFound", func(t *testing.T) {
		author := Author{}
		stmt := s.Select(authorIdColumn, authorNameColumn).From(authorTable).Where(authorIdColumn.Eq(-1)).Build()
		ok, err := stmt.FetchOneInto(&author)
		asserts.Nil(err)
		asserts.False(ok)
	})

	t.Run("InvalidType", func(t *testing.T) {
		v := make(map[string]interface{})
		stmt := s.Select(authorIdColumn, authorNameColumn).From(authorTable).Where(authorIdColumn.Eq(1)).Build()
		ok, err := stmt.FetchOneInto(&v)
		asserts.NotNil(err)
		asserts.False(ok)
	})

}
