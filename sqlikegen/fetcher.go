package main

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	QuerySelectTable = `
SELECT
  TABLE_NAME,
  TABLE_TYPE
FROM
  INFORMATION_SCHEMA.TABLES
WHERE
  TABLE_SCHEMA = ?
`
	QuerySelectColumn = `
SELECT
  COLUMN_NAME,
  COLUMN_DEFAULT,
  IS_NULLABLE,
  DATA_TYPE,
  CHARACTER_MAXIMUM_LENGTH,
  CHARACTER_OCTET_LENGTH,
  NUMERIC_PRECISION,
  NUMERIC_SCALE,
  CHARACTER_SET_NAME,
  COLLATION_NAME
FROM
  INFORMATION_SCHEMA.COLUMNS
WHERE
  TABLE_SCHEMA = ?
  AND 
  TABLE_NAME = ?
ORDER BY
  TABLE_SCHEMA ASC,
  TABLE_NAME ASC,
  ORDINAL_POSITION ASC
`
)

func Fetch(driver, username, password, address, database string) (*Schema, error) {
	db, err := sql.Open(driver, fmt.Sprintf("%s:%s@%s/information_schema", username, password, address))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database : %v", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("failed to close db connection : %+v", err)
		}
	}()

	tables, err := fetchTables(db, database)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch table infomation : %v", err)
	}

	defs := make([]Table, 0)
	for _, table := range tables {
		columns, err := fetchColumns(db, database, table)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch column infomation : %v", err)
		}
		table.Columns = columns
		defs = append(defs, table)
	}
	return &Schema{
		DBEngine: driver,
		Database: database,
		Schema:   defs,
	}, nil
}

func fetchTables(db *sql.DB, database string) ([]Table, error) {
	rows, err := db.Query(QuerySelectTable, database)
	if err != nil {
		return nil, fmt.Errorf("failed to query : %w", err)
	}
	defer rows.Close()

	tables := make([]Table, 0)
	for rows.Next() {
		table := Table{DBEngine: *dbtype}
		if err := rows.Scan(&table.Name, &table.Type); err != nil {
			return nil, fmt.Errorf("failed to scan row : %w", err)
		}
		tables = append(tables, table)

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tables, nil
}

func fetchColumns(db *sql.DB, database string, table Table) ([]Column, error) {
	rows, err := db.Query(QuerySelectColumn, database, table.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to query : %w", err)
	}
	defer rows.Close()

	columns := make([]Column, 0)
	for rows.Next() {
		column := Column{DBEngine: *dbtype}
		if err :=
			rows.Scan(
				&column.Name,
				&column.DefaultValue,
				&column.IsNullable,
				&column.DataType,
				&column.CharacterMaximumLength,
				&column.CharacterOctetLength,
				&column.NumericPrecision,
				&column.NumericScale,
				&column.CharacterSetName,
				&column.CollationName); err != nil {
			return nil, fmt.Errorf("failed to scan row : %w", err)
		}
		columns = append(columns, column)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return columns, nil
}
