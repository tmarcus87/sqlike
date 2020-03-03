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

func (s *SelectExplainStep) Accept(stmt *StatementImpl) {
	stmt.Statement += "EXPLAIN "
}

type SelectOneStep struct {
	parent StatementAcceptor
}

func (s *SelectOneStep) Parent() StatementAcceptor {
	return s.parent
}

func (s *SelectOneStep) Accept(stmt *StatementImpl) {
	stmt.Statement += getQueryer(s.parent).DialectStatement(dialect.StatementTypeSelectOne)
}

type SelectColumnStep struct {
	parent  StatementAcceptor
	columns []model.Column
}

func (s *SelectColumnStep) Parent() StatementAcceptor {
	return s.parent
}

func (s *SelectColumnStep) Accept(stmt *StatementImpl) {
	cols := make([]string, 0)
	for _, column := range s.columns {
		s := model.ColumnAsStatement(column)
		if column.SQLikeSelectModFmt() != "" {
			s = fmt.Sprintf(column.SQLikeSelectModFmt(), s)
		}
		cols = append(cols, s)
	}
	stmt.Statement += fmt.Sprintf("SELECT %s ", strings.Join(cols, ", "))
}

type SelectFromStep struct {
	parent StatementAcceptor
	table  model.Table
}

func (s *SelectFromStep) Parent() StatementAcceptor {
	return s.parent
}

func (s *SelectFromStep) Accept(stmt *StatementImpl) {
	stmt.Statement += fmt.Sprintf("FROM %s ", model.TableAsStatement(s.table))
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

func (s *SelectFromJoinStep) Accept(stmt *StatementImpl) {
	var onStmt string
	joinCondition(s.conditions, &onStmt, &stmt.Bindings, "AND")

	stmt.Statement += fmt.Sprintf("%s %s ON %s ", s.joinType, model.TableAsStatement(s.table), onStmt)
}

type SelectGroupByStep struct {
	parent  StatementAcceptor
	columns []model.Column
}

func (s *SelectGroupByStep) Parent() StatementAcceptor {
	return s.parent
}

func (s *SelectGroupByStep) Accept(stmt *StatementImpl) {
	if len(s.columns) == 0 {
		return
	}

	cols := make([]string, 0)
	for _, column := range s.columns {
		cols = append(cols, fmt.Sprintf("`%s`.`%s`", model.TableName(column.SQLikeTable()), column.SQLikeColumnName()))
	}

	stmt.Statement += fmt.Sprintf("GROUP BY %s ", strings.Join(cols, ", "))
}

const (
	OrderAsc  = "ASC"
	OrderDesc = "DESC"
)

func Order(column model.Column, order string) *SortOrder {
	return &SortOrder{
		column: column,
		order:  order,
	}
}

type SortOrder struct {
	column model.Column
	order  string
}

type SelectOrderByStep struct {
	parent StatementAcceptor
	orders []*SortOrder
}

func (s *SelectOrderByStep) Parent() StatementAcceptor {
	return s.parent
}

func (s *SelectOrderByStep) Accept(stmt *StatementImpl) {
	if len(s.orders) == 0 {
		return
	}

	orders := make([]string, 0)
	for _, order := range s.orders {
		orders =
			append(orders, fmt.Sprintf("`%s`.`%s` %s",
				model.TableName(order.column.SQLikeTable()), model.ColumnName(order.column), order.order))
	}

	stmt.Statement += fmt.Sprintf("ORDER BY %s ", strings.Join(orders, ", "))
}

type SelectLimitOffsetStep struct {
	parent StatementAcceptor
	limit  int32
	offset int64
}

func (s *SelectLimitOffsetStep) Parent() StatementAcceptor {
	return s.parent
}

func (s *SelectLimitOffsetStep) Accept(stmt *StatementImpl) {
	stmt.Statement += fmt.Sprintf("LIMIT %d ", s.limit)
	if s.offset > 0 {
		stmt.Statement += fmt.Sprintf("OFFSET %d ", s.offset)
	}
}
