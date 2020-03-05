package session

import (
	"github.com/tmarcus87/sqlike/model"
	"github.com/tmarcus87/sqlike/statement"
)

func (s *basicSession) DeleteFrom(table model.Table) statement.DeleteFromBranchStep {
	return statement.NewDeleteFromBranchStep(s.NewRootStep(), table)
}
