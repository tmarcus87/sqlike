package statement

import (
	"fmt"
	"github.com/tmarcus87/sqlike/model"
)

type TruncateStep struct {
	parent StatementAcceptor
	table  model.Table
}

func (s *TruncateStep) Parent() StatementAcceptor {
	return s.parent
}

func (s *TruncateStep) Accept(stmt *StatementImpl) error {
	stmt.Statement += fmt.Sprintf("TRUNCATE %s", s.table.SQLikeTableExpr())
	return nil
}
