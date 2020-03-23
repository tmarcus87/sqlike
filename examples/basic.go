package main

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tmarcus87/sqlike"
	"github.com/tmarcus87/sqlike/model"
	"time"
)

func init() {
	examples["basic"] = basic
}

func basic(e sqlike.Engine) error {
	type Book struct {
		Id       int64  `sqlike:"id"`
		Title    string `sqlike:"title"`
		AuthorId int64  `sqlike:"author_id"`
	}

	bookTable := &model.BasicTable{Name: "book"}
	nameColumn := model.NewTextColumn(bookTable, "title")
	authorIdColumn := model.NewInt64Column(bookTable, "author_id")

	// truncate
	if err := e.Master(context.Background()).Truncate(bookTable).Build().Execute().Error(); err != nil {
		return fmt.Errorf("failed to truncate : %+v", err)
	}

	ctx := context.Background()

	// Check no records
	{
		books := make([]*Book, 0)
		err :=
			e.Auto(ctx).
				SelectFrom(bookTable).
				Build().
				FetchInto(&books)
		if err != nil {
			return fmt.Errorf("failed to select : %w", err)
		}
		if len(books) != 0 {
			return fmt.Errorf("unexpected number of rows : %d", len(books))
		}
	}

	// Insert records w/ auto increment
	{
		result :=
			e.Master(ctx).
				InsertInto(bookTable).
				Columns(nameColumn, authorIdColumn).
				ValueStructs(
					&Book{
						Title:    "The Old Man and the Sea",
						AuthorId: time.Now().Unix() + 1,
					},
					&Book{
						Title:    "Hamlet",
						AuthorId: time.Now().Unix() + 1,
					},
				).
				Build().
				Execute()
		if result.Error() != nil {
			return result.Error()
		}
		fmt.Print("AffectedRows: ")
		fmt.Println(result.AffectedRows())
		fmt.Print("LastInsertId: ")
		fmt.Println(result.LastInsertId())
	}

	// Select w/ struct
	{
		books := make([]*Book, 0)
		err :=
			e.Auto(ctx).
				SelectFrom(bookTable).
				Build().
				FetchInto(&books)
		if err != nil {
			return fmt.Errorf("failed to select : %w", err)
		}
		if len(books) != 2 {
			return fmt.Errorf("unexpected number of rows : %d", len(books))
		}
		for _, book := range books {
			fmt.Printf("* %+v\n", *book)
		}
	}

	return nil
}
