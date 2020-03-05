package sqlike

import (
	"context"
	"database/sql"
	"github.com/tmarcus87/sqlike/session"
)

type Engine interface {
	Auto(ctx context.Context) session.Session
	Master(ctx context.Context) session.Session
	Slave(ctx context.Context) session.Session
	Close() error
}

type basicEngine struct {
	dialect      string
	master       *sql.DB
	slaves       []*sql.DB
	slaveHandler SlaveSelectionHandler
}

func (e *basicEngine) Auto(ctx context.Context) session.Session {
	if isExecuted(ctx) {
		return e.Master(ctx)
	}
	return e.Slave(ctx)
}

func (e *basicEngine) Master(ctx context.Context) session.Session {
	return session.NewSession(ctx, e.master, e.dialect, false)
}

func (e *basicEngine) Slave(ctx context.Context) session.Session {
	var slave *sql.DB
	if len(e.slaves) > 0 {
		slave = e.slaveHandler(e.slaves)
	} else {
		slave = e.master
	}
	return session.NewSession(ctx, slave, e.dialect, true)
}

func (e *basicEngine) Close() error {
	if err := e.master.Close(); err != nil {
		return nil
	}
	for _, slave := range e.slaves {
		if err := slave.Close(); err != nil {
			return nil
		}
	}
	return nil
}
