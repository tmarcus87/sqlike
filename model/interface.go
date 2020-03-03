package model

type Table interface {
	// SQLikeTableName テーブル名を返します
	SQLikeTableName() string
	// SQLikeTableAliasName エイリアス名を返します
	SQLikeTableAlias() string
	// SQLikeAs テーブルにエイリアスを設定します
	SQLikeAs(alias string) Table
}

type Column interface {
	// SQLikeTable
	SQLikeTable() Table
	// SQLikeColumnOriginalName
	SQLikeColumnName() string
	// SQLikeColumnAliasName
	SQLikeColumnAlias() string
	// SQLikeAs
	SQLikeAs(alias string) Column
	// SQLikeModifier
	SQLikeSelectModFmt() string
}

type Condition interface {
	Apply(stmt *string, bindings *[]interface{})
}
