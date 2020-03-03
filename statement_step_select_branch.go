package sqlike

type ExplainSelectBranchStep interface {
	SelectOne() SelectOneBranchStep
	Select(columns ...Column) SelectColumnBranchStep
}

type ExplainSelectBranchStepImpl struct {
	parent StatementAcceptor
}

func (s *ExplainSelectBranchStepImpl) Parent() StatementAcceptor {
	return s.parent
}

func (s *ExplainSelectBranchStepImpl) Accept(*StatementImpl) {}

func (s *ExplainSelectBranchStepImpl) SelectOne() SelectOneBranchStep {
	return &SelectOneBranchStepImpl{
		parent: &SelectOneStep{
			parent: s,
		},
	}
}

func (s *ExplainSelectBranchStepImpl) Select(columns ...Column) SelectColumnBranchStep {
	return &SelectColumnBranchStepImpl{
		parent: &SelectColumnStep{
			parent:  s,
			columns: columns,
		},
	}
}

type SelectOneBranchStep interface {
	Build() Statement
}

type SelectOneBranchStepImpl struct {
	parent StatementAcceptor
}

func (s *SelectOneBranchStepImpl) Parent() StatementAcceptor {
	return s.parent
}

func (s *SelectOneBranchStepImpl) Accept(*StatementImpl) {}

func (s *SelectOneBranchStepImpl) Build() Statement {
	return buildStatement(s.parent)
}

type SelectColumnBranchStep interface {
	From(table Table) SelectFromBranchStep
}

type SelectColumnBranchStepImpl struct {
	parent *SelectColumnStep
}

func (s *SelectColumnBranchStepImpl) Parent() StatementAcceptor {
	return s.parent
}

func (s *SelectColumnBranchStepImpl) Accept(*StatementImpl) {}

func (s *SelectColumnBranchStepImpl) From(table Table) SelectFromBranchStep {
	return &SelectFromBranchStepImpl{
		parent: &SelectFromStep{
			parent: s,
			table:  table,
		},
	}
}

type SelectFromBranchStep interface {
	Build() Statement
	LeftOuterJoin(table Table, conditions ...Condition) SelectFromJoinBranchStep
	RightOuterJoin(table Table, conditions ...Condition) SelectFromJoinBranchStep
	InnerJoin(table Table, conditions ...Condition) SelectFromJoinBranchStep
	Where(conditions ...Condition) SelectFromWhereBranchStep
	GroupBy(columns ...Column) SelectFromGroupByBranchStep
	OrderBy(orders ...*SortOrder) SelectFromOrderByBranchStep
	LimitAndOffset(limit int32, offset int64) SelectFromLimitAndOffsetBranchStep
}

type SelectFromBranchStepImpl struct {
	parent *SelectFromStep
}

func (s *SelectFromBranchStepImpl) Parent() StatementAcceptor {
	return s.parent
}

func (s *SelectFromBranchStepImpl) Accept(*StatementImpl) {}

func (s *SelectFromBranchStepImpl) Build() Statement {
	return buildStatement(s.parent)
}

func (s *SelectFromBranchStepImpl) LeftOuterJoin(table Table, conditions ...Condition) SelectFromJoinBranchStep {
	return &SelectFromJoinBranchStepImpl{
		parent: &SelectFromJoinStep{
			parent:     s,
			table:      table,
			conditions: conditions,
			joinType:   "LEFT OUTER JOIN",
		},
	}
}

func (s *SelectFromBranchStepImpl) RightOuterJoin(table Table, conditions ...Condition) SelectFromJoinBranchStep {
	return &SelectFromJoinBranchStepImpl{
		parent: &SelectFromJoinStep{
			parent:     s,
			table:      table,
			conditions: conditions,
			joinType:   "RIGHT OUTER JOIN",
		},
	}
}

func (s *SelectFromBranchStepImpl) InnerJoin(table Table, conditions ...Condition) SelectFromJoinBranchStep {
	return &SelectFromJoinBranchStepImpl{
		parent: &SelectFromJoinStep{
			parent:     s,
			table:      table,
			conditions: conditions,
			joinType:   "INNER JOIN",
		},
	}
}

func (s *SelectFromBranchStepImpl) Where(conditions ...Condition) SelectFromWhereBranchStep {
	return &SelectFromWhereBranchStepImpl{
		parent: &WhereStep{
			parent:     s,
			conditions: conditions,
		},
	}
}

func (s *SelectFromBranchStepImpl) GroupBy(columns ...Column) SelectFromGroupByBranchStep {
	return &SelectFromGroupByBranchStepImpl{
		parent: &SelectGroupByStep{
			parent:  s,
			columns: columns,
		},
	}
}

func (s *SelectFromBranchStepImpl) OrderBy(orders ...*SortOrder) SelectFromOrderByBranchStep {
	return &SelectFromOrderByBranchStepImpl{
		parent: &SelectOrderByStep{
			parent: s,
			orders: orders,
		},
	}
}

func (s *SelectFromBranchStepImpl) LimitAndOffset(limit int32, offset int64) SelectFromLimitAndOffsetBranchStep {
	return &SelectFromLimitAndOffsetBranchStepImpl{
		parent: &SelectLimitOffsetStep{
			parent: s,
			limit:  limit,
			offset: offset,
		},
	}
}

type SelectFromJoinBranchStep interface {
	Build() Statement
	Where(conditions ...Condition) SelectFromWhereBranchStep
	GroupBy(columns ...Column) SelectFromGroupByBranchStep
	OrderBy(orders ...*SortOrder) SelectFromOrderByBranchStep
	LimitAndOffset(limit int32, offset int64) SelectFromLimitAndOffsetBranchStep
}

type SelectFromJoinBranchStepImpl struct {
	parent StatementAcceptor
}

func (s *SelectFromJoinBranchStepImpl) Parent() StatementAcceptor {
	return s.parent
}

func (s *SelectFromJoinBranchStepImpl) Accept(*StatementImpl) {}

func (s *SelectFromJoinBranchStepImpl) Build() Statement {
	return buildStatement(s.parent)
}

func (s *SelectFromJoinBranchStepImpl) Where(conditions ...Condition) SelectFromWhereBranchStep {
	return &SelectFromWhereBranchStepImpl{
		parent: &WhereStep{
			parent:     s,
			conditions: conditions,
		},
	}
}

func (s *SelectFromJoinBranchStepImpl) GroupBy(columns ...Column) SelectFromGroupByBranchStep {
	return &SelectFromGroupByBranchStepImpl{
		parent: &SelectGroupByStep{
			parent:  s,
			columns: columns,
		},
	}
}

func (s *SelectFromJoinBranchStepImpl) OrderBy(orders ...*SortOrder) SelectFromOrderByBranchStep {
	return &SelectFromOrderByBranchStepImpl{
		parent: &SelectOrderByStep{
			parent: s,
			orders: orders,
		},
	}
}

func (s *SelectFromJoinBranchStepImpl) LimitAndOffset(limit int32, offset int64) SelectFromLimitAndOffsetBranchStep {
	return &SelectFromLimitAndOffsetBranchStepImpl{
		parent: &SelectLimitOffsetStep{
			parent: s,
			limit:  limit,
			offset: offset,
		},
	}
}

type SelectFromWhereBranchStep interface {
	Build() Statement
	GroupBy(columns ...Column) SelectFromGroupByBranchStep
	OrderBy(orders ...*SortOrder) SelectFromOrderByBranchStep
	LimitAndOffset(limit int32, offset int64) SelectFromLimitAndOffsetBranchStep
}

type SelectFromWhereBranchStepImpl struct {
	parent StatementAcceptor
}

func (s *SelectFromWhereBranchStepImpl) Parent() StatementAcceptor {
	return s.parent
}

func (s *SelectFromWhereBranchStepImpl) Accept(*StatementImpl) {}

func (s *SelectFromWhereBranchStepImpl) Build() Statement {
	return buildStatement(s.parent)
}

func (s *SelectFromWhereBranchStepImpl) GroupBy(columns ...Column) SelectFromGroupByBranchStep {
	return &SelectFromGroupByBranchStepImpl{
		parent: &SelectGroupByStep{
			parent:  s,
			columns: columns,
		},
	}
}

func (s *SelectFromWhereBranchStepImpl) OrderBy(orders ...*SortOrder) SelectFromOrderByBranchStep {
	return &SelectFromOrderByBranchStepImpl{
		parent: &SelectOrderByStep{
			parent: s,
			orders: orders,
		},
	}
}

func (s *SelectFromWhereBranchStepImpl) LimitAndOffset(limit int32, offset int64) SelectFromLimitAndOffsetBranchStep {
	return &SelectFromLimitAndOffsetBranchStepImpl{
		parent: &SelectLimitOffsetStep{
			parent: s,
			limit:  limit,
			offset: offset,
		},
	}
}

type SelectFromGroupByBranchStep interface {
	Build() Statement
	OrderBy(orders ...*SortOrder) SelectFromOrderByBranchStep
	LimitAndOffset(limit int32, offset int64) SelectFromLimitAndOffsetBranchStep
}

type SelectFromGroupByBranchStepImpl struct {
	parent StatementAcceptor
}

func (s *SelectFromGroupByBranchStepImpl) Parent() StatementAcceptor {
	return s.parent
}

func (s *SelectFromGroupByBranchStepImpl) Accept(*StatementImpl) {}

func (s *SelectFromGroupByBranchStepImpl) Build() Statement {
	return buildStatement(s.parent)
}

func (s *SelectFromGroupByBranchStepImpl) OrderBy(orders ...*SortOrder) SelectFromOrderByBranchStep {
	return &SelectFromOrderByBranchStepImpl{
		parent: &SelectOrderByStep{
			parent: s,
			orders: orders,
		},
	}
}

func (s *SelectFromGroupByBranchStepImpl) LimitAndOffset(limit int32, offset int64) SelectFromLimitAndOffsetBranchStep {
	return &SelectFromLimitAndOffsetBranchStepImpl{
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

type SelectFromOrderByBranchStepImpl struct {
	parent StatementAcceptor
}

func (s *SelectFromOrderByBranchStepImpl) Parent() StatementAcceptor {
	return s.parent
}

func (s *SelectFromOrderByBranchStepImpl) Accept(*StatementImpl) {}

func (s *SelectFromOrderByBranchStepImpl) Build() Statement {
	return buildStatement(s.parent)
}

func (s *SelectFromOrderByBranchStepImpl) LimitAndOffset(limit int32, offset int64) SelectFromLimitAndOffsetBranchStep {
	return &SelectFromLimitAndOffsetBranchStepImpl{
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

type SelectFromLimitAndOffsetBranchStepImpl struct {
	parent StatementAcceptor
}

func (s *SelectFromLimitAndOffsetBranchStepImpl) Parent() StatementAcceptor {
	return s.parent
}

func (s *SelectFromLimitAndOffsetBranchStepImpl) Accept(*StatementImpl) {}

func (s *SelectFromLimitAndOffsetBranchStepImpl) Build() Statement {
	return buildStatement(s.parent)
}
