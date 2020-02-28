package sqlike

import "github.com/tmarcus87/sqlike/logger"

func (session *basicSession) Begin() (err error) {
	if session.readonly {
		return ErrorReadonlySession
	}

	if session.isBegan {
		logger.Warn("Tx is already began")
		return nil
	}

	session.tx, err = session.db.Begin()
	if err != nil {
		return err
	}
	session.isBegan = true
	return nil
}

func (session *basicSession) Rollback() (err error) {
	if session.isBegan {
		return session.tx.Rollback()
	}
	return nil
}
