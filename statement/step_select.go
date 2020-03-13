package statement

import (
	"fmt"
	"github.com/tmarcus87/sqlike/dialect"
	"github.com/tmarcus87/sqlike/model"
	"strings"
)

type SelectExplainStep struct {
	parent StatementAcceptor
}

func (s *SelectExplainStep) Parent() StatementAcceptor {
	return s.parent
}

func (s *SelectExplainStep) Accept(stmt *StatementImpl) error {
	stmt.Statement += "EXPLAIN "
	return nil
}

type SelectOneStep struct {
	parent StatementAcceptor
}

func (s *SelectOneStep) Parent() StatementAcceptor {
	return s.parent
}

func (s *SelectOneStep) Accept(stmt *StatementImpl) error {
	q, err := getQueryer(s.parent)
	if err != nil {
		return err
	}

	st, err := q.DialectStatement(dialect.StatementTypeSelectOne)
	if err != nil {
		return err
	}
	stmt.Statement += st
	return nil
}

type SelectColumnStep struct {
	parent  StatementAcceptor
	columns []model.ColumnField
}

func (s *SelectColumnStep) Parent() StatementAcceptor {
	return s.parent
}

func (s *SelectColumnStep) Accept(stmt *StatementImpl) error {
	cols := make([]string, 0)
	for _, column := range s.columns {
		s := column
		cols = append(cols, fmt.Sprintf("%s", s.SQLikeFieldExpr()))
	}
	stmt.Statement += fmt.Sprintf("SELECT %s ", strings.Join(cols, ", "))
	return nil
}

type SelectFromStep struct {
	parent StatementAcceptor
	table  model.Table
}

func (s *SelectFromStep) Parent() StatementAcceptor {
	return s.parent
}

func (s *SelectFromStep) Accept(stmt *StatementImpl) error {
	stmt.Statement += fmt.Sprintf("FROM %s ", s.table.SQLikeTableExpr())
	return nil
}

type SelectFromJoinStep struct {
	parent     StatementAcceptor
	table      model.Table
	conditions []model.Condition
	joinType   string
}

func (s *SelectFromJoinStep) Parent() StatementAcceptor {
	return s.parent
}

func (s *SelectFromJoinStep) Accept(stmt *StatementImpl) error {
	var onStmt string
	joinCondition(s.conditions, &onStmt, &stmt.Bindings, "AND")

	stmt.Statement += fmt.Sprintf("%s %s ON %s ", s.joinType, s.table.SQLikeTableExpr(), onStmt)
	return nil
}

type SelectGroupByStep struct {
	parent  StatementAcceptor
	columns []model.ColumnField
}

func (s *SelectGroupByStep) Parent() StatementAcceptor {
	return s.parent
}

func (s *SelectGroupByStep) Accept(stmt *StatementImpl) error {
	if len(s.columns) == 0 {
		return nil
	}

	cols := make([]string, 0)
	for _, column := range s.columns {
		cols =
			append(cols,
				fmt.Sprintf("`%s`.`%s`",
					column.SQLikeTable().SQLikeAliasOrName(),
					column.SQLikeColumnName()))
	}

	stmt.Statement += fmt.Sprintf("GROUP BY %s ", strings.Join(cols, ", "))
	return nil
}

type SelectOrderByStep struct {
	parent StatementAcceptor
	orders []*model.SortOrder
}

func (s *SelectOrderByStep) Parent() StatementAcceptor {
	return s.parent
}

func (s *SelectOrderByStep) Accept(stmt *StatementImpl) error {
	if len(s.orders) == 0 {
		return nil
	}

	orders := make([]string, 0)
	for _, order := range s.orders {
		orders =
			append(orders,
				fmt.Sprintf("`%s`.`%s` %s",
					order.Column.SQLikeTable().SQLikeAliasOrName(),
					order.Column.SQLikeAliasOrName(),
					order.Order))
	}

	stmt.Statement += fmt.Sprintf("ORDER BY %s ", strings.Join(orders, ", "))
	return nil
}

type SelectLimitOffsetStep struct {
	parent StatementAcceptor
	limit  int32
	offset int64
}

func (s *SelectLimitOffsetStep) Parent() StatementAcceptor {
	return s.parent
}

func (s *SelectLimitOffsetStep) Accept(stmt *StatementImpl) error {
	stmt.Statement += fmt.Sprintf("LIMIT %d ", s.limit)
	if s.offset > 0 {
		stmt.Statement += fmt.Sprintf("OFFSET %d ", s.offset)
	}
	return nil
}
