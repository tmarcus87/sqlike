package session

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/tmarcus87/sqlike/dialect"
	"github.com/tmarcus87/sqlike/model"
	"testing"
)

type Author struct {
	Id   int    `sqlike:"id"`
	Name string `sqlike:"name"`
}

func TestFetchMap(t *testing.T) {
	authorTable := model.NewTable("author")

	authorIdColumn := model.NewInt64Column(authorTable, "id")
	authorNameColumn := model.NewTextColumn(authorTable, "name")

	db, err := sql.Open(dialect.MySQL, "user:password@tcp(127.0.0.1:3306)/sqlike")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	asserts := assert.New(t)

	if _, err := db.Exec("TRUNCATE `author`"); err != nil {
		panic(err)
	}
	if _, err := db.Exec("INSERT INTO `author` (`name`) VALUES ('William Shakespeare'), ('J. K. Rowling')"); err != nil {
		panic(err)
	}

	s := NewSession(context.Background(), db, dialect.MySQL, false)

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
	authorTable := model.NewTable("author")

	authorIdColumn := model.NewInt64Column(authorTable, "id")
	authorNameColumn := model.NewTextColumn(authorTable, "name")

	db, err := sql.Open(dialect.MySQL, "user:password@tcp(127.0.0.1:3306)/sqlike")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	asserts := assert.New(t)

	if _, err := db.Exec("TRUNCATE `author`"); err != nil {
		panic(err)
	}
	if _, err := db.Exec("INSERT INTO `author` (`name`) VALUES ('William Shakespeare'), ('J. K. Rowling')"); err != nil {
		panic(err)
	}

	s := NewSession(context.Background(), db, dialect.MySQL, false)

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
	authorTable := model.NewTable("author")

	authorIdColumn := model.NewInt64Column(authorTable, "id")
	authorNameColumn := model.NewTextColumn(authorTable, "name")

	db, err := sql.Open(dialect.MySQL, "user:password@tcp(127.0.0.1:3306)/sqlike")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	asserts := assert.New(t)

	if _, err := db.Exec("TRUNCATE `author`"); err != nil {
		panic(err)
	}
	if _, err := db.Exec("INSERT INTO `author` (`name`) VALUES ('William Shakespeare'), ('J. K. Rowling')"); err != nil {
		panic(err)
	}

	s := NewSession(context.Background(), db, dialect.MySQL, false)

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
