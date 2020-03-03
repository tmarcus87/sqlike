package sqlike

import (
	"github.com/tmarcus87/sqlike/model"
	"strings"
)

type AndCondition struct {
	conditions []model.Condition
}

func (c *AndCondition) Apply(stmt *string, bindings *[]interface{}) {
	joinCondition(c.conditions, stmt, bindings, "AND")
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

func Or(conditions ...model.Condition) model.Condition {
	return &OrCondition{
		conditions: conditions,
	}
}

func joinCondition(conditions []model.Condition, stmt *string, bindings *[]interface{}, operator string) {
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

	*bindings = b
}
