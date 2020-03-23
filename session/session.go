package session

import (
	"context"
	"database/sql"
	"github.com/tmarcus87/sqlike/dialect"
	"github.com/tmarcus87/sqlike/model"
	"github.com/tmarcus87/sqlike/statement"
)

type Session interface {
	Begin() error
	Commit() error
	Rollback() error
	Close() error

	Explain() statement.ExplainSelectBranchStep
	SelectOne() statement.SelectOneBranchStep
	Select(columns ...model.ColumnField) statement.SelectColumnBranchStep
	SelectFrom(table model.Table) statement.SelectFromBranchStep
	InsertInto(table model.Table) statement.InsertIntoBranchStep
	Update(table model.Table) statement.UpdateBranchStep
	DeleteFrom(table model.Table) statement.DeleteFromBranchStep
	Truncate(table model.Table) statement.TruncateBranchStep
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

func (s *basicSession) Explain() statement.ExplainSelectBranchStep {
	return statement.NewExplainSelectBranchStep(s.NewRootStep())
}

func (s *basicSession) SelectOne() statement.SelectOneBranchStep {
	return statement.NewSelectOneBranchStep(s.NewRootStep())
}

func (s *basicSession) Select(columns ...model.ColumnField) statement.SelectColumnBranchStep {
	return statement.NewSelectColumnBranchStep(s.NewRootStep(), columns...)
}

func (s *basicSession) SelectFrom(table model.Table) statement.SelectFromBranchStep {
	return statement.NewSelectFromBranchStep(s.NewRootStep(), table)
}

func (s *basicSession) InsertInto(table model.Table) statement.InsertIntoBranchStep {
	return statement.NewInsertIntoBranchStep(s.NewRootStep(), table)
}

func (s *basicSession) Update(table model.Table) statement.UpdateBranchStep {
	return statement.NewUpdateBranchStep(s.NewRootStep(), table)
}

func (s *basicSession) DeleteFrom(table model.Table) statement.DeleteFromBranchStep {
	return statement.NewDeleteFromBranchStep(s.NewRootStep(), table)
}

func (s *basicSession) Truncate(table model.Table) statement.TruncateBranchStep {
	return statement.NewTruncateBranchStep(s.NewRootStep(), table)
}
