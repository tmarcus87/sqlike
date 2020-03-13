package statement

import (
	"fmt"
	"github.com/tmarcus87/sqlike/model"
)

type DeleteFromStep struct {
	parent StatementAcceptor
	table  model.Table
}

func (s *DeleteFromStep) Parent() StatementAcceptor {
	return s.parent
}

func (s *DeleteFromStep) Accept(stmt *StatementImpl) error {
	stmt.Statement += fmt.Sprintf("DELETE FROM %s ", s.table.SQLikeTableExpr())
	return nil
}
