package session

import (
	"context"
	"database/sql"
	"errors"
	"github.com/tmarcus87/sqlike/dialect"
	"github.com/tmarcus87/sqlike/model"
	"github.com/tmarcus87/sqlike/statement"
)

var (
	ErrorReadonlySession = errors.New("readonly session can execute only SELECT")
)

type TxSession interface {
	SQLSession

	IsFlushed() bool
	Commit() error
	Rollback() error
	Close() error
}

type basicTxSession struct {
	tx      *sql.Tx
	ctx     context.Context
	dialect string

	flushed bool
}

func (s *basicTxSession) rootStep() *statement.RootStep {
	return statement.NewRootStep(
		s.ctx,
		dialect.GetDialectStatements(s.dialect),
		s.tx.QueryContext,
		s.tx.ExecContext)
}

func (s *basicTxSession) Explain() statement.ExplainSelectBranchStep {
	return statement.NewExplainSelectBranchStep(s.rootStep())
}

func (s *basicTxSession) Query(stmt string, bindings []interface{}) statement.Statement {
	return statement.NewInstantStep(s.rootStep(), stmt, bindings)
}

func (s *basicTxSession) SelectOne() statement.SelectOneBranchStep {
	return statement.NewSelectOneBranchStep(s.rootStep())
}

func (s *basicTxSession) Select(columns ...model.ColumnField) statement.SelectColumnBranchStep {
	return statement.NewSelectColumnBranchStep(s.rootStep(), columns...)
}

func (s *basicTxSession) SelectFrom(table model.Table) statement.SelectFromBranchStep {
	return statement.NewSelectFromBranchStep(s.rootStep(), table)
}

func (s *basicTxSession) InsertInto(table model.Table) statement.InsertIntoBranchStep {
	return statement.NewInsertIntoBranchStep(s.rootStep(), table)
}

func (s *basicTxSession) Update(table model.Table) statement.UpdateBranchStep {
	return statement.NewUpdateBranchStep(s.rootStep(), table)
}

func (s *basicTxSession) DeleteFrom(table model.Table) statement.DeleteFromBranchStep {
	return statement.NewDeleteFromBranchStep(s.rootStep(), table)
}

func (s *basicTxSession) Truncate(table model.Table) statement.TruncateBranchStep {
	return statement.NewTruncateBranchStep(s.rootStep(), table)
}

func (s *basicTxSession) Commit() (err error) {
	err = s.tx.Commit()
	if err == nil {
		s.flushed = true
	}
	return
}

func (s *basicTxSession) Rollback() (err error) {
	err = s.tx.Rollback()
	if err == nil {
		s.flushed = true
	}
	return
}

func (s *basicTxSession) Close() error {
	if !s.flushed {
		return s.tx.Rollback()
	}
	return nil
}

func (s *basicTxSession) IsFlushed() bool {
	return s.flushed
}
