package model

import "fmt"

type NumericField interface {
	ColumnField
	PlusInt(v int) NumericField
	PlusFloat(v float64) NumericField
	MinusInt(v int) NumericField
	MinusFloat(v float64) NumericField
	MultipleInt(v int) NumericField
	MultipleFloat(v float64) NumericField
	DivideInt(v int) NumericField
	DivideFloat(v float64) NumericField
}

type NumberColumn struct {
	Table Table
	Name  string
	alias string
	expr  string
}

func (c *NumberColumn) SQLikeTable() Table {
	return c.Table
}

func (c *NumberColumn) SQLikeColumnName() string {
	return c.Name
}

func (c *NumberColumn) SQLikeAliasOrName() string {
	if c.alias != "" {
		return c.alias
	}
	return c.Name

}

func (c *NumberColumn) SQLikeAs(alias string) ColumnField {
	c.alias = alias
	return c
}

func (c *NumberColumn) SQLikeFieldExpr() string {
	return fieldExpr(c, c.alias, c.expr)
}

func (c *NumberColumn) PlusInt(v int) NumericField {
	c.expr = calcExpr(c, c.expr, fmt.Sprintf("$$ + %d", v))
	return c
}

func (c *NumberColumn) PlusFloat(v float64) NumericField {
	c.expr = calcExpr(c, c.expr, fmt.Sprintf("$$ + %g", v))
	return c
}

func (c *NumberColumn) MinusInt(v int) NumericField {
	c.expr = calcExpr(c, c.expr, fmt.Sprintf("$$ - %d", v))
	return c
}

func (c *NumberColumn) MinusFloat(v float64) NumericField {
	c.expr = calcExpr(c, c.expr, fmt.Sprintf("$$ - %g", v))
	return c
}

func (c *NumberColumn) MultipleInt(v int) NumericField {
	c.expr = calcExpr(c, c.expr, fmt.Sprintf("$$ * %d", v))
	return c
}

func (c *NumberColumn) MultipleFloat(v float64) NumericField {
	c.expr = calcExpr(c, c.expr, fmt.Sprintf("$$ * %g", v))
	return c
}

func (c *NumberColumn) DivideInt(v int) NumericField {
	c.expr = calcExpr(c, c.expr, fmt.Sprintf("$$ / %d", v))
	return c
}

func (c *NumberColumn) DivideFloat(v float64) NumericField {
	c.expr = calcExpr(c, c.expr, fmt.Sprintf("$$ / %g", v))
	return c
}

func NewInt8Column(table Table, name string) *Int8Column {
	return &Int8Column{NumberColumn: NumberColumn{Table: table, Name: name}}
}

type Int8Column struct {
	NumberColumn
	value int8
}

func (c *Int8Column) SQLikeSet(v int8) ColumnValue {
	c.value = v
	return c
}

func (c *Int8Column) SQLikeColumnValue() interface{} {
	return c.value
}

func (c *Int8Column) CondEq(v int8) Condition {
	return &SingleValueCondition{Column: c, Operator: "=", Value: v}
}

func (c *Int8Column) CondNotEq(v int8) Condition {
	return &SingleValueCondition{Column: c, Operator: "!=", Value: v}
}

func (c *Int8Column) CondGt(v int8) Condition {
	return &SingleValueCondition{Column: c, Operator: ">", Value: v}
}

func (c *Int8Column) CondGtOrEq(v int8) Condition {
	return &SingleValueCondition{Column: c, Operator: ">=", Value: v}
}

func (c *Int8Column) CondLt(v int8) Condition {
	return &SingleValueCondition{Column: c, Operator: "<", Value: v}
}

func (c *Int8Column) CondLtOrEq(v int8) Condition {
	return &SingleValueCondition{Column: c, Operator: "<=", Value: v}
}

func (c *Int8Column) CondIsNull() Condition {
	return &NoValueCondition{Column: c, Operator: "IS NULL"}
}

func (c *Int8Column) CondIsNotNull() Condition {
	return &NoValueCondition{Column: c, Operator: "IS NOT NULL"}
}

func (c *Int8Column) CondEqCol(field ColumnField) Condition {
	return &SingleColumnCondition{Column: c, Operator: "=", Value: field}
}

func (c *Int8Column) CondIn(vs ...int8) Condition {
	return &MultiValueCondition{
		Column:   c,
		Operator: "IN",
		Values:   Int8SliceToInterfaceSlice(vs),
	}
}

func (c *Int8Column) CondNotIn(vs ...int8) Condition {
	return &MultiValueCondition{
		Column:   c,
		Operator: "NOT IN",
		Values:   Int8SliceToInterfaceSlice(vs),
	}
}

func Int8SliceToInterfaceSlice(in []int8) []interface{} {
	out := make([]interface{}, 0)
	for _, v := range in {
		out = append(out, v)
	}
	return out
}

func NewInt16Column(table Table, name string) *Int16Column {
	return &Int16Column{NumberColumn: NumberColumn{Table: table, Name: name}}
}

type Int16Column struct {
	NumberColumn
	value int16
}

func (c *Int16Column) SQLikeSet(v int16) ColumnValue {
	c.value = v
	return c
}

func (c *Int16Column) SQLikeColumnValue() interface{} {
	return c.value
}

func (c *Int16Column) CondEq(v int16) Condition {
	return &SingleValueCondition{Column: c, Operator: "=", Value: v}
}

func (c *Int16Column) CondNotEq(v int16) Condition {
	return &SingleValueCondition{Column: c, Operator: "!=", Value: v}
}

func (c *Int16Column) CondGt(v int16) Condition {
	return &SingleValueCondition{Column: c, Operator: ">", Value: v}
}

func (c *Int16Column) CondGtOrEq(v int16) Condition {
	return &SingleValueCondition{Column: c, Operator: ">=", Value: v}
}

func (c *Int16Column) CondLt(v int16) Condition {
	return &SingleValueCondition{Column: c, Operator: "<", Value: v}
}

func (c *Int16Column) CondLtOrEq(v int16) Condition {
	return &SingleValueCondition{Column: c, Operator: "<=", Value: v}
}

func (c *Int16Column) CondIsNull() Condition {
	return &NoValueCondition{Column: c, Operator: "IS NULL"}
}

func (c *Int16Column) CondIsNotNull() Condition {
	return &NoValueCondition{Column: c, Operator: "IS NOT NULL"}
}

func (c *Int16Column) CondEqCol(field ColumnField) Condition {
	return &SingleColumnCondition{Column: c, Operator: "=", Value: field}
}

func (c *Int16Column) CondIn(vs ...int16) Condition {
	return &MultiValueCondition{
		Column:   c,
		Operator: "IN",
		Values:   Int16SliceToInterfaceSlice(vs),
	}
}

func (c *Int16Column) CondNotIn(vs ...int16) Condition {
	return &MultiValueCondition{
		Column:   c,
		Operator: "NOT IN",
		Values:   Int16SliceToInterfaceSlice(vs),
	}
}

func Int16SliceToInterfaceSlice(in []int16) []interface{} {
	out := make([]interface{}, 0)
	for _, v := range in {
		out = append(out, v)
	}
	return out
}

func NewInt32Column(table Table, name string) *Int32Column {
	return &Int32Column{NumberColumn: NumberColumn{Table: table, Name: name}}
}

type Int32Column struct {
	NumberColumn
	value int32
}

func (c *Int32Column) SQLikeSet(v int32) ColumnValue {
	c.value = v
	return c
}

func (c *Int32Column) SQLikeColumnValue() interface{} {
	return c.value
}

func (c *Int32Column) CondEq(v int32) Condition {
	return &SingleValueCondition{Column: c, Operator: "=", Value: v}
}

func (c *Int32Column) CondNotEq(v int32) Condition {
	return &SingleValueCondition{Column: c, Operator: "!=", Value: v}
}

func (c *Int32Column) CondGt(v int32) Condition {
	return &SingleValueCondition{Column: c, Operator: ">", Value: v}
}

func (c *Int32Column) CondGtOrEq(v int32) Condition {
	return &SingleValueCondition{Column: c, Operator: ">=", Value: v}
}

func (c *Int32Column) CondLt(v int32) Condition {
	return &SingleValueCondition{Column: c, Operator: "<", Value: v}
}

func (c *Int32Column) CondLtOrEq(v int32) Condition {
	return &SingleValueCondition{Column: c, Operator: "<=", Value: v}
}

func (c *Int32Column) CondIsNull() Condition {
	return &NoValueCondition{Column: c, Operator: "IS NULL"}
}

func (c *Int32Column) CondIsNotNull() Condition {
	return &NoValueCondition{Column: c, Operator: "IS NOT NULL"}
}

func (c *Int32Column) CondEqCol(field ColumnField) Condition {
	return &SingleColumnCondition{Column: c, Operator: "=", Value: field}
}

func (c *Int32Column) CondIn(vs ...int32) Condition {
	return &MultiValueCondition{
		Column:   c,
		Operator: "IN",
		Values:   Int32SliceToInterfaceSlice(vs),
	}
}

func (c *Int32Column) CondNotIn(vs ...int32) Condition {
	return &MultiValueCondition{
		Column:   c,
		Operator: "NOT IN",
		Values:   Int32SliceToInterfaceSlice(vs),
	}
}

func Int32SliceToInterfaceSlice(in []int32) []interface{} {
	out := make([]interface{}, 0)
	for _, v := range in {
		out = append(out, v)
	}
	return out
}

func NewInt64Column(table Table, name string) *Int64Column {
	return &Int64Column{NumberColumn: NumberColumn{Table: table, Name: name}}
}

type Int64Column struct {
	NumberColumn
	value int64
}

func (c *Int64Column) SQLikeSet(v int64) ColumnValue {
	c.value = v
	return c
}

func (c *Int64Column) SQLikeColumnValue() interface{} {
	return c.value
}

func (c *Int64Column) CondEq(v int64) Condition {
	return &SingleValueCondition{Column: c, Operator: "=", Value: v}
}

func (c *Int64Column) CondNotEq(v int64) Condition {
	return &SingleValueCondition{Column: c, Operator: "!=", Value: v}
}

func (c *Int64Column) CondGt(v int64) Condition {
	return &SingleValueCondition{Column: c, Operator: ">", Value: v}
}

func (c *Int64Column) CondGtOrEq(v int64) Condition {
	return &SingleValueCondition{Column: c, Operator: ">=", Value: v}
}

func (c *Int64Column) CondLt(v int64) Condition {
	return &SingleValueCondition{Column: c, Operator: "<", Value: v}
}

func (c *Int64Column) CondLtOrEq(v int64) Condition {
	return &SingleValueCondition{Column: c, Operator: "<=", Value: v}
}

func (c *Int64Column) CondIsNull() Condition {
	return &NoValueCondition{Column: c, Operator: "IS NULL"}
}

func (c *Int64Column) CondIsNotNull() Condition {
	return &NoValueCondition{Column: c, Operator: "IS NOT NULL"}
}

func (c *Int64Column) CondEqCol(field ColumnField) Condition {
	return &SingleColumnCondition{Column: c, Operator: "=", Value: field}
}

func (c *Int64Column) CondIn(vs ...int64) Condition {
	return &MultiValueCondition{
		Column:   c,
		Operator: "IN",
		Values:   Int64SliceToInterfaceSlice(vs),
	}
}

func (c *Int64Column) CondNotIn(vs ...int64) Condition {
	return &MultiValueCondition{
		Column:   c,
		Operator: "NOT IN",
		Values:   Int64SliceToInterfaceSlice(vs),
	}
}

func Int64SliceToInterfaceSlice(in []int64) []interface{} {
	out := make([]interface{}, 0)
	for _, v := range in {
		out = append(out, v)
	}
	return out
}

func NewFloat32Column(table Table, name string) *Float32Column {
	return &Float32Column{NumberColumn: NumberColumn{Table: table, Name: name}}
}

type Float32Column struct {
	NumberColumn
	value float32
}

func (c *Float32Column) SQLikeSet(v float32) ColumnValue {
	c.value = v
	return c
}

func (c *Float32Column) SQLikeColumnValue() interface{} {
	return c.value
}

func (c *Float32Column) CondEq(v float32) Condition {
	return &SingleValueCondition{Column: c, Operator: "=", Value: v}
}

func (c *Float32Column) CondNotEq(v float32) Condition {
	return &SingleValueCondition{Column: c, Operator: "!=", Value: v}
}

func (c *Float32Column) CondGt(v float32) Condition {
	return &SingleValueCondition{Column: c, Operator: ">", Value: v}
}

func (c *Float32Column) CondGtOrEq(v float32) Condition {
	return &SingleValueCondition{Column: c, Operator: ">=", Value: v}
}

func (c *Float32Column) CondLt(v float32) Condition {
	return &SingleValueCondition{Column: c, Operator: "<", Value: v}
}

func (c *Float32Column) CondLtOrEq(v float32) Condition {
	return &SingleValueCondition{Column: c, Operator: "<=", Value: v}
}

func (c *Float32Column) CondIsNull() Condition {
	return &NoValueCondition{Column: c, Operator: "IS NULL"}
}

func (c *Float32Column) CondIsNotNull() Condition {
	return &NoValueCondition{Column: c, Operator: "IS NOT NULL"}
}

func (c *Float32Column) CondEqCol(field ColumnField) Condition {
	return &SingleColumnCondition{Column: c, Operator: "=", Value: field}
}

func (c *Float32Column) CondIn(vs ...float32) Condition {
	return &MultiValueCondition{
		Column:   c,
		Operator: "IN",
		Values:   Float32SliceToInterfaceSlice(vs),
	}
}

func (c *Float32Column) CondNotIn(vs ...float32) Condition {
	return &MultiValueCondition{
		Column:   c,
		Operator: "NOT IN",
		Values:   Float32SliceToInterfaceSlice(vs),
	}
}

func Float32SliceToInterfaceSlice(in []float32) []interface{} {
	out := make([]interface{}, 0)
	for _, v := range in {
		out = append(out, v)
	}
	return out
}

func NewFloat64Column(table Table, name string) *Float64Column {
	return &Float64Column{NumberColumn: NumberColumn{Table: table, Name: name}}
}

type Float64Column struct {
	NumberColumn
	value float64
}

func (c *Float64Column) SQLikeSet(v float64) ColumnValue {
	c.value = v
	return c
}

func (c *Float64Column) SQLikeColumnValue() interface{} {
	return c.value
}

func (c *Float64Column) CondEq(v float64) Condition {
	return &SingleValueCondition{Column: c, Operator: "=", Value: v}
}

func (c *Float64Column) CondNotEq(v float64) Condition {
	return &SingleValueCondition{Column: c, Operator: "!=", Value: v}
}

func (c *Float64Column) CondGt(v float64) Condition {
	return &SingleValueCondition{Column: c, Operator: ">", Value: v}
}

func (c *Float64Column) CondGtOrEq(v float64) Condition {
	return &SingleValueCondition{Column: c, Operator: ">=", Value: v}
}

func (c *Float64Column) CondLt(v float64) Condition {
	return &SingleValueCondition{Column: c, Operator: "<", Value: v}
}

func (c *Float64Column) CondLtOrEq(v float64) Condition {
	return &SingleValueCondition{Column: c, Operator: "<=", Value: v}
}

func (c *Float64Column) CondIsNull() Condition {
	return &NoValueCondition{Column: c, Operator: "IS NULL"}
}

func (c *Float64Column) CondIsNotNull() Condition {
	return &NoValueCondition{Column: c, Operator: "IS NOT NULL"}
}

func (c *Float64Column) CondEqCol(field ColumnField) Condition {
	return &SingleColumnCondition{Column: c, Operator: "=", Value: field}
}

func (c *Float64Column) CondIn(vs ...float64) Condition {
	return &MultiValueCondition{
		Column:   c,
		Operator: "IN",
		Values:   Float64SliceToInterfaceSlice(vs),
	}
}

func (c *Float64Column) CondNotIn(vs ...float64) Condition {
	return &MultiValueCondition{
		Column:   c,
		Operator: "NOT IN",
		Values:   Float64SliceToInterfaceSlice(vs),
	}
}

func Float64SliceToInterfaceSlice(in []float64) []interface{} {
	out := make([]interface{}, 0)
	for _, v := range in {
		out = append(out, v)
	}
	return out
}
