package main

func init() {
	EngineDefs["mysql"] =
		EngineDefinition{
			DataType: map[string]DataTypeDefinition{
				// Numeric
				"bit": {
					GoType:         "uint8",
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullInt32,
				},
				"tinyint": {
					GoType:         "int8",
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullInt32,
				},
				"bool": {
					GoType:         "bool",
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullBool,
				},
				"boolean": {
					GoType:         "bool",
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullBool,
				},
				"smallint": {
					GoType:         "int16",
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullInt32,
				},
				"mediumint": {
					GoType:         "int32",
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullInt32,
				},
				"int": {
					GoType:         "int32",
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullInt32,
				},
				"integer": {
					GoType:         "int32",
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullInt32,
				},
				"bigint": {
					GoType:         "int64",
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullInt64,
				},
				"decimal": {
					GoType:         "float32",
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullFloat64,
				},
				"dec": {
					GoType:         "float32",
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullFloat64,
				},
				"float": {
					GoType:         "float32",
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullFloat64,
				},
				"double": {
					GoType:         "float64",
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullFloat64,
				},
				"double precision": {
					GoType:         "float64",
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullFloat64,
				},

				// Date&Time
				"date": {
					GoType:         typeTime,
					Import:         pkgTime,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullTime,
				},
				"datetime": {
					GoType:         typeTime,
					Import:         pkgTime,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullTime,
				},
				"timestamp": {
					GoType:         typeTime,
					Import:         pkgTime,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullTime,
				},
				"time": {
					GoType:         typeTime,
					Import:         pkgTime,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullTime,
				},
				"year": {
					GoType:         typeTime,
					Import:         pkgTime,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullTime,
				},

				// Textual
				"char": {
					GoType:         typeString,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullString,
				},
				"varchar": {
					GoType:         typeString,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullString,
				},
				"binary": {
					GoType:         typeString,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullString,
				},
				"varbinary": {
					GoType:         typeString,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullString,
				},
				"tinyblob": {
					GoType:         typeString,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullString,
				},
				"tinytext": {
					GoType:         typeString,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullString,
				},
				"blob": {
					GoType:         typeString,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullString,
				},
				"text": {
					GoType:         typeString,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullString,
				},
				"mediumblob": {
					GoType:         typeString,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullString,
				},
				"mediumtext": {
					GoType:         typeString,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullString,
				},
				"longblob": {
					GoType:         typeString,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullString,
				},
				"longtext": {
					GoType:         typeString,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullString,
				},
				"enum": {
					GoType:         typeString,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullString,
				},
				"set": {
					GoType:         typeString,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullString,
				},
			},
		}
}
