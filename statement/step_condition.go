package statement

import "github.com/tmarcus87/sqlike/model"

type WhereStep struct {
	parent     StatementAcceptor
	conditions []model.Condition
}

func (s *WhereStep) Parent() StatementAcceptor {
	return s.parent
}

func (s *WhereStep) Accept(stmt *StatementImpl) error {
	if len(s.conditions) == 0 {
		return nil
	}

	stmt.Statement += "WHERE "

	joinCondition(s.conditions, &stmt.Statement, &stmt.Bindings, "AND")

	stmt.Statement += " "

	return nil
}
