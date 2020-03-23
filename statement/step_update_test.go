package statement

import (
	"github.com/stretchr/testify/assert"
	"github.com/tmarcus87/sqlike/dialect"
	"github.com/tmarcus87/sqlike/model"
	"testing"
)

func TestUpdateStep_Accept(t *testing.T) {

}

func TestUpdateSetStep_Accept(t *testing.T) {

	t1 := model.NewTable("t1")

	c1 := model.NewBoolColumn(t1, "c1")
	c2 := model.NewBoolColumn(t1, "c2")

	t.Run("OneSet", func(t *testing.T) {
		stmt, bindings, err :=
			NewUpdateBranchStep(root(dialect.MySQL), t1).
				Set(c1, 1).
				Build().
				StatementAndBindings()

		asserts := assert.New(t)
		asserts.Nil(err)
		asserts.Equal("UPDATE `t1` SET `c1` = ?", stmt)
		asserts.Len(bindings, 1)
		asserts.Equal(1, bindings[0])
	})

	t.Run("TwoSet", func(t *testing.T) {
		stmt, bindings, err :=
			NewUpdateBranchStep(root(dialect.MySQL), t1).
				Set(c1, 1).
				Set(c2, 2).
				Build().
				StatementAndBindings()

		asserts := assert.New(t)
		asserts.Nil(err)
		asserts.Equal("UPDATE `t1` SET `c1` = ?, `c2` = ?", stmt)
		asserts.Len(bindings, 2)
		asserts.Equal(1, bindings[0])
		asserts.Equal(2, bindings[1])
	})

	t.Run("WithWhere", func(t *testing.T) {
		stmt, bindings, err :=
			NewUpdateBranchStep(root(dialect.MySQL), t1).
				Set(c1, 1).
				Set(c2, 2).
				Where(c1.Eq(true)).
				Build().
				StatementAndBindings()

		asserts := assert.New(t)
		asserts.Nil(err)
		asserts.Equal("UPDATE `t1` SET `c1` = ?, `c2` = ? WHERE `t1`.`c1` = ?", stmt)
		asserts.Len(bindings, 3)
		asserts.Equal(1, bindings[0])
		asserts.Equal(2, bindings[1])
		asserts.Equal(true, bindings[2])
	})
}

func TestUpdateSetRecordStep_Accept(t *testing.T) {

	t1 := model.NewTable("t1")

	c1 := model.NewBoolColumn(t1, "c1")
	c2 := model.NewBoolColumn(t1, "c2")
	c3 := model.NewBoolColumn(t1, "c3")

	type Value struct {
		C1      int
		Column2 int `sqlike:"c2"`
		Column3 int `sqlike:"c3"`
	}

	t.Run("All", func(t *testing.T) {
		stmt, bindings, err :=
			NewUpdateBranchStep(root(dialect.MySQL), t1).
				SetRecord(&model.Record{Value: &Value{C1: 1, Column2: 2, Column3: 3}}).
				Where(c1.Eq(true)).
				Build().
				StatementAndBindings()
		asserts := assert.New(t)
		asserts.Nil(err)
		asserts.Equal("UPDATE `t1` SET `c1` = ?, `c2` = ?, `c3` = ? WHERE `t1`.`c1` = ?", stmt)
		asserts.Len(bindings, 4)
		asserts.Equal(1, bindings[0])
		asserts.Equal(2, bindings[1])
		asserts.Equal(3, bindings[2])
		asserts.Equal(true, bindings[3])
	})

	t.Run("WithOnly", func(t *testing.T) {
		stmt, bindings, err :=
			NewUpdateBranchStep(root(dialect.MySQL), t1).
				SetRecord(&model.Record{
					Only:  []model.Column{c1, c2},
					Value: &Value{C1: 1, Column2: 2},
				}).
				Where(c1.Eq(true)).
				Build().
				StatementAndBindings()
		asserts := assert.New(t)
		asserts.Nil(err)
		asserts.Equal("UPDATE `t1` SET `c1` = ?, `c2` = ? WHERE `t1`.`c1` = ?", stmt)
		asserts.Len(bindings, 3)
		asserts.Equal(1, bindings[0])
		asserts.Equal(2, bindings[1])
		asserts.Equal(true, bindings[2])
	})

	t.Run("WithSkip", func(t *testing.T) {
		stmt, bindings, err :=
			NewUpdateBranchStep(root(dialect.MySQL), t1).
				SetRecord(&model.Record{
					Skip:  []model.Column{c3},
					Value: &Value{C1: 1, Column2: 2},
				}).
				Where(c1.Eq(true)).
				Build().
				StatementAndBindings()
		asserts := assert.New(t)
		asserts.Nil(err)
		asserts.Equal("UPDATE `t1` SET `c1` = ?, `c2` = ? WHERE `t1`.`c1` = ?", stmt)
		asserts.Len(bindings, 3)
		asserts.Equal(1, bindings[0])
		asserts.Equal(2, bindings[1])
		asserts.Equal(true, bindings[2])

	})
}
