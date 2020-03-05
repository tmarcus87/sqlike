package model

const (
	OrderAsc  = "ASC"
	OrderDesc = "DESC"
)

type SortOrder struct {
	Column Column
	Order  string
}
