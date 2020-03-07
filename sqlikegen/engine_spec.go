package main

const (
	pkgDatabaseSql     = "database/sql"
	typeSqlNullBool    = "sql.NullBool"
	typeSqlNullInt32   = "sql.NullInt32"
	typeSqlNullInt64   = "sql.NullInt64"
	typeSqlNullFloat64 = "sql.NullFloat64"
	typeSqlNullString  = "sql.NullString"
	typeSqlNullTime    = "sql.NullTime"

	pkgTime  = "time"
	typeTime = "time.Time"

	typeString = "string"
)

type DataTypeDefinition struct {
	Import         string
	GoType         string
	NullableImport string
	NullableGoType string
}

type EngineDefinition struct {
	// DBType -> GoType
	DataType map[string]DataTypeDefinition
}

var EngineDefs = make(map[string]EngineDefinition)
