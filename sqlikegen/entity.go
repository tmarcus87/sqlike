package main

import (
	"database/sql"
	"fmt"
	"strings"
)

type Schema struct {
	DBEngine string
	Database string
	Schema   map[Table][]Column
}

type Table struct {
	DBEngine string
	Name     string
	Type     string
}

type Column struct {
	DBEngine               string
	Name                   string
	DefaultValue           sql.NullString
	IsNullable             string
	DataType               string
	CharacterMaximumLength sql.NullInt64
	CharacterOctetLength   sql.NullInt64
	NumericPrecision       sql.NullInt64
	NumericScale           sql.NullInt64
	CharacterSetName       sql.NullString
	CollationName          sql.NullString
}

func (c *Column) Nullable() bool {
	return c.IsNullable == "YES"
}

func (c *Column) ValueImport() (string, error) {
	if c.Nullable() {
		return c.NullableImport()
	} else {
		return c.Import()
	}
}

func (c *Column) ValueFieldType() (string, error) {
	if c.Nullable() {
		return c.NullableGoType()
	} else {
		return c.GoType()
	}
}

func (c *Column) dataType() (*DataTypeDefinition, error) {
	defs, ok := EngineDefs[c.DBEngine]
	if !ok {
		return nil, fmt.Errorf("unsupported engine(%s)", c.DBEngine)
	}
	dataType, ok := defs.DataType[strings.ToLower(c.DataType)]
	if !ok {
		return nil, fmt.Errorf("unsupported DataType(%s)", c.DataType)
	}
	return &dataType, nil
}

func (c *Column) GoType() (string, error) {
	dataType, err := c.dataType()
	if err != nil {
		return "", err
	}
	return dataType.GoType, nil
}

func (c *Column) Import() (string, error) {
	dataType, err := c.dataType()
	if err != nil {
		return "", err
	}
	return dataType.Import, nil
}

func (c *Column) NullableGoType() (string, error) {
	dataType, err := c.dataType()
	if err != nil {
		return "", err
	}
	return dataType.NullableGoType, nil
}

func (c *Column) NullableImport() (string, error) {
	dataType, err := c.dataType()
	if err != nil {
		return "", err
	}
	return dataType.NullableImport, nil
}
