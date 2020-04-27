package statement

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/tmarcus87/sqlike/dialect"
)

type Queryer interface {
	DialectStatement(st dialect.StatementType) (string, error)
	Context() context.Context
	Query(string, ...interface{}) (*sql.Rows, error)
	Execute(string, ...interface{}) (sql.Result, error)
}

type RootStep struct {
	ctx              context.Context
	q                func(context.Context, string, ...interface{}) (*sql.Rows, error)
	e                func(context.Context, string, ...interface{}) (sql.Result, error)
	dialectStatement map[dialect.StatementType]string
}

func (s *RootStep) DialectStatement(st dialect.StatementType) (string, error) {
	stmt, ok := s.dialectStatement[st]
	if !ok {
		return "", fmt.Errorf("no dialect statement for %v", st)
	}
	return stmt, nil
}

func (s *RootStep) Parent() StatementAcceptor {
	return nil
}

func (s *RootStep) Accept(*StatementImpl) error { return nil }

func (s *RootStep) Context() context.Context {
	return s.ctx
}

func (s *RootStep) Query(stmt string, args ...interface{}) (*sql.Rows, error) {
	return s.q(s.ctx, stmt, args...)
}

func (s *RootStep) Execute(stmt string, args ...interface{}) (sql.Result, error) {
	return s.e(s.ctx, stmt, args...)
}

func NewRootStep(
	ctx context.Context,
	dialectStatement map[dialect.StatementType]string,
	q func(context.Context, string, ...interface{}) (*sql.Rows, error),
	e func(context.Context, string, ...interface{}) (sql.Result, error)) *RootStep {
	return &RootStep{
		ctx:              ctx,
		q:                q,
		e:                e,
		dialectStatement: dialectStatement,
	}
}

type InstantStep struct {
	parent    StatementAcceptor
	statement string
	bindings  []interface{}
}

func (s *InstantStep) Parent() StatementAcceptor {
	return s.parent
}

func (s *InstantStep) Accept(stmt *StatementImpl) error {
	stmt.Statement += s.statement
	stmt.Bindings = append(stmt.Bindings, s.bindings...)
	return nil
}

func NewInstantStep(parent StatementAcceptor, statement string, bindings []interface{}) Statement {
	return NewStatementBuilder(
		&InstantStep{
			parent:    parent,
			statement: statement,
			bindings:  bindings,
		})
}
