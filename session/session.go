package session

import (
	"context"
	"database/sql"
	"github.com/tmarcus87/sqlike/dialect"
	"github.com/tmarcus87/sqlike/model"
	"github.com/tmarcus87/sqlike/statement"
)

type Session interface {
	Explain() statement.ExplainSelectBranchStep
	SelectOne() statement.SelectOneBranchStep
	Select(columns ...model.ColumnField) statement.SelectColumnBranchStep
	InsertInto(table model.Table) statement.InsertIntoBranchStep
	Update(table model.Table) statement.UpdateBranchStep
	DeleteFrom(table model.Table) statement.DeleteFromBranchStep
}

type basicSession struct {
	db      *sql.DB
	ctx     context.Context
	dialect string

	// Query condition
	readonly bool

	// Tx
	tx      *sql.Tx
	isBegan bool
}

func NewSession(ctx context.Context, db *sql.DB, dialect string, readonly bool) Session {
	return &basicSession{
		db:       db,
		ctx:      ctx,
		dialect:  dialect,
		readonly: readonly,
	}
}

func (s *basicSession) NewRootStep() *statement.RootStep {
	var (
		q func(context.Context, string, ...interface{}) (*sql.Rows, error)
		e func(context.Context, string, ...interface{}) (sql.Result, error)
	)

	if s.tx != nil {
		q = s.tx.QueryContext
	} else if s.db != nil {
		q = s.db.QueryContext
	} else {
		panic("No available queryer")
	}

	if s.tx != nil {
		e = s.tx.ExecContext
	} else if s.db != nil {
		e = s.db.ExecContext
	} else {
		panic("No available queryer")
	}

	return statement.NewRootStep(s.ctx, dialect.GetDialectStatements(s.dialect), q, e)
}
