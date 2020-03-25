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

	typeBoolColumn    = "BoolColumn"
	typeInt8Column    = "Int8Column"
	typeInt16Column   = "Int16Column"
	typeInt32Column   = "Int32Column"
	typeInt64Column   = "Int64Column"
	typeFloat32Column = "Float32Column"
	typeFloat64Column = "Float64Column"
	typeTextColumn    = "TextColumn"
	typeTimeColumn    = "TimeColumn"
)

type DataTypeDefinition struct {
	Import         string
	GoType         string
	NullableImport string
	NullableGoType string
	BaseStructType string
}

type EngineDefinition struct {
	// DBType -> GoType
	DataType map[string]DataTypeDefinition
}

var EngineDefs = make(map[string]EngineDefinition)
