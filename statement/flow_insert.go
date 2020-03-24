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
	OnDuplicateKeyIgnore() InsertOnDuplicateKeyIgnoreBranchStep
	OnDuplicateKeyUpdate() InsertOnDuplicateKeyUpdateBranchStep
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

func (s *insertIntoValuesBranchStepImpl) OnDuplicateKeyIgnore() InsertOnDuplicateKeyIgnoreBranchStep {
	return &insertOnDuplicateKeyIgnoreBranchStepImpl{
		parent: &InsertOnDuplicateKeyIgnoreStep{
			parent: s,
		},
	}
}

func (s *insertIntoValuesBranchStepImpl) OnDuplicateKeyUpdate() InsertOnDuplicateKeyUpdateBranchStep {
	return &insertOnDuplicateKeyUpdateBranchStepImpl{
		parent: &InsertOnDuplicateKeyUpdateStep{
			parent: s,
		},
	}
}

type InsertIntoValueStructsBranchStep interface {
	Build() Statement
	OnDuplicateKeyIgnore() InsertOnDuplicateKeyIgnoreBranchStep
	OnDuplicateKeyUpdate() InsertOnDuplicateKeyUpdateBranchStep
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

func (s *insertIntoValueStructsBranchStepImpl) OnDuplicateKeyIgnore() InsertOnDuplicateKeyIgnoreBranchStep {
	return &insertOnDuplicateKeyIgnoreBranchStepImpl{
		parent: &InsertOnDuplicateKeyIgnoreStep{
			parent: s,
		},
	}
}

func (s *insertIntoValueStructsBranchStepImpl) OnDuplicateKeyUpdate() InsertOnDuplicateKeyUpdateBranchStep {
	return &insertOnDuplicateKeyUpdateBranchStepImpl{
		parent: &InsertOnDuplicateKeyUpdateStep{
			parent: s,
		},
	}
}

type InsertOnDuplicateKeyIgnoreBranchStep interface {
	Build() Statement
}

type insertOnDuplicateKeyIgnoreBranchStepImpl struct {
	parent StatementAcceptor
}

func (s *insertOnDuplicateKeyIgnoreBranchStepImpl) Parent() StatementAcceptor {
	return s.parent
}

func (s insertOnDuplicateKeyIgnoreBranchStepImpl) Accept(*StatementImpl) error { return nil }

func (s *insertOnDuplicateKeyIgnoreBranchStepImpl) Build() Statement {
	return NewStatementBuilder(s)
}

type InsertOnDuplicateKeyUpdateBranchStep interface {
	Set(column model.Column, value interface{}) InsertOnDuplicateKeyUpdateSetBranchStep
	SetRecord(record *model.Record) InsertOnDuplicateKeyUpdateSetRecordBranchStep
}

type insertOnDuplicateKeyUpdateBranchStepImpl struct {
	parent StatementAcceptor
}

func (s *insertOnDuplicateKeyUpdateBranchStepImpl) Parent() StatementAcceptor {
	return s.parent
}

func (s *insertOnDuplicateKeyUpdateBranchStepImpl) Accept(*StatementImpl) error { return nil }

func (s *insertOnDuplicateKeyUpdateBranchStepImpl) Set(column model.Column, value interface{}) InsertOnDuplicateKeyUpdateSetBranchStep {
	return &insertOnDuplicateKeyUpdateSetBranchStepImpl{
		parent: &InsertOnDuplicateKeyUpdateSetStep{
			parent: s,
			column: column,
			value:  value,
		},
	}
}

func (s *insertOnDuplicateKeyUpdateBranchStepImpl) SetRecord(record *model.Record) InsertOnDuplicateKeyUpdateSetRecordBranchStep {
	return &insertOnDuplicateKeyUpdateSetRecordBranchStepImpl{
		parent: &InsertOnDuplicateKeyUpdateSetRecordStep{
			parent: s,
			record: record,
		},
	}
}

type InsertOnDuplicateKeyUpdateSetBranchStep interface {
	Build() Statement
	Set(column model.Column, value interface{}) InsertOnDuplicateKeyUpdateSetBranchStep
}

type insertOnDuplicateKeyUpdateSetBranchStepImpl struct {
	parent StatementAcceptor
}

func (s *insertOnDuplicateKeyUpdateSetBranchStepImpl) Parent() StatementAcceptor {
	return s.parent
}

func (s *insertOnDuplicateKeyUpdateSetBranchStepImpl) Accept(*StatementImpl) error { return nil }

func (s *insertOnDuplicateKeyUpdateSetBranchStepImpl) Build() Statement {
	return NewStatementBuilder(s)
}

func (s *insertOnDuplicateKeyUpdateSetBranchStepImpl) Set(column model.Column, value interface{}) InsertOnDuplicateKeyUpdateSetBranchStep {
	return &insertOnDuplicateKeyUpdateSetBranchStepImpl{
		parent: &InsertOnDuplicateKeyUpdateSetStep{
			parent: s,
			column: column,
			value:  value,
		},
	}
}

type InsertOnDuplicateKeyUpdateSetRecordBranchStep interface {
	Build() Statement
}

type insertOnDuplicateKeyUpdateSetRecordBranchStepImpl struct {
	parent StatementAcceptor
}

func (s *insertOnDuplicateKeyUpdateSetRecordBranchStepImpl) Parent() StatementAcceptor {
	return s.parent
}

func (s *insertOnDuplicateKeyUpdateSetRecordBranchStepImpl) Accept(*StatementImpl) error { return nil }

func (s *insertOnDuplicateKeyUpdateSetRecordBranchStepImpl) Build() Statement {
	return NewStatementBuilder(s)
}
