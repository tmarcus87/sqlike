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

func (s *basicSession) Commit() error {
	if !s.isBegan {
		logger.Warn("Tx is not began")
		return nil
	}

	if err := s.tx.Commit(); err != nil {
		return err
	}

	s.isBegan = false
	return nil
}

func (s *basicSession) Rollback() error {
	if s.isBegan {
		return s.tx.Rollback()
	}
	return nil
}

func (s *basicSession) Close() error {
	if s.isBegan {
		logger.Debug("Rollback session")
		if err := s.tx.Rollback(); err != nil {
			return err
		}
		s.isBegan = false
	}
	return nil
}
