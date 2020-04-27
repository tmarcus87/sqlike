package session

import (
	"context"
	"database/sql"
	"github.com/tmarcus87/sqlike/dialect"
	"github.com/tmarcus87/sqlike/model"
	"github.com/tmarcus87/sqlike/statement"
)

type SQLSession interface {
	Explain() statement.ExplainSelectBranchStep
	Query(stmt string, bindings []interface{}) statement.Statement
	SelectOne() statement.SelectOneBranchStep
	Select(columns ...model.ColumnField) statement.SelectColumnBranchStep
	SelectFrom(table model.Table) statement.SelectFromBranchStep
	InsertInto(table model.Table) statement.InsertIntoBranchStep
	Update(table model.Table) statement.UpdateBranchStep
	DeleteFrom(table model.Table) statement.DeleteFromBranchStep
	Truncate(table model.Table) statement.TruncateBranchStep
}

type Session interface {
	Begin() (TxSession, error)

	SQLSession
}

type basicSession struct {
	db       *sql.DB
	ctx      context.Context
	dialect  string
	readonly bool
}

func NewSession(ctx context.Context, db *sql.DB, dialect string, readonly bool) Session {
	return &basicSession{
		db:       db,
		ctx:      ctx,
		dialect:  dialect,
		readonly: readonly,
	}
}

func (s *basicSession) Begin() (TxSession, error) {
	if s.readonly {
		return nil, ErrorReadonlySession
	}

	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}

	return &basicTxSession{
		tx:      tx,
		ctx:     s.ctx,
		dialect: s.dialect,
	}, nil
}

func (s *basicSession) rootStep() *statement.RootStep {
	return statement.NewRootStep(
		s.ctx,
		dialect.GetDialectStatements(s.dialect),
		s.db.QueryContext,
		s.db.ExecContext)
}

func (s *basicSession) Explain() statement.ExplainSelectBranchStep {
	return statement.NewExplainSelectBranchStep(s.rootStep())
}

func (s *basicSession) Query(stmt string, bindings []interface{}) statement.Statement {
	return statement.NewInstantStep(s.rootStep(), stmt, bindings)
}

func (s *basicSession) SelectOne() statement.SelectOneBranchStep {
	return statement.NewSelectOneBranchStep(s.rootStep())
}

func (s *basicSession) Select(columns ...model.ColumnField) statement.SelectColumnBranchStep {
	return statement.NewSelectColumnBranchStep(s.rootStep(), columns...)
}

func (s *basicSession) SelectFrom(table model.Table) statement.SelectFromBranchStep {
	return statement.NewSelectFromBranchStep(s.rootStep(), table)
}

func (s *basicSession) InsertInto(table model.Table) statement.InsertIntoBranchStep {
	return statement.NewInsertIntoBranchStep(s.rootStep(), table)
}

func (s *basicSession) Update(table model.Table) statement.UpdateBranchStep {
	return statement.NewUpdateBranchStep(s.rootStep(), table)
}

func (s *basicSession) DeleteFrom(table model.Table) statement.DeleteFromBranchStep {
	return statement.NewDeleteFromBranchStep(s.rootStep(), table)
}

func (s *basicSession) Truncate(table model.Table) statement.TruncateBranchStep {
	return statement.NewTruncateBranchStep(s.rootStep(), table)
}
