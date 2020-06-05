package statement

import (
	"github.com/tmarcus87/sqlike/model"
)

type AndCondition struct {
	conditions []model.Condition
}

func (c *AndCondition) Apply(stmt *string, bindings *[]interface{}) {
	joinCondition(c.conditions, stmt, bindings, "AND")
}

func (c *AndCondition) And(condition model.Condition) model.Condition {
	return model.And(c, condition)
}

func (c *AndCondition) Or(condition model.Condition) model.Condition {
	return model.Or(c, condition)
}

func And(conditions ...model.Condition) model.Condition {
	return &AndCondition{
		conditions: conditions,
	}
}

type OrCondition struct {
	conditions []model.Condition
}

func (c *OrCondition) Apply(stmt *string, bindings *[]interface{}) {
	joinCondition(c.conditions, stmt, bindings, "OR")
}

func (c *OrCondition) And(condition model.Condition) model.Condition {
	return model.And(c, condition)
}

func (c *OrCondition) Or(condition model.Condition) model.Condition {
	return model.Or(c, condition)
}

func Or(conditions ...model.Condition) model.Condition {
	return &OrCondition{
		conditions: conditions,
	}
}

func joinCondition(conditions []model.Condition, stmt *string, bindings *[]interface{}, operator string) {
	model.JoinCondition(conditions, stmt, bindings, operator)
}
