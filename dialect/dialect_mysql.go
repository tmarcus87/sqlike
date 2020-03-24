package dialect

const MySQL = "mysql"

func init() {
	sqlDialect[MySQL] =
		map[StatementType]string{
			StatementTypeSelectOne:            "SELECT 1 FROM dual",
			StatementTypeOnDuplicateKeyIgnore: "ON DUPLICATE KEY IGNORE",
			StatementTypeOnDuplicateKeyUpdate: "ON DUPLICATE KEY UPDATE",
		}
}
