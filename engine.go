package sqlike

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/tmarcus87/sqlike/session"
)

type contextKey int

const (
	txKey contextKey = iota
	commitedKey
)

var (
	ErrCtxCastFail  = errors.New("failed to cast ctx value")
	ErrNoTx         = errors.New("not tx found")
	ErrAlreadyBegan = errors.New("tx is already began")
)

type Engine interface {
	// Create session
	//
	// If the executed flag is set in Context,
	// a session will be created using the connection of the primary DB.
	// If there is no flag, create a session using replica DB connection.
	NewSession(ctx context.Context) session.Session

	// Begin transaction
	//
	// Begin transaction and set to context
	BeginTx(ctx context.Context) (context.Context, error)

	// Rollback transaction
	//
	// Rollback transaction and delete from context
	RollbackTx(ctx context.Context) (context.Context, error)

	// Commit transaction
	//
	// Commit transaction and delete  from context
	CommitTx(ctx context.Context) (context.Context, error)

	// Close transaction
	//
	// Close transaction(uncommited transaction will be rollbacked).
	CloseTx(ctx context.Context) (context.Context, error)

	// Get session from context
	GetSession(ctx context.Context) (sess session.SQLSession, isTx bool, err error)

	// Get transaction session
	GetTxSession(ctx context.Context) (session.TxSession, error)

	// Close engine
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

func (e *basicEngine) BeginTx(ctx context.Context) (context.Context, error) {
	if ctx.Value(txKey) != nil {
		return ctx, ErrAlreadyBegan
	}

	txs, err := e.newMasterSession(ctx).Begin()
	if err != nil {
		return ctx, err
	}

	return MarkAsExecuted(context.WithValue(ctx, txKey, txs)), nil
}

func (e *basicEngine) RollbackTx(ctx context.Context) (context.Context, error) {
	txs, err := e.GetTxSession(ctx)
	if err != nil {
		return ctx, err
	}

	// Rollback
	if err = txs.Rollback(); err != nil {
		return ctx, err
	}

	// Contextを更新
	return context.WithValue(ctx, txKey, nil), nil
}

func (e *basicEngine) CommitTx(ctx context.Context) (context.Context, error) {
	txs, err := e.GetTxSession(ctx)
	if err != nil {
		return ctx, err
	}

	// Commit
	if err = txs.Commit(); err != nil {
		return ctx, err
	}

	// Contextを更新
	return context.WithValue(ctx, txKey, nil), nil
}

func (e *basicEngine) CloseTx(ctx context.Context) (context.Context, error) {
	txs, err := e.GetTxSession(ctx)
	if err != nil {
		return ctx, err
	}

	// Commit
	if err = txs.Close(); err != nil {
		return ctx, err
	}

	// Contextを更新
	return context.WithValue(ctx, txKey, nil), nil
}

func (e *basicEngine) GetSession(ctx context.Context) (session.SQLSession, bool, error) {
	s := ctx.Value(txKey)
	if s == nil {
		return e.NewSession(ctx), false, nil
	}
	txs, ok := s.(session.SQLSession)
	if !ok {
		return nil, false, fmt.Errorf("unexpected context session : %T(%v)", s, s)
	}
	return txs, true, nil
}

func (e *basicEngine) GetTxSession(ctx context.Context) (session.TxSession, error) {
	// Session取得
	s, tx, err := e.GetSession(ctx)
	if err != nil {
		return nil, err
	}
	if !tx {
		return nil, ErrNoTx
	}

	// TxSessionにキャストできなければエラー
	txs, ok := s.(session.TxSession)
	if !ok {
		return nil, fmt.Errorf("unexpected context session : %T(%v)", s, s)
	}
	return txs, nil

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
