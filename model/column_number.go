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
