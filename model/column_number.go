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

func (c *NumberColumn) As(alias string) ColumnField {
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

func (c *NumberColumn) Asc() *SortOrder {
	return &SortOrder{
		Column: c,
		Order:  OrderAsc,
	}
}

func (c *NumberColumn) Desc() *SortOrder {
	return &SortOrder{
		Column: c,
		Order:  OrderDesc,
	}
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

func (c *Int8Column) Eq(v int8) Condition {
	return &SingleValueCondition{Column: c, Operator: "=", Value: v}
}

func (c *Int8Column) NotEq(v int8) Condition {
	return &SingleValueCondition{Column: c, Operator: "!=", Value: v}
}

func (c *Int8Column) Gt(v int8) Condition {
	return &SingleValueCondition{Column: c, Operator: ">", Value: v}
}

func (c *Int8Column) GtOrEq(v int8) Condition {
	return &SingleValueCondition{Column: c, Operator: ">=", Value: v}
}

func (c *Int8Column) Lt(v int8) Condition {
	return &SingleValueCondition{Column: c, Operator: "<", Value: v}
}

func (c *Int8Column) LtOrEq(v int8) Condition {
	return &SingleValueCondition{Column: c, Operator: "<=", Value: v}
}

func (c *Int8Column) IsNull() Condition {
	return &NoValueCondition{Column: c, Operator: "IS NULL"}
}

func (c *Int8Column) IsNotNull() Condition {
	return &NoValueCondition{Column: c, Operator: "IS NOT NULL"}
}

func (c *Int8Column) EqCol(field ColumnField) Condition {
	return &SingleColumnCondition{Column: c, Operator: "=", Value: field}
}

func (c *Int8Column) In(vs ...int8) Condition {
	return &MultiValueCondition{
		Column:   c,
		Operator: "IN",
		Values:   Int8SliceToInterfaceSlice(vs),
	}
}

func (c *Int8Column) NotIn(vs ...int8) Condition {
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

func (c *Int16Column) Eq(v int16) Condition {
	return &SingleValueCondition{Column: c, Operator: "=", Value: v}
}

func (c *Int16Column) NotEq(v int16) Condition {
	return &SingleValueCondition{Column: c, Operator: "!=", Value: v}
}

func (c *Int16Column) Gt(v int16) Condition {
	return &SingleValueCondition{Column: c, Operator: ">", Value: v}
}

func (c *Int16Column) GtOrEq(v int16) Condition {
	return &SingleValueCondition{Column: c, Operator: ">=", Value: v}
}

func (c *Int16Column) Lt(v int16) Condition {
	return &SingleValueCondition{Column: c, Operator: "<", Value: v}
}

func (c *Int16Column) LtOrEq(v int16) Condition {
	return &SingleValueCondition{Column: c, Operator: "<=", Value: v}
}

func (c *Int16Column) IsNull() Condition {
	return &NoValueCondition{Column: c, Operator: "IS NULL"}
}

func (c *Int16Column) IsNotNull() Condition {
	return &NoValueCondition{Column: c, Operator: "IS NOT NULL"}
}

func (c *Int16Column) EqCol(field ColumnField) Condition {
	return &SingleColumnCondition{Column: c, Operator: "=", Value: field}
}

func (c *Int16Column) In(vs ...int16) Condition {
	return &MultiValueCondition{
		Column:   c,
		Operator: "IN",
		Values:   Int16SliceToInterfaceSlice(vs),
	}
}

func (c *Int16Column) NotIn(vs ...int16) Condition {
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

func (c *Int32Column) Eq(v int32) Condition {
	return &SingleValueCondition{Column: c, Operator: "=", Value: v}
}

func (c *Int32Column) NotEq(v int32) Condition {
	return &SingleValueCondition{Column: c, Operator: "!=", Value: v}
}

func (c *Int32Column) Gt(v int32) Condition {
	return &SingleValueCondition{Column: c, Operator: ">", Value: v}
}

func (c *Int32Column) GtOrEq(v int32) Condition {
	return &SingleValueCondition{Column: c, Operator: ">=", Value: v}
}

func (c *Int32Column) Lt(v int32) Condition {
	return &SingleValueCondition{Column: c, Operator: "<", Value: v}
}

func (c *Int32Column) LtOrEq(v int32) Condition {
	return &SingleValueCondition{Column: c, Operator: "<=", Value: v}
}

func (c *Int32Column) IsNull() Condition {
	return &NoValueCondition{Column: c, Operator: "IS NULL"}
}

func (c *Int32Column) IsNotNull() Condition {
	return &NoValueCondition{Column: c, Operator: "IS NOT NULL"}
}

func (c *Int32Column) EqCol(field ColumnField) Condition {
	return &SingleColumnCondition{Column: c, Operator: "=", Value: field}
}

func (c *Int32Column) In(vs ...int32) Condition {
	return &MultiValueCondition{
		Column:   c,
		Operator: "IN",
		Values:   Int32SliceToInterfaceSlice(vs),
	}
}

func (c *Int32Column) NotIn(vs ...int32) Condition {
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

func (c *Int64Column) Eq(v int64) Condition {
	return &SingleValueCondition{Column: c, Operator: "=", Value: v}
}

func (c *Int64Column) NotEq(v int64) Condition {
	return &SingleValueCondition{Column: c, Operator: "!=", Value: v}
}

func (c *Int64Column) Gt(v int64) Condition {
	return &SingleValueCondition{Column: c, Operator: ">", Value: v}
}

func (c *Int64Column) GtOrEq(v int64) Condition {
	return &SingleValueCondition{Column: c, Operator: ">=", Value: v}
}

func (c *Int64Column) Lt(v int64) Condition {
	return &SingleValueCondition{Column: c, Operator: "<", Value: v}
}

func (c *Int64Column) LtOrEq(v int64) Condition {
	return &SingleValueCondition{Column: c, Operator: "<=", Value: v}
}

func (c *Int64Column) IsNull() Condition {
	return &NoValueCondition{Column: c, Operator: "IS NULL"}
}

func (c *Int64Column) IsNotNull() Condition {
	return &NoValueCondition{Column: c, Operator: "IS NOT NULL"}
}

func (c *Int64Column) EqCol(field ColumnField) Condition {
	return &SingleColumnCondition{Column: c, Operator: "=", Value: field}
}

func (c *Int64Column) In(vs ...int64) Condition {
	return &MultiValueCondition{
		Column:   c,
		Operator: "IN",
		Values:   Int64SliceToInterfaceSlice(vs),
	}
}

func (c *Int64Column) NotIn(vs ...int64) Condition {
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

func (c *Float32Column) Eq(v float32) Condition {
	return &SingleValueCondition{Column: c, Operator: "=", Value: v}
}

func (c *Float32Column) NotEq(v float32) Condition {
	return &SingleValueCondition{Column: c, Operator: "!=", Value: v}
}

func (c *Float32Column) Gt(v float32) Condition {
	return &SingleValueCondition{Column: c, Operator: ">", Value: v}
}

func (c *Float32Column) GtOrEq(v float32) Condition {
	return &SingleValueCondition{Column: c, Operator: ">=", Value: v}
}

func (c *Float32Column) Lt(v float32) Condition {
	return &SingleValueCondition{Column: c, Operator: "<", Value: v}
}

func (c *Float32Column) LtOrEq(v float32) Condition {
	return &SingleValueCondition{Column: c, Operator: "<=", Value: v}
}

func (c *Float32Column) IsNull() Condition {
	return &NoValueCondition{Column: c, Operator: "IS NULL"}
}

func (c *Float32Column) IsNotNull() Condition {
	return &NoValueCondition{Column: c, Operator: "IS NOT NULL"}
}

func (c *Float32Column) EqCol(field ColumnField) Condition {
	return &SingleColumnCondition{Column: c, Operator: "=", Value: field}
}

func (c *Float32Column) In(vs ...float32) Condition {
	return &MultiValueCondition{
		Column:   c,
		Operator: "IN",
		Values:   Float32SliceToInterfaceSlice(vs),
	}
}

func (c *Float32Column) NotIn(vs ...float32) Condition {
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

func (c *Float64Column) Eq(v float64) Condition {
	return &SingleValueCondition{Column: c, Operator: "=", Value: v}
}

func (c *Float64Column) NotEq(v float64) Condition {
	return &SingleValueCondition{Column: c, Operator: "!=", Value: v}
}

func (c *Float64Column) Gt(v float64) Condition {
	return &SingleValueCondition{Column: c, Operator: ">", Value: v}
}

func (c *Float64Column) GtOrEq(v float64) Condition {
	return &SingleValueCondition{Column: c, Operator: ">=", Value: v}
}

func (c *Float64Column) Lt(v float64) Condition {
	return &SingleValueCondition{Column: c, Operator: "<", Value: v}
}

func (c *Float64Column) LtOrEq(v float64) Condition {
	return &SingleValueCondition{Column: c, Operator: "<=", Value: v}
}

func (c *Float64Column) IsNull() Condition {
	return &NoValueCondition{Column: c, Operator: "IS NULL"}
}

func (c *Float64Column) IsNotNull() Condition {
	return &NoValueCondition{Column: c, Operator: "IS NOT NULL"}
}

func (c *Float64Column) EqCol(field ColumnField) Condition {
	return &SingleColumnCondition{Column: c, Operator: "=", Value: field}
}

func (c *Float64Column) In(vs ...float64) Condition {
	return &MultiValueCondition{
		Column:   c,
		Operator: "IN",
		Values:   Float64SliceToInterfaceSlice(vs),
	}
}

func (c *Float64Column) NotIn(vs ...float64) Condition {
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
