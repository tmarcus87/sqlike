package dialect

const Sqlite3 = "seqlite3"

func init() {
	sqlDialect[Sqlite3] =
		map[StatementType]string{
			StatementTypeSelectOne: "SELECT 1",
		}
}
