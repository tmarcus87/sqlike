package session

import (
	"errors"
	"github.com/tmarcus87/sqlike/logger"
)

var (
	ErrorReadonlySession = errors.New("readonly session can execute only SELECT")
)

func (s *basicSession) Begin() (err error) {
	if s.readonly {
		return ErrorReadonlySession
	}

	if s.isBegan {
		logger.Warn("Tx is already began")
		return nil
	}

	s.tx, err = s.db.Begin()
	if err != nil {
		return err
	}
	s.isBegan = true
	return nil
}

func (s *basicSession) Rollback() (err error) {
	if s.isBegan {
		return s.tx.Rollback()
	}
	return nil
}
