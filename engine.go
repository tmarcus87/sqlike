package sqlike

import (
	"context"
	"database/sql"
	"github.com/tmarcus87/sqlike/session"
)

type Engine interface {
	NewSession(ctx context.Context) session.Session
	Close() error
}

type basicEngine struct {
	dialect      string
	master       *sql.DB
	slaves       []*sql.DB
	slaveHandler SlaveSelectionHandler
}

func (e *basicEngine) NewSession(ctx context.Context) session.Session {
	if isExecuted(ctx) {
		return e.newMasterSession(ctx)
	}
	return e.newSlaveSession(ctx)
}

func (e *basicEngine) newMasterSession(ctx context.Context) session.Session {
	return session.NewSession(ctx, e.master, e.dialect, false)
}

func (e *basicEngine) newSlaveSession(ctx context.Context) session.Session {
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
