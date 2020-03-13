package session

import (
	"github.com/tmarcus87/sqlike/model"
	"github.com/tmarcus87/sqlike/statement"
)

func (s *basicSession) Explain() statement.ExplainSelectBranchStep {
	return statement.NewExplainSelectBranchStep(s.NewRootStep())
}

func (s *basicSession) SelectOne() statement.SelectOneBranchStep {
	return statement.NewSelectOneBranchStep(s.NewRootStep())
}

func (s *basicSession) Select(columns ...model.ColumnField) statement.SelectColumnBranchStep {
	return statement.NewSelectColumnBranchStep(s.NewRootStep(), columns...)
}
