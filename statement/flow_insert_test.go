package statement

import (
	"github.com/stretchr/testify/assert"
	"github.com/tmarcus87/sqlike/dialect"
	"github.com/tmarcus87/sqlike/model"
	"testing"
)

func TestInsertIntoBranchStep(t *testing.T) {
	t1 := model.NewTable("t1")
	c1 := model.NewInt32Column(t1, "c1")
	c2 := model.NewInt32Column(t1, "c2")

	t.Run("InsertOnDuplicateKeyUpdateSet", func(t *testing.T) {
		stmt, bindings, err :=
			NewInsertIntoBranchStep(root(dialect.MySQL), t1).
				Columns(c1, c2).
				Values(1, 2).
				OnDuplicateKeyUpdate().
				SetValue(c2.Value(3)).
				Build().
				StatementAndBindings()

		asserts := assert.New(t)
		asserts.Nil(err)
		asserts.Equal("INSERT INTO `t1` (`c1`, `c2`) VALUES (?, ?) ON DUPLICATE KEY UPDATE SET `c2` = ?", stmt)
		asserts.Len(bindings, 3)
		asserts.Equal(1, bindings[0])
		asserts.Equal(2, bindings[1])
		asserts.Equal(int32(3), bindings[2])
	})

	t.Run("InsertOnDuplicateKeyUpdateSetRecord", func(t *testing.T) {
		type Value struct {
			C1 int32 `sqlike:"c1"`
			C2 int32 `sqlike:"c2"`
		}

		stmt, bindings, err :=
			NewInsertIntoBranchStep(root(dialect.MySQL), t1).
				Columns(c1, c2).
				Values(1, 2).
				OnDuplicateKeyUpdate().
				SetRecord(&model.Record{
					Value: &Value{
						C1: 3,
						C2: 4,
					},
					Skip: []model.Column{c1},
				}).
				Build().
				StatementAndBindings()

		asserts := assert.New(t)
		asserts.Nil(err)
		asserts.Equal("INSERT INTO `t1` (`c1`, `c2`) VALUES (?, ?) ON DUPLICATE KEY UPDATE SET `c2` = ?", stmt)
		asserts.Len(bindings, 3)
		asserts.Equal(1, bindings[0])
		asserts.Equal(2, bindings[1])
		asserts.Equal(int32(4), bindings[2])
	})

}
