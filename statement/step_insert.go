package statement

import (
	"errors"
	"fmt"
	"github.com/tmarcus87/sqlike/model"
	"reflect"
	"strings"
)

var (
	ErrorNoColumnInfo = errors.New("no column info")
)

const (
	StateInsertStmtColumns  = "INSERT_STMT_COLUMNS"
	StateInsertStmtHasValue = "INSERT_STMT_HAS_VALUE"
)

type InsertIntoStep struct {
	parent StatementAcceptor
	table  model.Table
}

func (s *InsertIntoStep) Parent() StatementAcceptor {
	return s.parent
}

func (s *InsertIntoStep) Accept(stmt *StatementImpl) error {
	stmt.Statement += fmt.Sprintf("INSERT INTO `%s` ", s.table.SQLikeTableName())
	return nil
}

type InsertIntoColumnStep struct {
	parent  StatementAcceptor
	columns []model.Column
}

func (s *InsertIntoColumnStep) Parent() StatementAcceptor {
	return s.parent
}

func (s *InsertIntoColumnStep) Accept(stmt *StatementImpl) error {
	cols := make([]string, 0)
	for _, column := range s.columns {
		cols = append(cols, fmt.Sprintf("`%s`", column.SQLikeColumnName()))
	}
	stmt.Statement += fmt.Sprintf("(%s) ", strings.Join(cols, ", "))
	stmt.State[StateInsertStmtColumns] = s.columns
	return nil
}

type InsertIntoValuesStep struct {
	parent StatementAcceptor
	values []interface{}
}

func (s *InsertIntoValuesStep) Parent() StatementAcceptor {
	return s.parent
}

func (s *InsertIntoValuesStep) Accept(stmt *StatementImpl) error {
	if _, ok := stmt.State[StateInsertStmtHasValue]; ok {
		stmt.Statement += ", "
	} else {
		stmt.Statement += "VALUES "
	}

	stmt.Statement += insertValueStatement(len(s.values))
	stmt.Bindings = append(stmt.Bindings, s.values...)
	stmt.State[StateInsertStmtHasValue] = true
	return nil
}

type InsertIntoValueStructStep struct {
	parent StatementAcceptor
	values []interface{}
}

func (s *InsertIntoValueStructStep) Parent() StatementAcceptor {
	return s.parent
}

func (s *InsertIntoValueStructStep) Accept(stmt *StatementImpl) error {
	if len(s.values) == 0 {
		return nil
	}

	if _, ok := stmt.State[StateInsertStmtHasValue]; ok {
		stmt.Statement += ", "
	} else {
		stmt.Statement += "VALUES "
	}

	for _, value := range s.values {
		v := reflect.ValueOf(value)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}

		name2index := make(map[string]int)
		for i := 0; i < v.Type().NumField(); i++ {
			f := v.Type().Field(i)
			name2index[strings.ToLower(f.Name)] = i
			if tag, ok := f.Tag.Lookup("sqlike"); ok {
				name2index[strings.ToLower(tag)] = i
			}
		}

		columnsV, ok := stmt.State[StateInsertStmtColumns]
		if !ok {
			return ErrorNoColumnInfo
		}
		columns, ok := columnsV.([]model.Column)
		if !ok {
			panic("invalid type of column")
		}

		if _, ok := stmt.State[StateInsertStmtHasValue]; ok {
			stmt.Statement += ", "
		}

		stmt.Statement += insertValueStatement(len(columns))
		stmt.State[StateInsertStmtHasValue] = true

		for _, column := range columns {
			fieldIndex, ok := name2index[column.SQLikeColumnName()]
			if !ok {
				return fmt.Errorf("struct field for '%s' is not found", column)
			}
			columnValue := v.Field(fieldIndex).Interface()
			stmt.Bindings = append(stmt.Bindings, columnValue)
		}
	}

	return nil
}

func insertValueStatement(numOfValue int) string {
	vs := make([]string, 0)
	for i := 0; i < numOfValue; i++ {
		vs = append(vs, "?")
	}
	return fmt.Sprintf("(%s)", strings.Join(vs, ", "))
}
