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
					BaseStructType: typeInt8Column,
				},
				"tinyint": {
					GoType:         "int8",
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullInt32,
					BaseStructType: typeInt8Column,
				},
				"bool": {
					GoType:         "bool",
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullBool,
					BaseStructType: typeBoolColumn,
				},
				"boolean": {
					GoType:         "bool",
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullBool,
					BaseStructType: typeBoolColumn,
				},
				"smallint": {
					GoType:         "int16",
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullInt32,
					BaseStructType: typeInt16Column,
				},
				"mediumint": {
					GoType:         "int32",
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullInt32,
					BaseStructType: typeInt32Column,
				},
				"int": {
					GoType:         "int32",
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullInt32,
					BaseStructType: typeInt32Column,
				},
				"integer": {
					GoType:         "int32",
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullInt32,
					BaseStructType: typeInt32Column,
				},
				"bigint": {
					GoType:         "int64",
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullInt64,
					BaseStructType: typeInt64Column,
				},
				"decimal": {
					GoType:         "float32",
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullFloat64,
					BaseStructType: typeFloat32Column,
				},
				"dec": {
					GoType:         "float32",
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullFloat64,
					BaseStructType: typeFloat32Column,
				},
				"float": {
					GoType:         "float32",
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullFloat64,
					BaseStructType: typeFloat32Column,
				},
				"double": {
					GoType:         "float64",
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullFloat64,
					BaseStructType: typeFloat64Column,
				},
				"double precision": {
					GoType:         "float64",
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullFloat64,
					BaseStructType: typeFloat64Column,
				},

				// Date&Time
				"date": {
					GoType:         typeTime,
					Import:         pkgTime,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullTime,
					BaseStructType: typeTimeColumn,
				},
				"datetime": {
					GoType:         typeTime,
					Import:         pkgTime,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullTime,
					BaseStructType: typeTimeColumn,
				},
				"timestamp": {
					GoType:         typeTime,
					Import:         pkgTime,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullTime,
					BaseStructType: typeTimeColumn,
				},
				"time": {
					GoType:         typeTime,
					Import:         pkgTime,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullTime,
					BaseStructType: typeTimeColumn,
				},
				"year": {
					GoType:         typeTime,
					Import:         pkgTime,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullTime,
					BaseStructType: typeTimeColumn,
				},

				// Textual
				"char": {
					GoType:         typeString,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullString,
					BaseStructType: typeTextColumn,
				},
				"varchar": {
					GoType:         typeString,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullString,
					BaseStructType: typeTextColumn,
				},
				"binary": {
					GoType:         typeString,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullString,
					BaseStructType: typeTextColumn,
				},
				"varbinary": {
					GoType:         typeString,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullString,
					BaseStructType: typeTextColumn,
				},
				"tinyblob": {
					GoType:         typeString,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullString,
					BaseStructType: typeTextColumn,
				},
				"tinytext": {
					GoType:         typeString,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullString,
					BaseStructType: typeTextColumn,
				},
				"blob": {
					GoType:         typeString,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullString,
					BaseStructType: typeTextColumn,
				},
				"text": {
					GoType:         typeString,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullString,
					BaseStructType: typeTextColumn,
				},
				"mediumblob": {
					GoType:         typeString,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullString,
					BaseStructType: typeTextColumn,
				},
				"mediumtext": {
					GoType:         typeString,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullString,
					BaseStructType: typeTextColumn,
				},
				"longblob": {
					GoType:         typeString,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullString,
					BaseStructType: typeTextColumn,
				},
				"longtext": {
					GoType:         typeString,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullString,
					BaseStructType: typeTextColumn,
				},
				"enum": {
					GoType:         typeString,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullString,
					BaseStructType: typeTextColumn,
				},
				"set": {
					GoType:         typeString,
					NullableImport: pkgDatabaseSql,
					NullableGoType: typeSqlNullString,
					BaseStructType: typeTextColumn,
				},
			},
		}
}
