package statement

import (
	"github.com/tmarcus87/sqlike/model"
)

type ExplainSelectBranchStep interface {
	SelectOne() SelectOneBranchStep
	Select(columns ...model.ColumnField) SelectColumnBranchStep
}

func NewExplainSelectBranchStep(parent StatementAcceptor) ExplainSelectBranchStep {
	return &explainSelectBranchStepImpl{
		parent: &SelectExplainStep{
			parent: parent,
		},
	}
}

type explainSelectBranchStepImpl struct {
	parent StatementAcceptor
}

func (s *explainSelectBranchStepImpl) Parent() StatementAcceptor {
	return s.parent
}

func (s *explainSelectBranchStepImpl) Accept(*StatementImpl) error { return nil }

func (s *explainSelectBranchStepImpl) SelectOne() SelectOneBranchStep {
	return &selectOneBranchStepImpl{
		parent: &SelectOneStep{
			parent: s,
		},
	}
}

func (s *explainSelectBranchStepImpl) Select(columns ...model.ColumnField) SelectColumnBranchStep {
	return &selectColumnBranchStepImpl{
		parent: &SelectColumnStep{
			parent:  s,
			columns: columns,
		},
	}
}

type SelectOneBranchStep interface {
	Build() Statement
}

func NewSelectOneBranchStep(parent StatementAcceptor) SelectOneBranchStep {
	return &selectOneBranchStepImpl{
		parent: &SelectOneStep{
			parent: parent,
		},
	}
}

type selectOneBranchStepImpl struct {
	parent StatementAcceptor
}

func (s *selectOneBranchStepImpl) Parent() StatementAcceptor {
	return s.parent
}

func (s *selectOneBranchStepImpl) Accept(*StatementImpl) error { return nil }

func (s *selectOneBranchStepImpl) Build() Statement {
	return NewStatementBuilder(s)
}

type SelectColumnBranchStep interface {
	From(table model.Table) SelectFromBranchStep
}

func NewSelectColumnBranchStep(parent StatementAcceptor, columns ...model.ColumnField) SelectColumnBranchStep {
	return &selectColumnBranchStepImpl{
		parent: &SelectColumnStep{
			parent:  parent,
			columns: columns,
		},
	}
}

type selectColumnBranchStepImpl struct {
	parent *SelectColumnStep
}

func (s *selectColumnBranchStepImpl) Parent() StatementAcceptor {
	return s.parent
}

func (s *selectColumnBranchStepImpl) Accept(*StatementImpl) error { return nil }

func (s *selectColumnBranchStepImpl) From(table model.Table) SelectFromBranchStep {
	return &selectFromBranchStepImpl{
		parent: &SelectFromStep{
			parent: s,
			table:  table,
		},
	}
}

type SelectFromBranchStep interface {
	Build() Statement
	LeftOuterJoin(table model.Table, conditions ...model.Condition) SelectFromJoinBranchStep
	RightOuterJoin(table model.Table, conditions ...model.Condition) SelectFromJoinBranchStep
	InnerJoin(table model.Table, conditions ...model.Condition) SelectFromJoinBranchStep
	Where(conditions ...model.Condition) SelectFromWhereBranchStep
	GroupBy(columns ...model.ColumnField) SelectFromGroupByBranchStep
	OrderBy(orders ...*model.SortOrder) SelectFromOrderByBranchStep
	LimitAndOffset(limit int32, offset int64) SelectFromLimitAndOffsetBranchStep
}

func NewSelectFromBranchStep(parent StatementAcceptor, table model.Table) SelectFromBranchStep {
	return &selectFromBranchStepImpl{
		parent: &SelectFromStep{
			parent: &SelectColumnStep{
				parent: parent,
				columns: []model.ColumnField{
					model.NewAllColumnField(),
				},
			},
			table: table,
		},
	}
}

type selectFromBranchStepImpl struct {
	parent *SelectFromStep
}

func (s *selectFromBranchStepImpl) Parent() StatementAcceptor {
	return s.parent
}

func (s *selectFromBranchStepImpl) Accept(*StatementImpl) error { return nil }

func (s *selectFromBranchStepImpl) Build() Statement {
	return NewStatementBuilder(s)
}

func (s *selectFromBranchStepImpl) LeftOuterJoin(table model.Table, conditions ...model.Condition) SelectFromJoinBranchStep {
	return &selectFromJoinBranchStepImpl{
		parent: &SelectFromJoinStep{
			parent:     s,
			table:      table,
			conditions: conditions,
			joinType:   "LEFT OUTER JOIN",
		},
	}
}

func (s *selectFromBranchStepImpl) RightOuterJoin(table model.Table, conditions ...model.Condition) SelectFromJoinBranchStep {
	return &selectFromJoinBranchStepImpl{
		parent: &SelectFromJoinStep{
			parent:     s,
			table:      table,
			conditions: conditions,
			joinType:   "RIGHT OUTER JOIN",
		},
	}
}

func (s *selectFromBranchStepImpl) InnerJoin(table model.Table, conditions ...model.Condition) SelectFromJoinBranchStep {
	return &selectFromJoinBranchStepImpl{
		parent: &SelectFromJoinStep{
			parent:     s,
			table:      table,
			conditions: conditions,
			joinType:   "INNER JOIN",
		},
	}
}

func (s *selectFromBranchStepImpl) Where(conditions ...model.Condition) SelectFromWhereBranchStep {
	return &selectFromWhereBranchStepImpl{
		parent: &WhereStep{
			parent:     s,
			conditions: conditions,
		},
	}
}

func (s *selectFromBranchStepImpl) GroupBy(columns ...model.ColumnField) SelectFromGroupByBranchStep {
	return &selectFromGroupByBranchStepImpl{
		parent: &SelectGroupByStep{
			parent:  s,
			columns: columns,
		},
	}
}

func (s *selectFromBranchStepImpl) OrderBy(orders ...*model.SortOrder) SelectFromOrderByBranchStep {
	return &selectFromOrderByBranchStepImpl{
		parent: &SelectOrderByStep{
			parent: s,
			orders: orders,
		},
	}
}

func (s *selectFromBranchStepImpl) LimitAndOffset(limit int32, offset int64) SelectFromLimitAndOffsetBranchStep {
	return &selectFromLimitAndOffsetBranchStepImpl{
		parent: &SelectLimitOffsetStep{
			parent: s,
			limit:  limit,
			offset: offset,
		},
	}
}

type SelectFromJoinBranchStep interface {
	Build() Statement
	Where(conditions ...model.Condition) SelectFromWhereBranchStep
	GroupBy(columns ...model.ColumnField) SelectFromGroupByBranchStep
	OrderBy(orders ...*model.SortOrder) SelectFromOrderByBranchStep
	LimitAndOffset(limit int32, offset int64) SelectFromLimitAndOffsetBranchStep
}

type selectFromJoinBranchStepImpl struct {
	parent StatementAcceptor
}

func (s *selectFromJoinBranchStepImpl) Parent() StatementAcceptor {
	return s.parent
}

func (s *selectFromJoinBranchStepImpl) Accept(*StatementImpl) error { return nil }

func (s *selectFromJoinBranchStepImpl) Build() Statement {
	return NewStatementBuilder(s)
}

func (s *selectFromJoinBranchStepImpl) Where(conditions ...model.Condition) SelectFromWhereBranchStep {
	return &selectFromWhereBranchStepImpl{
		parent: &WhereStep{
			parent:     s,
			conditions: conditions,
		},
	}
}

func (s *selectFromJoinBranchStepImpl) GroupBy(columns ...model.ColumnField) SelectFromGroupByBranchStep {
	return &selectFromGroupByBranchStepImpl{
		parent: &SelectGroupByStep{
			parent:  s,
			columns: columns,
		},
	}
}

func (s *selectFromJoinBranchStepImpl) OrderBy(orders ...*model.SortOrder) SelectFromOrderByBranchStep {
	return &selectFromOrderByBranchStepImpl{
		parent: &SelectOrderByStep{
			parent: s,
			orders: orders,
		},
	}
}

func (s *selectFromJoinBranchStepImpl) LimitAndOffset(limit int32, offset int64) SelectFromLimitAndOffsetBranchStep {
	return &selectFromLimitAndOffsetBranchStepImpl{
		parent: &SelectLimitOffsetStep{
			parent: s,
			limit:  limit,
			offset: offset,
		},
	}
}

type SelectFromWhereBranchStep interface {
	Build() Statement
	GroupBy(columns ...model.ColumnField) SelectFromGroupByBranchStep
	OrderBy(orders ...*model.SortOrder) SelectFromOrderByBranchStep
	LimitAndOffset(limit int32, offset int64) SelectFromLimitAndOffsetBranchStep
}

type selectFromWhereBranchStepImpl struct {
	parent StatementAcceptor
}

func (s *selectFromWhereBranchStepImpl) Parent() StatementAcceptor {
	return s.parent
}

func (s *selectFromWhereBranchStepImpl) Accept(*StatementImpl) error { return nil }

func (s *selectFromWhereBranchStepImpl) Build() Statement {
	return NewStatementBuilder(s)
}

func (s *selectFromWhereBranchStepImpl) GroupBy(columns ...model.ColumnField) SelectFromGroupByBranchStep {
	return &selectFromGroupByBranchStepImpl{
		parent: &SelectGroupByStep{
			parent:  s,
			columns: columns,
		},
	}
}

func (s *selectFromWhereBranchStepImpl) OrderBy(orders ...*model.SortOrder) SelectFromOrderByBranchStep {
	return &selectFromOrderByBranchStepImpl{
		parent: &SelectOrderByStep{
			parent: s,
			orders: orders,
		},
	}
}

func (s *selectFromWhereBranchStepImpl) LimitAndOffset(limit int32, offset int64) SelectFromLimitAndOffsetBranchStep {
	return &selectFromLimitAndOffsetBranchStepImpl{
		parent: &SelectLimitOffsetStep{
			parent: s,
			limit:  limit,
			offset: offset,
		},
	}
}

type SelectFromGroupByBranchStep interface {
	Build() Statement
	OrderBy(orders ...*model.SortOrder) SelectFromOrderByBranchStep
	LimitAndOffset(limit int32, offset int64) SelectFromLimitAndOffsetBranchStep
}

type selectFromGroupByBranchStepImpl struct {
	parent StatementAcceptor
}

func (s *selectFromGroupByBranchStepImpl) Parent() StatementAcceptor {
	return s.parent
}

func (s *selectFromGroupByBranchStepImpl) Accept(*StatementImpl) error { return nil }

func (s *selectFromGroupByBranchStepImpl) Build() Statement {
	return NewStatementBuilder(s)
}

func (s *selectFromGroupByBranchStepImpl) OrderBy(orders ...*model.SortOrder) SelectFromOrderByBranchStep {
	return &selectFromOrderByBranchStepImpl{
		parent: &SelectOrderByStep{
			parent: s,
			orders: orders,
		},
	}
}

func (s *selectFromGroupByBranchStepImpl) LimitAndOffset(limit int32, offset int64) SelectFromLimitAndOffsetBranchStep {
	return &selectFromLimitAndOffsetBranchStepImpl{
		parent: &SelectLimitOffsetStep{
			parent: s,
			limit:  limit,
			offset: offset,
		},
	}
}

type SelectFromOrderByBranchStep interface {
	Build() Statement
	LimitAndOffset(limit int32, offset int64) SelectFromLimitAndOffsetBranchStep
}

type selectFromOrderByBranchStepImpl struct {
	parent StatementAcceptor
}

func (s *selectFromOrderByBranchStepImpl) Parent() StatementAcceptor {
	return s.parent
}

func (s *selectFromOrderByBranchStepImpl) Accept(*StatementImpl) error { return nil }

func (s *selectFromOrderByBranchStepImpl) Build() Statement {
	return NewStatementBuilder(s.parent)
}

func (s *selectFromOrderByBranchStepImpl) LimitAndOffset(limit int32, offset int64) SelectFromLimitAndOffsetBranchStep {
	return &selectFromLimitAndOffsetBranchStepImpl{
		parent: &SelectLimitOffsetStep{
			parent: s,
			limit:  limit,
			offset: offset,
		},
	}
}

type SelectFromLimitAndOffsetBranchStep interface {
	Build() Statement
}

type selectFromLimitAndOffsetBranchStepImpl struct {
	parent StatementAcceptor
}

func (s *selectFromLimitAndOffsetBranchStepImpl) Parent() StatementAcceptor {
	return s.parent
}

func (s *selectFromLimitAndOffsetBranchStepImpl) Accept(*StatementImpl) error { return nil }

func (s *selectFromLimitAndOffsetBranchStepImpl) Build() Statement {
	return NewStatementBuilder(s.parent)
}
