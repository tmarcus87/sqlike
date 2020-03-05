package statement

import "github.com/tmarcus87/sqlike/model"

type DeleteFromBranchStep interface {
	Where(conditions ...model.Condition) DeleteWhereBranchStep
	Build() Statement
}

func NewDeleteFromBranchStep(parent StatementAcceptor, table model.Table) DeleteFromBranchStep {
	return &deleteFromBranchStepImpl{
		parent: &DeleteFromStep{
			parent: parent,
			table:  table,
		},
	}
}

type deleteFromBranchStepImpl struct {
	parent StatementAcceptor
}

func (s *deleteFromBranchStepImpl) Parent() StatementAcceptor {
	return s.parent
}

func (s *deleteFromBranchStepImpl) Accept(*StatementImpl) error { return nil }

func (s *deleteFromBranchStepImpl) Where(conditions ...model.Condition) DeleteWhereBranchStep {
	return &deleteWhereBranchStepImpl{
		parent: &WhereStep{
			parent:     s,
			conditions: conditions,
		},
	}
}

func (s *deleteFromBranchStepImpl) Build() Statement {
	return NewStatementBuilder(s)
}

type DeleteWhereBranchStep interface {
	Build() Statement
}

type deleteWhereBranchStepImpl struct {
	parent StatementAcceptor
}

func (s *deleteWhereBranchStepImpl) Parent() StatementAcceptor {
	return s.parent
}

func (s *deleteWhereBranchStepImpl) Accept(*StatementImpl) error { return nil }

func (s *deleteWhereBranchStepImpl) Build() Statement {
	return NewStatementBuilder(s)
}
