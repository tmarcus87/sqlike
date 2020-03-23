package main

import (
	"context"
	"fmt"
	"github.com/tmarcus87/sqlike"
	"github.com/tmarcus87/sqlike/model"
)

func init() {
	examples["join"] = join
}

func join(e sqlike.Engine) error {

	type Author struct {
		Id   int64  `sqlike:"id"`
		Name string `sqlike:"name"`
	}

	type Book struct {
		Id       int64  `sqlike:"id"`
		Title    string `sqlike:"title"`
		AuthorId int64  `sqlike:"author_id"`
	}

	authorTable := model.NewTable("author")
	authorIdColumn := model.NewInt64Column(authorTable, "id")
	authorNameColumn := model.NewTextColumn(authorTable, "name")
	bookTable := model.NewTable("book")
	bookIdColumn := model.NewInt64Column(bookTable, "id")
	bookTitleColumn := model.NewTextColumn(bookTable, "title")
	bookAuthorIdColumn := model.NewTextColumn(bookTable, "author_id")

	if err := e.Master(context.Background()).Truncate(authorTable).Build().Execute().Error(); err != nil {
		return fmt.Errorf("failed to truncate : %w", err)
	}
	if err := e.Master(context.Background()).Truncate(bookTable).Build().Execute().Error(); err != nil {
		return fmt.Errorf("failed to truncate : %w", err)
	}

	// Insert author
	{
		for _, td := range []Author{
			{Id: 1, Name: "Lev Nikolayevich Tolstoy"},
			{Id: 2, Name: "Georges Simenon"},
			{Id: 3, Name: "J. K. Rowling"},
		} {

			if err :=
				e.Master(context.Background()).
					InsertInto(authorTable).
					Columns(authorNameColumn).
					ValueStructs(&td).
					Build().
					Execute().
					Error(); err != nil {
				return fmt.Errorf("failed to insert : %w", err)
			}

			author := Author{}
			ok, err :=
				e.Master(context.Background()).
					SelectFrom(authorTable).
					Where(authorIdColumn.CondEq(td.Id)).
					Build().
					FetchOneInto(&author)
			if err != nil {
				return fmt.Errorf("unexpected error : %w", err)
			}
			if !ok {
				return fmt.Errorf("no record(%d)", td.Id)
			}
			if author.Name != td.Name {
				return fmt.Errorf("author name is mismatched : %+v", author)
			}
			fmt.Printf("%+v\n", author)
		}
	}

	// Insert book
	{
		for _, td := range []Book{
			{Id: 1, Title: "Childhood", AuthorId: 1},
			{Id: 2, Title: "Boyhood", AuthorId: 1},
			{Id: 3, Title: "Youth", AuthorId: 1},
			{Id: 4, Title: "Sevastopol Sketches", AuthorId: 1},
			{Id: 5, Title: "The Cossacks", AuthorId: 1},
			{Id: 6, Title: "Family Happiness", AuthorId: 1},
			{Id: 7, Title: "Pietre-le-Letton", AuthorId: 2},
			{Id: 8, Title: "Monsieur Gallet décédé", AuthorId: 2},
			{Id: 9, Title: "Le Pendu de Saint-Phollien", AuthorId: 2},
			{Id: 10, Title: "Le Charretier de la Providence", AuthorId: 2},
			{Id: 11, Title: "La Tête d'un homme", AuthorId: 2},
			{Id: 12, Title: "Le Chien jaune", AuthorId: 2},
			{Id: 13, Title: "Harry Potter and the Philosopher's Stone", AuthorId: 3},
			{Id: 14, Title: "Harry Potter and the Chamber of Secrets", AuthorId: 3},
			{Id: 15, Title: "Harry Potter and the Prisoner of Azkaban", AuthorId: 3},
			{Id: 16, Title: "Harry Potter and the Goblet of Fire", AuthorId: 3},
			{Id: 17, Title: "Harry Potter and the Order of the Phoenix", AuthorId: 3},
			{Id: 18, Title: "Harry Potter and the Half-Blood Prince", AuthorId: 3},
			{Id: 19, Title: "Harry Potter and the Deathly Hallows", AuthorId: 3},
		} {
			if err :=
				e.Master(context.Background()).
					InsertInto(bookTable).
					Columns(bookTitleColumn, bookAuthorIdColumn).
					ValueStructs(&td).
					Build().
					Execute().
					Error(); err != nil {
				return fmt.Errorf("failed to insert : %w", err)
			}

			book := Book{}
			ok, err :=
				e.Master(context.Background()).
					SelectFrom(bookTable).
					Where(bookIdColumn.CondEq(td.Id)).
					Build().
					FetchOneInto(&book)
			if err != nil {
				return fmt.Errorf("unexpected error : %w", err)
			}
			if !ok {
				return fmt.Errorf("no record(%d)", td.Id)
			}
			if book.Title != td.Title {
				return fmt.Errorf("author name is mismatched : %+v", book)
			}
			fmt.Printf("%+v\n", book)
		}
	}

	{
		type BookAndAuthor struct {
			BookId     int64  `sqlike:"book_id"`
			BookTitle  string `sqlike:"book_title"`
			AuthorName string `sqlike:"author_name"`
		}

		books := make([]BookAndAuthor, 0)

		if err :=
			e.Master(context.Background()).
				Select(
					bookIdColumn.As("book_id"),
					bookTitleColumn.As("book_title"),
					authorNameColumn.As("author_name")).
				From(bookTable).
				InnerJoin(authorTable, bookAuthorIdColumn.CondEqCol(authorIdColumn)).
				Where(authorIdColumn.CondEq(2)).
				OrderBy(bookTitleColumn.Asc()).
				Build().
				FetchInto(&books); err != nil {
			return fmt.Errorf("failed to query : %w", err)
		}

		for _, book := range books {
			fmt.Printf("%+v\n", book)
		}
	}
	return nil
}
