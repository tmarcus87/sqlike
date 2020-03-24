package dialect

type StatementType int

const (
	StatementTypeUnknown StatementType = iota
	StatementTypeSelectOne
	StatementTypeOnDuplicateKeyIgnore
	StatementTypeOnDuplicateKeyUpdate
)

var sqlDialect = make(map[string]map[StatementType]string)

func GetDialectStatements(dialect string) map[StatementType]string {
	return sqlDialect[dialect]
}
