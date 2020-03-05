package session

import (
	"github.com/tmarcus87/sqlike/model"
	"github.com/tmarcus87/sqlike/statement"
)

func (s *basicSession) Update(table model.Table) statement.UpdateBranchStep {
	return statement.NewUpdateBranchStep(s.NewRootStep(), table)
}
