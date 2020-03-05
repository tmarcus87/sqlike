package main

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tmarcus87/sqlike"
	"github.com/tmarcus87/sqlike/model"
	"log"
	"time"
)

func main() {
	engine, err :=
		sqlike.NewEngine(
			sqlike.FromHostAndPort("mysql", "127.0.0.1", 3306, "user", "password", "sqlike"))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := engine.Close(); err != nil {
			log.Printf("[Error] %+v\n", err)
		}
	}()

	type Book struct {
		Id       int64  `sqlike:"id"`
		Name     string `sqlike:"title"`
		AuthorId int64  `sqlike:"author_id"`
	}

	bookTable := &model.BasicTable{Name: "book"}
	//idColumn := &model.BasicColumn{Table: bookTable, Name: "id"}
	nameColumn := &model.BasicColumn{Table: bookTable, Name: "title"}
	authorIdColumn := &model.BasicColumn{Table: bookTable, Name: "author_id"}

	result :=
		engine.Master(context.Background()).
			InsertInto(bookTable).
			Columns(nameColumn, authorIdColumn).
			ValueStructs(
				&Book{
					Name:     "The Old Man and the Sea",
					AuthorId: time.Now().Unix(),
				},
			).
			Build().
			Execute()
	if result.Error() != nil {
		panic(err)
	}
	fmt.Println(result.AffectedRows())
	fmt.Println(result.LastInsertId())
}
