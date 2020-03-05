SQLike
=======

SQL-like syntax O/R mapper for Golang inspired by [jOOQ](https://www.jooq.org/)

```
go get github.com/tmarcus87/sqlike
```

## Requirements

* Golang 1.13+

## SchemaGenerator

Will be supported in the future.

## Example

```
package main

import (
    _ "github.com/go-sql-driver/mysql"
    "github.com/tmarcus87/sqlike"
)

type Book struct {
    Id   int64  `sqlike:"id"`
    Name string `sqlike:"name"`
}

func main() {
    engine, err :=
        sqlike.NewEngine(
            sqlike.FromHostAndPort("mysql", "master.example.com", 3306, "user", "password", "database"),
            sqlike.WithSlaveByHostAndPort("slave.example.com", 13306, "user", "password"))
    if err != nil {
        panic(err)
    }

    books := make([]*Book, 0)

    if err :=
        engine.
            Auto(context.Background()).
            Select(
                dbschema.Book().Id(),
                dbschema.Book().Name()).
            From(dbschema.Author()).
            InnerJoin(dbschema.Book(), dbschema.Book().AuthorId().EqCol(dbschema.Author().Id())).
            Where(dbschema.Author().Name().Eq("William Shakespeare")).
            OrderBy(sqlike.NewOrder(dbschema.Book().Name(), sqlike.OrderASC))
            LimitAndOffset(5, 0).
            Build().
            FetchInto(&books); err != nil {
        panic(err)
    }

    for i, book := range books {
        fmt.Printf("%4d %s", book.Id, book.Name)
    }
}

```

