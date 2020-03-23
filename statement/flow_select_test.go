package statement

import (
	"github.com/stretchr/testify/assert"
	"github.com/tmarcus87/sqlike/dialect"
	"github.com/tmarcus87/sqlike/model"
	"testing"
)

func TestExplainSelectBranchStep(t *testing.T) {
	stmt, _, err :=
		NewExplainSelectBranchStep(root(dialect.MySQL)).
			SelectOne().
			Build().
			StatementAndBindings()

	asserts := assert.New(t)

	asserts.Nil(err)
	asserts.Equal("EXPLAIN SELECT 1 FROM dual", stmt)
}

func TestSelectFromBranchStep(t *testing.T) {
	t1 := model.NewTable("t1")

	stmt, _, err :=
		NewSelectFromBranchStep(root(dialect.MySQL), t1).
			Build().
			StatementAndBindings()

	asserts := assert.New(t)
	asserts.Nil(err)
	asserts.Equal("SELECT * FROM `t1`", stmt)
}
