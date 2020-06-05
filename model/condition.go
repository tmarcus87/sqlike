package model

import (
	"fmt"
	"strings"
)

type Condition interface {
	Apply(stmt *string, bindings *[]interface{})
	And(condition Condition) Condition
	Or(condition Condition) Condition
}

func And(left, right Condition) Condition {
	return &AndCondition{
		left:  left,
		right: right,
	}
}

type AndCondition struct {
	left  Condition
	right Condition
}

func (c *AndCondition) Apply(stmt *string, bindings *[]interface{}) {
	JoinCondition([]Condition{c.left, c.right}, stmt, bindings, "AND")
}

func (c *AndCondition) And(condition Condition) Condition {
	return &AndCondition{
		left:  c,
		right: condition,
	}
}

func (c *AndCondition) Or(condition Condition) Condition {
	return &OrCondition{
		left:  c,
		right: condition,
	}
}

func Or(left, right Condition) Condition {
	return &OrCondition{
		left:  left,
		right: right,
	}
}

type OrCondition struct {
	left  Condition
	right Condition
}

func (c *OrCondition) Apply(stmt *string, bindings *[]interface{}) {
	JoinCondition([]Condition{c.left, c.right}, stmt, bindings, "OR")
}

func (c *OrCondition) And(condition Condition) Condition {
	return &AndCondition{
		left:  c,
		right: condition,
	}
}

func (c *OrCondition) Or(condition Condition) Condition {
	return &OrCondition{
		left:  c,
		right: condition,
	}
}

func JoinCondition(conditions []Condition, stmt *string, bindings *[]interface{}, operator string) {
	statements := make([]string, 0)
	b := make([]interface{}, 0)

	for _, condition := range conditions {
		statement := ""
		condition.Apply(&statement, &b)
		statements = append(statements, statement)
	}

	if len(conditions) > 1 {
		*stmt += "("
	}

	*stmt += strings.Join(statements, " "+operator+" ")

	if len(conditions) > 1 {
		*stmt += ")"
	}

	*bindings = append(*bindings, b...)
}

type NoValueCondition struct {
	Column   ColumnField
	Operator string
}

func (c *NoValueCondition) Apply(stmt *string, bindings *[]interface{}) {
	*stmt += fmt.Sprintf("`%s`.`%s` %s", c.Column.Table().SQLikeAliasOrName(), c.Column.ColumnName(), c.Operator)
}

func (c *NoValueCondition) And(condition Condition) Condition {
	return &AndCondition{
		left:  c,
		right: condition,
	}
}

func (c *NoValueCondition) Or(condition Condition) Condition {
	return &OrCondition{
		left:  c,
		right: condition,
	}
}

type SingleValueCondition struct {
	Column   ColumnField
	Operator string
	Value    interface{}
}

func (c *SingleValueCondition) Apply(stmt *string, bindings *[]interface{}) {
	*stmt += fmt.Sprintf("`%s`.`%s` %s ?", c.Column.Table().SQLikeAliasOrName(), c.Column.ColumnName(), c.Operator)
	*bindings = append(*bindings, c.Value)
}

func (c *SingleValueCondition) And(condition Condition) Condition {
	return &AndCondition{
		left:  c,
		right: condition,
	}
}

func (c *SingleValueCondition) Or(condition Condition) Condition {
	return &OrCondition{
		left:  c,
		right: condition,
	}
}

type MultiValueCondition struct {
	Column   ColumnField
	Operator string
	Values   []interface{}
}

func (c *MultiValueCondition) Apply(stmt *string, bindings *[]interface{}) {
	conds := make([]string, 0)
	for i := 0; i < len(c.Values); i++ {
		conds = append(conds, "?")
	}

	*stmt +=
		fmt.Sprintf("`%s`.`%s` %s (%s)",
			c.Column.Table().SQLikeAliasOrName(), c.Column.ColumnName(),
			c.Operator,
			strings.Join(conds, ", "))
	*bindings = append(*bindings, c.Values...)
}

func (c *MultiValueCondition) And(condition Condition) Condition {
	return &AndCondition{
		left:  c,
		right: condition,
	}
}

func (c *MultiValueCondition) Or(condition Condition) Condition {
	return &OrCondition{
		left:  c,
		right: condition,
	}
}

type SingleColumnCondition struct {
	Column   ColumnField
	Operator string
	Value    ColumnField
}

func (c *SingleColumnCondition) Apply(stmt *string, bindings *[]interface{}) {
	*stmt += fmt.Sprintf("`%s`.`%s` %s `%s`.`%s`",
		c.Column.Table().SQLikeAliasOrName(), c.Column.ColumnName(),
		c.Operator,
		c.Value.Table().SQLikeAliasOrName(), c.Value.ColumnName())
}

func (c *SingleColumnCondition) And(condition Condition) Condition {
	return &AndCondition{
		left:  c,
		right: condition,
	}
}

func (c *SingleColumnCondition) Or(condition Condition) Condition {
	return &OrCondition{
		left:  c,
		right: condition,
	}
}
