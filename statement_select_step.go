package sqlike

import (
	"fmt"
	"strings"
)

type SelectExplainStep struct {
	parent StatementAcceptor
}

func (s *SelectExplainStep) DialectStatement(st StatementType) string {
	return s.parent.DialectStatement(st)
}

func (s *SelectExplainStep) Parent() StatementAcceptor {
	return s.parent
}

func (s *SelectExplainStep) Accept(stmt *StatementImpl) {
	stmt.Statement += "EXPLAIN "
}

type SelectColumnStep struct {
	parent  StatementAcceptor
	columns []Column
}

func (s *SelectColumnStep) DialectStatement(st StatementType) string {
	return s.parent.DialectStatement(st)
}

func (s *SelectColumnStep) Parent() StatementAcceptor {
	return s.parent
}

func (s *SelectColumnStep) Accept(stmt *StatementImpl) {
	cols := make([]string, 0)
	for _, column := range s.columns {
		cols = append(cols, ColumnAsStatement(column))
	}
	stmt.Statement += fmt.Sprintf("SELECT %s ", strings.Join(cols, ", "))
}

type SelectFromStep struct {
	parent StatementAcceptor
	table  Table
}

func (s *SelectFromStep) DialectStatement(st StatementType) string {
	return s.parent.DialectStatement(st)
}

func (s *SelectFromStep) Parent() StatementAcceptor {
	return s.parent
}

func (s *SelectFromStep) Accept(stmt *StatementImpl) {
	stmt.Statement += fmt.Sprintf("FROM %s ", TableAsStatement(s.table))
}

type SelectGroupByStep struct {
	parent  StatementAcceptor
	columns []Column
}

func (s *SelectGroupByStep) DialectStatement(st StatementType) string {
	return s.parent.DialectStatement(st)
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
		cols = append(cols, fmt.Sprintf("`%s`.`%s`", TableName(column.SQLikeTable()), column.SQLikeColumnName()))
	}

	stmt.Statement += fmt.Sprintf("ORDER BY %s ", strings.Join(cols, ", "))
}

const (
	OrderAsc  = "ASC"
	OrderDesc = "DESC"
)

func NewOrder(column Column, order string) *SortOrder {
	return &SortOrder{
		column: column,
		order:  order,
	}
}

type SortOrder struct {
	column Column
	order  string
}

type SelectOrderByStep struct {
	parent StatementAcceptor
	orders []*SortOrder
}

func (s *SelectOrderByStep) DialectStatement(st StatementType) string {
	return s.DialectStatement(st)
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
				TableName(order.column.SQLikeTable()), ColumnName(order.column), order.order))
	}

	stmt.Statement += fmt.Sprintf("ORDER BY %s ", strings.Join(orders, ", "))
}

type SelectLimitOffsetStep struct {
	parent StatementAcceptor
	limit  int32
	offset int64
}

func (s *SelectLimitOffsetStep) DialectStatement(st StatementType) string {
	return s.parent.DialectStatement(st)
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
