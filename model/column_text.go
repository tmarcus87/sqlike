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

func (c *TextColumn) CondEq(v string) Condition {
	return &SingleValueCondition{Column: c, Operator: "=", Value: v}
}

func (c *TextColumn) CondNotEq(v string) Condition {
	return &SingleValueCondition{Column: c, Operator: "!=", Value: v}
}

func (c *TextColumn) CondGt(v string) Condition {
	return &SingleValueCondition{Column: c, Operator: ">", Value: v}
}

func (c *TextColumn) CondGtOrEq(v string) Condition {
	return &SingleValueCondition{Column: c, Operator: ">=", Value: v}
}

func (c *TextColumn) CondLt(v string) Condition {
	return &SingleValueCondition{Column: c, Operator: "<", Value: v}
}

func (c *TextColumn) CondLtOrEq(v string) Condition {
	return &SingleValueCondition{Column: c, Operator: "<=", Value: v}
}

func (c *TextColumn) CondLike(v string) Condition {
	return &SingleValueCondition{Column: c, Operator: "LIKE", Value: v}
}

func (c *TextColumn) CondIsNull() Condition {
	return &NoValueCondition{Column: c, Operator: "IS NULL"}
}

func (c *TextColumn) CondIsNotNull() Condition {
	return &NoValueCondition{Column: c, Operator: "IS NOT NULL"}
}

func (c *TextColumn) CondEqCol(field ColumnField) Condition {
	return &SingleColumnCondition{Column: c, Operator: "=", Value: field}
}

func (c *TextColumn) CondIn(vs ...string) Condition {
	return &MultiValueCondition{
		Column:   c,
		Operator: "IN",
		Values:   StringSliceToInterfaceSlice(vs),
	}
}

func (c *TextColumn) CondNotIn(vs ...string) Condition {
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
