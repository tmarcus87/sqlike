package model

import (
	"fmt"
	"strings"
)

type Condition interface {
	Apply(stmt *string, bindings *[]interface{})
}

type NoValueCondition struct {
	Column   ColumnField
	Operator string
}

func (c *NoValueCondition) Apply(stmt *string, bindings *[]interface{}) {
	*stmt += fmt.Sprintf("`%s`.`%s` %s", c.Column.SQLikeTable().SQLikeAliasOrName(), c.Column.SQLikeColumnName(), c.Operator)
}

type SingleValueCondition struct {
	Column   ColumnField
	Operator string
	Value    interface{}
}

func (c *SingleValueCondition) Apply(stmt *string, bindings *[]interface{}) {
	*stmt += fmt.Sprintf("`%s`.`%s` %s ?", c.Column.SQLikeTable().SQLikeAliasOrName(), c.Column.SQLikeColumnName(), c.Operator)
	*bindings = append(*bindings, c.Value)
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
			c.Column.SQLikeTable().SQLikeAliasOrName(), c.Column.SQLikeColumnName(),
			c.Operator,
			strings.Join(conds, ", "))
	*bindings = append(*bindings, c.Values...)
}

type SingleColumnCondition struct {
	Column   ColumnField
	Operator string
	Value    ColumnField
}

func (c *SingleColumnCondition) Apply(stmt *string, bindings *[]interface{}) {
	*stmt += fmt.Sprintf("`%s`.`%s` %s `%s`.`%s`",
		c.Column.SQLikeTable().SQLikeAliasOrName(), c.Column.SQLikeColumnName(),
		c.Operator,
		c.Value.SQLikeTable().SQLikeAliasOrName(), c.Value.SQLikeColumnName())
}
