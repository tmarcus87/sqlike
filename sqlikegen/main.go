package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	dbtype      = flag.String("t", "", "Database type for generating code")
	database    = flag.String("d", "", "Database for generating code")
	username    = flag.String("u", "", "Username for connecting to database")
	password    = flag.String("p", "", "Password for connecting to database")
	hostAndPort = flag.String("h", "", "Host&Port for connecting to database format : 'host:port'")
	outdir      = flag.String("o", "", "Output dir")
	pkg         = flag.String("pkg", "", "Output package")
)

func main() {
	flag.Parse()

	schema, err := Fetch(*dbtype, *username, *password, fmt.Sprintf("tcp(%s)", *hostAndPort), *database)
	if err != nil {
		log.Fatalf("failed to fetch schema : %+v", err)
	}

	if *pkg == "" {
		*pkg = *database
	}

	generators := make([]Generator, 0)
	generators = append(generators, NewConstGenerator(NewWriter(*outdir, "name.go")))
	generators = append(generators, NewSchemaSourceGenerator(NewWriter(*outdir, "schema.go")))
	generators = append(generators, NewValueEntityGenerator(NewWriter(*outdir, "model/value.go")))

	for _, g := range generators {
		if err := g.Generate(*pkg, schema); err != nil {
			log.Fatalf("failed to generate : %+v", err)
		}
	}

}
