package dialect

const MySQL = "mysql"

func init() {
	sqlDialect[MySQL] =
		map[StatementType]string{
			StatementTypeSelectOne: "SELECT 1 FROM dual",
		}
}
