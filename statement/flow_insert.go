package statement

import "github.com/tmarcus87/sqlike/model"

type InsertIntoBranchStep interface {
	Columns(cols ...model.ColumnField) InsertIntoColumnBranchStep
	Values(values ...interface{}) InsertIntoValuesBranchStep
	ValueStructs(values ...interface{}) InsertIntoValueStructsBranchStep
	Select(columns ...model.ColumnField) SelectColumnBranchStep
}

func NewInsertIntoBranchStep(parent StatementAcceptor, table model.Table) InsertIntoBranchStep {
	return &insertIntoBranchStepImpl{
		parent: &InsertIntoStep{
			parent: parent,
			table:  table,
		},
	}
}

type insertIntoBranchStepImpl struct {
	parent StatementAcceptor
}

func (s *insertIntoBranchStepImpl) Parent() StatementAcceptor {
	return s.parent
}

func (s *insertIntoBranchStepImpl) Accept(*StatementImpl) error { return nil }

func (s *insertIntoBranchStepImpl) Columns(columns ...model.ColumnField) InsertIntoColumnBranchStep {
	return &insertIntoColumnBranchStepImpl{
		parent: &InsertIntoColumnStep{
			parent:  s,
			columns: columns,
		},
	}
}

func (s *insertIntoBranchStepImpl) Values(values ...interface{}) InsertIntoValuesBranchStep {
	return &insertIntoValuesBranchStepImpl{
		parent: &InsertIntoValuesStep{
			parent: s,
			values: values,
		},
	}
}

func (s *insertIntoBranchStepImpl) ValueStructs(values ...interface{}) InsertIntoValueStructsBranchStep {
	return &insertIntoValueStructsBranchStepImpl{
		parent: &InsertIntoValueStructStep{
			parent: s,
			values: values,
		},
	}
}

func (s *insertIntoBranchStepImpl) Select(cols ...model.ColumnField) SelectColumnBranchStep {
	return NewSelectColumnBranchStep(s.parent, cols...)
}

type InsertIntoColumnBranchStep interface {
	Values(values ...interface{}) InsertIntoValuesBranchStep
	ValueStructs(values ...interface{}) InsertIntoValueStructsBranchStep
	Select(columns ...model.ColumnField) SelectColumnBranchStep
}

type insertIntoColumnBranchStepImpl struct {
	parent StatementAcceptor
}

func (s *insertIntoColumnBranchStepImpl) Parent() StatementAcceptor {
	return s.parent
}

func (s *insertIntoColumnBranchStepImpl) Accept(*StatementImpl) error { return nil }

func (s *insertIntoColumnBranchStepImpl) Values(values ...interface{}) InsertIntoValuesBranchStep {
	return &insertIntoValuesBranchStepImpl{
		parent: &InsertIntoValuesStep{
			parent: s,
			values: values,
		},
	}
}

func (s *insertIntoColumnBranchStepImpl) ValueStructs(values ...interface{}) InsertIntoValueStructsBranchStep {
	return &insertIntoValueStructsBranchStepImpl{
		parent: &InsertIntoValueStructStep{
			parent: s,
			values: values,
		},
	}
}

func (s *insertIntoColumnBranchStepImpl) Select(columns ...model.ColumnField) SelectColumnBranchStep {
	return NewSelectColumnBranchStep(s.parent, columns...)
}

type InsertIntoValuesBranchStep interface {
	Build() Statement
	Values(values ...interface{}) InsertIntoValuesBranchStep
}

type insertIntoValuesBranchStepImpl struct {
	parent StatementAcceptor
}

func (s *insertIntoValuesBranchStepImpl) Parent() StatementAcceptor {
	return s.parent
}

func (s *insertIntoValuesBranchStepImpl) Accept(*StatementImpl) error { return nil }

func (s *insertIntoValuesBranchStepImpl) Build() Statement {
	return NewStatementBuilder(s)
}

func (s *insertIntoValuesBranchStepImpl) Values(values ...interface{}) InsertIntoValuesBranchStep {
	return &insertIntoValuesBranchStepImpl{
		parent: &InsertIntoValuesStep{
			parent: s,
			values: values,
		},
	}
}

type InsertIntoValueStructsBranchStep interface {
	Build() Statement
}

type insertIntoValueStructsBranchStepImpl struct {
	parent StatementAcceptor
}

func (s *insertIntoValueStructsBranchStepImpl) Parent() StatementAcceptor {
	return s.parent
}

func (s *insertIntoValueStructsBranchStepImpl) Accept(*StatementImpl) error { return nil }

func (s *insertIntoValueStructsBranchStepImpl) Build() Statement {
	return NewStatementBuilder(s)
}
