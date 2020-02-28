package sqlike

import (
	"context"
	"database/sql"
)

type Engine interface {
}

type basicEngine struct {
	dialect string
	master  *sql.DB
	slaves  []*sql.DB
}

func (e *basicEngine) Auto(ctx context.Context) Session {
	if isExecuted(ctx) {
		return e.Master(ctx)
	}
	return e.Slave(ctx)
}

func (e *basicEngine) Master(ctx context.Context) Session {
	return &basicSession{
		db:       e.master,
		ctx:      ctx,
		readonly: false,
	}
}

func (e *basicEngine) Slave(ctx context.Context) Session {
	return &basicSession{
		db:       e.master,
		ctx:      ctx,
		readonly: true,
	}
}
