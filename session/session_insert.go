package session

import (
	"github.com/tmarcus87/sqlike/model"
	"github.com/tmarcus87/sqlike/statement"
)

func (s *basicSession) InsertInto(table model.Table) statement.InsertIntoBranchStep {
	return statement.NewInsertIntoBranchStep(s.NewRootStep(), table)
}
