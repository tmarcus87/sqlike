package model

func NewTextColumn(table Table, name string) *TextColumn {
	return &TextColumn{Table: table, Name: name}
}

type TextField interface {
	ColumnField
}

type TextColumn struct {
	Table Table
	Name  string
	alias string
	expr  string
	value string
}

func (c *TextColumn) SQLikeTable() Table {
	return c.Table
}

func (c *TextColumn) SQLikeColumnName() string {
	return c.Name
}

func (c *TextColumn) SQLikeAliasOrName() string {
	if c.alias != "" {
		return c.alias
	}
	return c.Name
}

func (c *TextColumn) As(alias string) ColumnField {
	c.alias = alias
	return c
}

func (c *TextColumn) SQLikeFieldExpr() string {
	return fieldExpr(c, c.alias, c.expr)
}

func (c *TextColumn) SQLikeSet(v string) ColumnValue {
	c.value = v
	return c
}

func (c *TextColumn) SQLikeColumnValue() interface{} {
	return c.value
}

func (c *TextColumn) Eq(v string) Condition {
	return &SingleValueCondition{Column: c, Operator: "=", Value: v}
}

func (c *TextColumn) NotEq(v string) Condition {
	return &SingleValueCondition{Column: c, Operator: "!=", Value: v}
}

func (c *TextColumn) Gt(v string) Condition {
	return &SingleValueCondition{Column: c, Operator: ">", Value: v}
}

func (c *TextColumn) GtOrEq(v string) Condition {
	return &SingleValueCondition{Column: c, Operator: ">=", Value: v}
}

func (c *TextColumn) Lt(v string) Condition {
	return &SingleValueCondition{Column: c, Operator: "<", Value: v}
}

func (c *TextColumn) LtOrEq(v string) Condition {
	return &SingleValueCondition{Column: c, Operator: "<=", Value: v}
}

func (c *TextColumn) Like(v string) Condition {
	return &SingleValueCondition{Column: c, Operator: "LIKE", Value: v}
}

func (c *TextColumn) IsNull() Condition {
	return &NoValueCondition{Column: c, Operator: "IS NULL"}
}

func (c *TextColumn) IsNotNull() Condition {
	return &NoValueCondition{Column: c, Operator: "IS NOT NULL"}
}

func (c *TextColumn) EqCol(field ColumnField) Condition {
	return &SingleColumnCondition{Column: c, Operator: "=", Value: field}
}

func (c *TextColumn) In(vs ...string) Condition {
	return &MultiValueCondition{
		Column:   c,
		Operator: "IN",
		Values:   StringSliceToInterfaceSlice(vs),
	}
}

func (c *TextColumn) NotIn(vs ...string) Condition {
	return &MultiValueCondition{
		Column:   c,
		Operator: "NOT IN",
		Values:   StringSliceToInterfaceSlice(vs),
	}
}

func (c *TextColumn) Asc() *SortOrder {
	return &SortOrder{
		Column: c,
		Order:  OrderAsc,
	}
}

func (c *TextColumn) Desc() *SortOrder {
	return &SortOrder{
		Column: c,
		Order:  OrderDesc,
	}
}

func StringSliceToInterfaceSlice(in []string) []interface{} {
	out := make([]interface{}, 0)
	for _, v := range in {
		out = append(out, v)
	}
	return out
}
