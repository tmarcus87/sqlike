package sqlike

type WhereStep struct {
	parent     StatementAcceptor
	conditions []Condition
}

func (s *WhereStep) Parent() StatementAcceptor {
	return s.parent
}

func (s *WhereStep) Accept(stmt *StatementImpl) {
	if len(s.conditions) == 0 {
		return
	}

	stmt.Statement += "WHERE "

	joinCondition(s.conditions, &stmt.Statement, &stmt.Bindings, "AND")
}
