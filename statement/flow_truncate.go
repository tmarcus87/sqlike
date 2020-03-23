package statement

import "github.com/tmarcus87/sqlike/model"

type TruncateBranchStep interface {
	Build() Statement
}

func NewTruncateBranchStep(parent StatementAcceptor, table model.Table) TruncateBranchStep {
	return &truncateBranchStepImpl{
		parent: &TruncateStep{
			parent: parent,
			table:  table,
		},
	}
}

type truncateBranchStepImpl struct {
	parent StatementAcceptor
}

func (s *truncateBranchStepImpl) Parent() StatementAcceptor {
	return s.parent
}

func (s *truncateBranchStepImpl) Accept(*StatementImpl) error { return nil }

func (s *truncateBranchStepImpl) Build() Statement {
	return NewStatementBuilder(s)
}
