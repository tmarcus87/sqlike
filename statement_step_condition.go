package sqlike

type WhereStep struct {
	parent     StatementAcceptor
	conditions []Condition
}

func (s *WhereStep) DialectStatement(st StatementType) string {
	return s.parent.DialectStatement(st)
}

func (s *WhereStep) Parent() StatementAcceptor {
	return s.parent
}

func (s *WhereStep) Accept(stmt *StatementImpl) {
	if len(s.conditions) == 0 {
		return
	}
	stmt.Statement += "WHERE "

	for _, condition := range s.conditions {
		condition.Apply(&stmt.Statement, &stmt.Bindings)
	}
}
