package sqlike

const DialectSqlite3 = "seqlite3"

func init() {
	sqlDialect[DialectSqlite3] =
		map[StatementType]string{
			StatementTypeSelectOne: "SELECT 1",
		}
}
