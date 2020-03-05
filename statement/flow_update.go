package statement

import "github.com/tmarcus87/sqlike/model"

type UpdateBranchStep interface {
	Set(column model.Column, value interface{}) UpdateSetBranchStep
	SetRecord(record *model.Record) UpdateSetRecordBranchStep
}

func NewUpdateBranchStep(parent StatementAcceptor, table model.Table) UpdateBranchStep {
	return &updateBranchStepImpl{
		parent: &UpdateStep{
			parent: parent,
			table:  table,
		},
	}
}

type updateBranchStepImpl struct {
	parent StatementAcceptor
}

func (s *updateBranchStepImpl) Parent() StatementAcceptor {
	return s.parent
}

func (s *updateBranchStepImpl) Accept(*StatementImpl) error { return nil }

func (s *updateBranchStepImpl) Set(column model.Column, value interface{}) UpdateSetBranchStep {
	return &updateSetBranchStepImpl{
		parent: &UpdateSetStep{
			parent: s,
			column: column,
			value:  value,
		},
	}
}

func (s *updateBranchStepImpl) SetRecord(record *model.Record) UpdateSetRecordBranchStep {
	return &updateSetRecordBranchStepImpl{
		parent: &UpdateSetRecordStep{
			parent: s,
			record: record,
		},
	}
}

type UpdateSetBranchStep interface {
	Set(column model.Column, value interface{}) UpdateSetBranchStep
	Where(conditions ...model.Condition) UpdateWhereBranchStep
	Build() Statement
}

type updateSetBranchStepImpl struct {
	parent StatementAcceptor
}

func (s *updateSetBranchStepImpl) Parent() StatementAcceptor {
	return s.parent
}

func (s *updateSetBranchStepImpl) Accept(stmt *StatementImpl) error { return nil }

func (s *updateSetBranchStepImpl) Set(column model.Column, value interface{}) UpdateSetBranchStep {
	return &updateSetBranchStepImpl{
		parent: &UpdateSetStep{
			parent: s,
			column: column,
			value:  value,
		},
	}
}

func (s *updateSetBranchStepImpl) Where(conditions ...model.Condition) UpdateWhereBranchStep {
	return &updateWhereBranchStepImpl{
		parent: &WhereStep{
			parent:     s,
			conditions: conditions,
		},
	}
}

func (s *updateSetBranchStepImpl) Build() Statement {
	return NewStatementBuilder(s)
}

type UpdateSetRecordBranchStep interface {
	Where(conditions ...model.Condition) UpdateWhereBranchStep
}

type updateSetRecordBranchStepImpl struct {
	parent StatementAcceptor
}

func (s *updateSetRecordBranchStepImpl) Parent() StatementAcceptor {
	return s.parent
}

func (s *updateSetRecordBranchStepImpl) Accept(stmt *StatementImpl) error { return nil }

func (s *updateSetRecordBranchStepImpl) Where(conditions ...model.Condition) UpdateWhereBranchStep {
	return &updateWhereBranchStepImpl{
		parent: &WhereStep{
			parent:     s,
			conditions: conditions,
		},
	}
}

type UpdateWhereBranchStep interface {
	Build() Statement
}

type updateWhereBranchStepImpl struct {
	parent StatementAcceptor
}

func (s *updateWhereBranchStepImpl) Parent() StatementAcceptor {
	return s.parent
}

func (s *updateWhereBranchStepImpl) Accept(*StatementImpl) error { return nil }

func (s *updateWhereBranchStepImpl) Build() Statement {
	return NewStatementBuilder(s)
}
