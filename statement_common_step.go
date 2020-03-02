package sqlike

import (
	"context"
	"database/sql"
)

type Queryer interface {
	DialectStatement(st StatementType) string
	Context() context.Context
	Query(string, ...interface{}) (*sql.Rows, error)
	Execute(string, ...interface{}) (sql.Result, error)
}

type RootStep struct {
	ctx              context.Context
	q                func(context.Context, string, ...interface{}) (*sql.Rows, error)
	e                func(context.Context, string, ...interface{}) (sql.Result, error)
	dialectStatement map[StatementType]string
}

func (s *RootStep) DialectStatement(st StatementType) string {
	return s.dialectStatement[st]
}

func (s *RootStep) Parent() StatementAcceptor {
	return nil
}

func (s *RootStep) Accept(*StatementImpl) {
	/* Do nothing */
}

func (s *RootStep) Context() context.Context {
	return s.ctx
}

func (s *RootStep) Query(stmt string, args ...interface{}) (*sql.Rows, error) {
	return s.q(s.ctx, stmt, args...)
}

func (s *RootStep) Execute(stmt string, args ...interface{}) (sql.Result, error) {
	return s.e(s.ctx, stmt, args...)
}

type InstantStep struct {
	parent    StatementAcceptor
	statement string
	bindings  []interface{}
}

func (s *InstantStep) Parent() StatementAcceptor {
	return s.parent
}

func (s *InstantStep) Accept(stmt *StatementImpl) {
	stmt.Statement += s.statement
	stmt.Bindings = append(stmt.Bindings, s.bindings)
}
