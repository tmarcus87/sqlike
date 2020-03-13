package model

import "fmt"

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
