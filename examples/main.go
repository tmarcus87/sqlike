package main

import (
	"fmt"
	"github.com/tmarcus87/sqlike"
	"github.com/tmarcus87/sqlike/logger"
	"github.com/tmarcus87/sqlike/model"
	"log"
	"os"
)

type Author struct {
	Id   int64  `sqlike:"id"`
	Name string `sqlike:"name"`
}

type Book struct {
	Id       int64  `sqlike:"id"`
	Title    string `sqlike:"title"`
	AuthorId int64  `sqlike:"author_id"`
}

var (
	examples = make(map[string]func(e sqlike.Engine) error)

	authorTable        = model.NewTable("author")
	authorIdColumn     = model.NewInt64Column(authorTable, "id")
	authorNameColumn   = model.NewTextColumn(authorTable, "name")
	bookTable          = model.NewTable("book")
	bookIdColumn       = model.NewInt64Column(bookTable, "id")
	bookTitleColumn    = model.NewTextColumn(bookTable, "title")
	bookAuthorIdColumn = model.NewTextColumn(bookTable, "author_id")
)

func main() {
	logger.SetLevel(logger.DebugLevel)

	var name string

	if len(os.Args) >= 2 {
		name = os.Args[1]
	} else {
		fmt.Printf("Example name? ")
		if _, err := fmt.Scan(&name); err != nil {
			log.Fatalf("failed to read from stdin : %+v", err)
		}
	}

	fn, ok := examples[name]
	if !ok {
		log.Fatalf("Example(%s) is not found", name)
	}

	fmt.Printf("Run '%s'\n--------\n", name)

	engine, err :=
		sqlike.NewEngine(
			sqlike.FromHostAndPort("mysql", "127.0.0.1", 3306, "user", "password", "sqlike"))
	if err != nil {
		log.Fatalf("Failed to create engine : %+v", err)
	}
	defer func() {
		if err := engine.Close(); err != nil {
			log.Printf("[Error] %+v\n", err)
		}
	}()

	if err := fn(engine); err != nil {
		log.Fatalf("Failed to execute example(%s) : %+v", name, err)
	}
}
