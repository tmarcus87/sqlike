package sqlike

type StatementType int

const (
	StatementTypeUnknown StatementType = iota
	StatementTypeSelectOne
)

var sqlDialect = make(map[string]map[StatementType]string)
