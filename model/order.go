package model

const (
	OrderAsc  = "ASC"
	OrderDesc = "DESC"
)

type SortOrder struct {
	Column ColumnField
	Order  string
}
