package sqlike

import (
	"context"
	"database/sql"
)

type Session interface {
	Explain() ExplainSelectBranchStep
	SelectOne() SelectOneBranchStep
	Select(columns ...Column) SelectColumnBranchStep
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

func (session *basicSession) NewRootStep() *RootStep {
	var (
		q func(context.Context, string, ...interface{}) (*sql.Rows, error)
		e func(context.Context, string, ...interface{}) (sql.Result, error)
	)

	if session.tx != nil {
		q = session.tx.QueryContext
	} else if session.db != nil {
		q = session.db.QueryContext
	} else {
		panic("No available queryer")
	}

	if session.tx != nil {
		e = session.tx.ExecContext
	} else if session.db != nil {
		e = session.db.ExecContext
	} else {
		panic("No available queryer")
	}

	return &RootStep{
		ctx:              session.ctx,
		q:                q,
		e:                e,
		dialectStatement: sqlDialect[session.dialect],
	}
}
