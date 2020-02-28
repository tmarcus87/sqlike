package sqlike

import "strings"

type AndCondition struct {
	conditions []Condition
}

func (c *AndCondition) Apply(stmt *string, bindings *[]interface{}) {
	joinCondition(c.conditions, stmt, bindings, "AND")
}

func And(conditions ...Condition) Condition {
	return &AndCondition{
		conditions: conditions,
	}
}

type OrCondition struct {
	conditions []Condition
}

func (c *OrCondition) Apply(stmt *string, bindings *[]interface{}) {
	joinCondition(c.conditions, stmt, bindings, "OR")
}

func Or(conditions ...Condition) Condition {
	return &OrCondition{
		conditions: conditions,
	}
}

func joinCondition(conditions []Condition, stmt *string, bindings *[]interface{}, operator string) {
	statements := make([]string, 0)
	b := make([]interface{}, 0)
	for _, condition := range conditions {
		statement := ""
		condition.Apply(&statement, &b)
		statements = append(statements, statement)
	}
	*stmt += "(" + strings.Join(statements, " "+operator+" ") + ")"
	*bindings = b
}
