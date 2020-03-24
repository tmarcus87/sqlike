package statement

import (
	"errors"
	"fmt"
	"github.com/tmarcus87/sqlike/dialect"
	"github.com/tmarcus87/sqlike/model"
	"strings"
)

var (
	ErrorNoColumnInfo = errors.New("no column info")
)

const (
	StateInsertStmtColumns                 = "INSERT_STMT_COLUMNS"
	StateInsertStmtHasValue                = "INSERT_STMT_HAS_VALUE"
	StateInsertOnDuplicateKeyUpdateStmtSet = "INSERT_STMT_DUPLICATE_UPDATE_SET"
)

type InsertIntoStep struct {
	parent StatementAcceptor
	table  model.Table
}

func (s *InsertIntoStep) Parent() StatementAcceptor {
	return s.parent
}

func (s *InsertIntoStep) Accept(stmt *StatementImpl) error {
	stmt.Statement += fmt.Sprintf("INSERT INTO %s ", s.table.SQLikeTableExpr())
	return nil
}

type InsertIntoColumnStep struct {
	parent  StatementAcceptor
	columns []model.ColumnField
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
		stmt.Statement = strings.TrimSuffix(stmt.Statement, " ")
		stmt.Statement += ", "
	} else {
		stmt.Statement += "VALUES "
	}

	stmt.Statement += insertValueStatement(len(s.values)) + " "
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
		stmt.Statement = strings.TrimSuffix(stmt.Statement, " ")
		stmt.Statement += ", "
	} else {
		stmt.Statement += "VALUES "
	}

	for _, value := range s.values {
		fvm, err := getColumnName2FieldValueMap(value)
		if err != nil {
			return err
		}

		columnsV, ok := stmt.State[StateInsertStmtColumns]
		if !ok {
			return ErrorNoColumnInfo
		}
		columns, ok := columnsV.([]model.ColumnField)
		if !ok {
			panic("invalid type of column")
		}

		if _, ok := stmt.State[StateInsertStmtHasValue]; ok {
			stmt.Statement += ", "
		}

		stmt.Statement += insertValueStatement(len(columns))
		stmt.State[StateInsertStmtHasValue] = true

		for _, column := range columns {
			fv, ok := fvm[column.SQLikeColumnName()]
			if !ok {
				return fmt.Errorf("struct field for '%s' is not found", column)
			}
			stmt.Bindings = append(stmt.Bindings, fv.Interface())
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

type InsertOnDuplicateKeyIgnoreStep struct {
	parent StatementAcceptor
}

func (s *InsertOnDuplicateKeyIgnoreStep) Parent() StatementAcceptor {
	return s.parent
}

func (s *InsertOnDuplicateKeyIgnoreStep) Accept(stmt *StatementImpl) error {
	q, err := getQueryer(s)
	if err != nil {
		return err
	}

	st, err := q.DialectStatement(dialect.StatementTypeOnDuplicateKeyIgnore)
	if err != nil {
		return err
	}
	stmt.Statement += st + " "
	return nil
}

type InsertOnDuplicateKeyUpdateStep struct {
	parent StatementAcceptor
}

func (s *InsertOnDuplicateKeyUpdateStep) Parent() StatementAcceptor {
	return s.parent
}

func (s *InsertOnDuplicateKeyUpdateStep) Accept(stmt *StatementImpl) error {
	q, err := getQueryer(s)
	if err != nil {
		return err
	}

	st, err := q.DialectStatement(dialect.StatementTypeOnDuplicateKeyUpdate)
	if err != nil {
		return err
	}
	stmt.Statement += st + " "
	return nil
}

type InsertOnDuplicateKeyUpdateSetStep struct {
	parent StatementAcceptor
	column model.Column
	value  interface{}
}

func (s *InsertOnDuplicateKeyUpdateSetStep) Parent() StatementAcceptor {
	return s.parent
}

func (s *InsertOnDuplicateKeyUpdateSetStep) Accept(stmt *StatementImpl) error {
	if _, ok := stmt.State[StateInsertOnDuplicateKeyUpdateStmtSet]; ok {
		stmt.Statement = strings.TrimSuffix(stmt.Statement, " ")
		stmt.Statement += ", "
	} else {
		stmt.Statement += "SET "
	}
	stmt.Statement += fmt.Sprintf("`%s` = ?", s.column.SQLikeColumnName())
	stmt.Bindings = append(stmt.Bindings, s.value)
	stmt.State[StateUpdateStmtSet] = true
	return nil
}

type InsertOnDuplicateKeyUpdateSetRecordStep struct {
	parent StatementAcceptor
	record *model.Record
}

func (s *InsertOnDuplicateKeyUpdateSetRecordStep) Parent() StatementAcceptor {
	return s.parent
}

func (s *InsertOnDuplicateKeyUpdateSetRecordStep) Accept(stmt *StatementImpl) error {
	return ApplySetStatement(stmt, s.record)
}
