package sqlike

const DialectMySQL = "mysql"

func init() {
	sqlDialect[DialectMySQL] =
		map[StatementType]string{
			StatementTypeSelectOne: "SELECT 1 FROM dual",
		}
}
