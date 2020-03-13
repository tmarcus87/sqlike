package model

//
//import (
//	"fmt"
//	"strings"
//)
//
//type BasicEqCondition struct {
//	column Column
//	v      interface{}
//}
//
//func (c *BasicEqCondition) Apply(stmt *string, bindings *[]interface{}) {
//	*stmt += fmt.Sprintf("`%s`.`%s` = ?", TableName(c.column.SQLikeTable()), c.column.SQLikeColumnName())
//	*bindings = append(*bindings, c.v)
//}
//
//type BasicEqColCondition struct {
//	left  Column
//	right Column
//}
//
//func (c *BasicEqColCondition) Apply(stmt *string, bindings *[]interface{}) {
//	*stmt +=
//		fmt.Sprintf("`%s`.`%s` = `%s`.`%s`",
//			TableName(c.left.SQLikeTable()), c.left.SQLikeColumnName(),
//			TableName(c.right.SQLikeTable()), c.right.SQLikeColumnName())
//}
//
//type NoValueCondition struct {
//	Column   Column
//	Operator string
//}
//
//func (c *NoValueCondition) Apply(stmt *string, bindings *[]interface{}) {
//	*stmt += fmt.Sprintf("`%s`.`%s` %s", TableName(c.Column.SQLikeTable()), c.Column.SQLikeColumnName(), c.Operator)
//}
//
//type SingleValueCondition struct {
//	Column   Column
//	Operator string
//	Value    interface{}
//}
//
//func (c *SingleValueCondition) Apply(stmt *string, bindings *[]interface{}) {
//	*stmt += fmt.Sprintf("`%s`.`%s` %s ?", TableName(c.Column.SQLikeTable()), c.Column.SQLikeColumnName(), c.Operator)
//	*bindings = append(*bindings, c.Value)
//}
//
//type MultiValueCondition struct {
//	Column   Column
//	Operator string
//	Values   []interface{}
//}
//
//func (c *MultiValueCondition) Apply(stmt *string, bindings *[]interface{}) {
//	conds := make([]string, 0)
//	for i := 0; i < len(c.Values); i++ {
//		conds = append(conds, "?")
//	}
//	*stmt += fmt.Sprintf("`%s`.`%s` %s (%s)", TableName(c.Column.SQLikeTable()), c.Column.SQLikeColumnName(), c.Operator, strings.Join(conds, ", "))
//	*bindings = append(*bindings, c.Values...)
//}
