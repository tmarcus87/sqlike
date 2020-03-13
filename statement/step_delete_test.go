package statement

import (
	"github.com/stretchr/testify/assert"
	"github.com/tmarcus87/sqlike/dialect"
	"github.com/tmarcus87/sqlike/model"
	"testing"
)

func TestDeleteFromStep_Accept(t *testing.T) {
	t1 := model.NewTable("t1")
	c1 := model.NewBoolColumn(t1, "c1")

	t.Run("WithoutWhere", func(t *testing.T) {
		stmt, _, err :=
			NewDeleteFromBranchStep(root(dialect.MySQL), t1).
				Build().
				StatementAndBindings()

		asserts := assert.New(t)
		asserts.Nil(err)
		asserts.Equal("DELETE FROM `t1`", stmt)
	})

	t.Run("WithWhere", func(t *testing.T) {
		stmt, bindings, err :=
			NewDeleteFromBranchStep(root(dialect.MySQL), t1).
				Where(c1.CondEq(true)).
				Build().
				StatementAndBindings()

		asserts := assert.New(t)
		asserts.Nil(err)
		asserts.Equal("DELETE FROM `t1` WHERE `t1`.`c1` = ?", stmt)
		asserts.Len(bindings, 1)
		asserts.Equal(true, bindings[0])
	})
}
