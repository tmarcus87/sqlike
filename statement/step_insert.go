package statement

import (
	"fmt"
	"github.com/tmarcus87/sqlike/model"
	"strings"
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
	stmt.State[StateInsertStmtColumns] = cols
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
	if !strings.Contains(stmt.Statement, "VALUES") {
		stmt.Statement += "VALUES "
	}

	if _, ok := stmt.State[StateInsertStmtHasValue]; ok {
		stmt.Statement += ", "
	}

	vs := make([]string, 0)
	for i := 0; i < len(s.values); i++ {
		vs = append(vs, "?")
	}

	stmt.Statement += fmt.Sprintf("(%s)", strings.Join(vs, ", "))
	stmt.Bindings = append(stmt.Bindings, s.values...)
	stmt.State[StateInsertStmtHasValue] = struct{}{}
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
	panic("implement me")
}
