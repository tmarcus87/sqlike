package statement

import (
	"github.com/stretchr/testify/assert"
	"github.com/tmarcus87/sqlike/dialect"
	"github.com/tmarcus87/sqlike/model"
	"testing"
)

func TestInsertIntoColumns_Accept(t *testing.T) {
	t1 := &model.BasicTable{Name: "t1"}

	tests := []struct {
		name string
		c1   model.Column
		c2   model.Column
	}{
		{
			name: "WithoutAs",
			c1:   &model.BasicColumn{Table: t1, Name: "c1"},
			c2:   &model.BasicColumn{Table: t1, Name: "c2"},
		},
		{
			name: "WithAs",
			c1:   (&model.BasicColumn{Table: t1, Name: "c1"}).SQLikeAs("c1alt"),
			c2:   (&model.BasicColumn{Table: t1, Name: "c2"}).SQLikeAs("c2alt"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			asserts := assert.New(t)

			stmt, _, err :=
				NewInsertIntoBranchStep(root(dialect.MySQL), t1).
					Columns(test.c1, test.c2).
					Values(1, 2).
					Build().
					StatementAndBindings()

			asserts.Nil(err)
			asserts.Equal("INSERT INTO `t1` (`c1`, `c2`) VALUES (?, ?)", stmt)

		})
	}

}

func TestInsertIntoValues_Accept(t *testing.T) {

	t.Run("OneValues", func(t *testing.T) {
		asserts := assert.New(t)

		t1 := &model.BasicTable{Name: "t1"}
		c1 := &model.BasicColumn{Table: t1, Name: "c1"}
		c2 := &model.BasicColumn{Table: t1, Name: "c2"}

		stmt, bindings, err :=
			NewInsertIntoBranchStep(root(dialect.MySQL), t1).
				Columns(c1, c2).
				Values(1, 2).
				Build().
				StatementAndBindings()

		asserts.Nil(err)
		asserts.Equal("INSERT INTO `t1` (`c1`, `c2`) VALUES (?, ?)", stmt)
		asserts.Len(bindings, 2)
		asserts.Equal(1, bindings[0])
		asserts.Equal(2, bindings[1])
	})

	t.Run("TwoValues", func(t *testing.T) {
		asserts := assert.New(t)

		t1 := &model.BasicTable{Name: "t1"}
		c1 := &model.BasicColumn{Table: t1, Name: "c1"}
		c2 := &model.BasicColumn{Table: t1, Name: "c2"}

		stmt, bindings, err :=
			NewInsertIntoBranchStep(root(dialect.MySQL), t1).
				Columns(c1, c2).
				Values(1, 2).
				Values(3, 4).
				Build().
				StatementAndBindings()

		asserts.Nil(err)
		asserts.Equal("INSERT INTO `t1` (`c1`, `c2`) VALUES (?, ?), (?, ?)", stmt)
		asserts.Len(bindings, 4)
		asserts.Equal(1, bindings[0])
		asserts.Equal(2, bindings[1])
		asserts.Equal(3, bindings[2])
		asserts.Equal(4, bindings[3])
	})

	t.Run("WithoutColumns", func(t *testing.T) {
		asserts := assert.New(t)

		t1 := &model.BasicTable{Name: "t1"}

		stmt, bindings, err :=
			NewInsertIntoBranchStep(root(dialect.MySQL), t1).
				Values(1, 2).
				Values(3, 4).
				Build().
				StatementAndBindings()

		asserts.Nil(err)
		asserts.Equal("INSERT INTO `t1` VALUES (?, ?), (?, ?)", stmt)
		asserts.Len(bindings, 4)
		asserts.Equal(1, bindings[0])
		asserts.Equal(2, bindings[1])
		asserts.Equal(3, bindings[2])
		asserts.Equal(4, bindings[3])
	})
}

func TestInsertIntoStructValues_Accept(t *testing.T) {

	t1 := &model.BasicTable{Name: "t1"}
	c1 := &model.BasicColumn{Table: t1, Name: "c1"}
	c2 := &model.BasicColumn{Table: t1, Name: "c2"}

	t.Run("WithoutTag", func(t *testing.T) {
		type ValueStruct struct {
			C1 int
			C2 string
		}

		t.Run("OneWithoutPtr", func(t *testing.T) {
			asserts := assert.New(t)

			stmt, bindings, err :=
				NewInsertIntoBranchStep(root(dialect.MySQL), t1).
					Columns(c1, c2).
					ValueStructs(ValueStruct{C1: 1, C2: "hoge"}).
					Build().
					StatementAndBindings()
			asserts.Nil(err)
			asserts.Equal("INSERT INTO `t1` (`c1`, `c2`) VALUES (?, ?)", stmt)
			asserts.Len(bindings, 2)
			asserts.Equal(1, bindings[0])
			asserts.Equal("hoge", bindings[1])

		})
		t.Run("OneWithPtr", func(t *testing.T) {
			asserts := assert.New(t)

			stmt, bindings, err :=
				NewInsertIntoBranchStep(root(dialect.MySQL), t1).
					Columns(c1, c2).
					ValueStructs(&ValueStruct{C1: 1, C2: "hoge"}).
					Build().
					StatementAndBindings()
			asserts.Nil(err)
			asserts.Equal("INSERT INTO `t1` (`c1`, `c2`) VALUES (?, ?)", stmt)
			asserts.Len(bindings, 2)
			asserts.Equal(1, bindings[0])
			asserts.Equal("hoge", bindings[1])
		})
		t.Run("TwoWithoutPtr", func(t *testing.T) {
			asserts := assert.New(t)

			stmt, bindings, err :=
				NewInsertIntoBranchStep(root(dialect.MySQL), t1).
					Columns(c1, c2).
					ValueStructs(
						ValueStruct{C1: 1, C2: "hoge"},
						ValueStruct{C1: 2, C2: "fuga"}).
					Build().
					StatementAndBindings()
			asserts.Nil(err)
			asserts.Equal("INSERT INTO `t1` (`c1`, `c2`) VALUES (?, ?), (?, ?)", stmt)
			asserts.Len(bindings, 4)
			asserts.Equal(1, bindings[0])
			asserts.Equal("hoge", bindings[1])
			asserts.Equal(2, bindings[2])
			asserts.Equal("fuga", bindings[3])
		})
		t.Run("TwoWithPtr", func(t *testing.T) {
			asserts := assert.New(t)

			stmt, bindings, err :=
				NewInsertIntoBranchStep(root(dialect.MySQL), t1).
					Columns(c1, c2).
					ValueStructs(
						&ValueStruct{C1: 1, C2: "hoge"},
						&ValueStruct{C1: 2, C2: "fuga"}).
					Build().
					StatementAndBindings()
			asserts.Nil(err)
			asserts.Equal("INSERT INTO `t1` (`c1`, `c2`) VALUES (?, ?), (?, ?)", stmt)
			asserts.Len(bindings, 4)
			asserts.Equal(1, bindings[0])
			asserts.Equal("hoge", bindings[1])
			asserts.Equal(2, bindings[2])
			asserts.Equal("fuga", bindings[3])
		})

	})

	t.Run("WithTag", func(t *testing.T) {
		type TaggedValueStruct struct {
			Column1 int    `sqlike:"c1"`
			Column2 string `sqlike:"c2"`
		}

		t.Run("OneWithoutPtr", func(t *testing.T) {
			asserts := assert.New(t)

			stmt, bindings, err :=
				NewInsertIntoBranchStep(root(dialect.MySQL), t1).
					Columns(c1, c2).
					ValueStructs(TaggedValueStruct{Column1: 1, Column2: "hoge"}).
					Build().
					StatementAndBindings()
			asserts.Nil(err)
			asserts.Equal("INSERT INTO `t1` (`c1`, `c2`) VALUES (?, ?)", stmt)
			asserts.Len(bindings, 2)
			asserts.Equal(1, bindings[0])
			asserts.Equal("hoge", bindings[1])

		})
		t.Run("OneWithPtr", func(t *testing.T) {
			asserts := assert.New(t)

			stmt, bindings, err :=
				NewInsertIntoBranchStep(root(dialect.MySQL), t1).
					Columns(c1, c2).
					ValueStructs(&TaggedValueStruct{Column1: 1, Column2: "hoge"}).
					Build().
					StatementAndBindings()
			asserts.Nil(err)
			asserts.Equal("INSERT INTO `t1` (`c1`, `c2`) VALUES (?, ?)", stmt)
			asserts.Len(bindings, 2)
			asserts.Equal(1, bindings[0])
			asserts.Equal("hoge", bindings[1])
		})
		t.Run("TwoWithoutPtr", func(t *testing.T) {
			asserts := assert.New(t)

			stmt, bindings, err :=
				NewInsertIntoBranchStep(root(dialect.MySQL), t1).
					Columns(c1, c2).
					ValueStructs(
						TaggedValueStruct{Column1: 1, Column2: "hoge"},
						TaggedValueStruct{Column1: 2, Column2: "fuga"}).
					Build().
					StatementAndBindings()
			asserts.Nil(err)
			asserts.Equal("INSERT INTO `t1` (`c1`, `c2`) VALUES (?, ?), (?, ?)", stmt)
			asserts.Len(bindings, 4)
			asserts.Equal(1, bindings[0])
			asserts.Equal("hoge", bindings[1])
			asserts.Equal(2, bindings[2])
			asserts.Equal("fuga", bindings[3])
		})
		t.Run("TwoWithPtr", func(t *testing.T) {
			asserts := assert.New(t)

			stmt, bindings, err :=
				NewInsertIntoBranchStep(root(dialect.MySQL), t1).
					Columns(c1, c2).
					ValueStructs(
						&TaggedValueStruct{Column1: 1, Column2: "hoge"},
						&TaggedValueStruct{Column1: 2, Column2: "fuga"}).
					Build().
					StatementAndBindings()
			asserts.Nil(err)
			asserts.Equal("INSERT INTO `t1` (`c1`, `c2`) VALUES (?, ?), (?, ?)", stmt)
			asserts.Len(bindings, 4)
			asserts.Equal(1, bindings[0])
			asserts.Equal("hoge", bindings[1])
			asserts.Equal(2, bindings[2])
			asserts.Equal("fuga", bindings[3])
		})
	})
}

func TestInsertIntoSelect_Accept(t *testing.T) {
	t.Run("WithColumns", func(t *testing.T) {
		asserts := assert.New(t)

		t1 := &model.BasicTable{Name: "t1"}
		t2 := &model.BasicTable{Name: "t2"}
		c1 := &model.BasicColumn{Table: t1, Name: "c1"}
		c2 := &model.BasicColumn{Table: t1, Name: "c2"}
		c3 := &model.BasicColumn{Table: t2, Name: "c3"}
		c4 := &model.BasicColumn{Table: t2, Name: "c4"}

		stmt, _, err :=
			NewInsertIntoBranchStep(root(dialect.MySQL), t1).
				Columns(c1, c2).
				Select(c3, c4).
				From(t2).
				Build().
				StatementAndBindings()
		asserts.Nil(err)
		asserts.Equal("INSERT INTO `t1` (`c1`, `c2`) SELECT `t2`.`c3`, `t2`.`c4` FROM `t2`", stmt)
	})

	t.Run("WithoutColumns", func(t *testing.T) {
		asserts := assert.New(t)

		t1 := &model.BasicTable{Name: "t1"}
		t2 := &model.BasicTable{Name: "t2"}
		c3 := &model.BasicColumn{Table: t2, Name: "c3"}
		c4 := &model.BasicColumn{Table: t2, Name: "c4"}

		stmt, _, err :=
			NewInsertIntoBranchStep(root(dialect.MySQL), t1).
				Select(c3, c4).
				From(t2).
				Build().
				StatementAndBindings()
		asserts.Nil(err)
		asserts.Equal("INSERT INTO `t1` SELECT `t2`.`c3`, `t2`.`c4` FROM `t2`", stmt)
	})

	t.Run("WithSelectAs", func(t *testing.T) {
		asserts := assert.New(t)

		t1 := &model.BasicTable{Name: "t1"}
		t2 := &model.BasicTable{Name: "t2"}
		c1 := &model.BasicColumn{Table: t2, Name: "c1"}
		c2 := &model.BasicColumn{Table: t2, Name: "c2"}

		stmt, _, err :=
			NewInsertIntoBranchStep(root(dialect.MySQL), t1).
				Select(c1.SQLikeAs("c1alt"), c2.SQLikeAs("c2alt")).
				From(t2.SQLikeAs("t2alt")).
				Build().
				StatementAndBindings()
		asserts.Nil(err)
		asserts.Equal("INSERT INTO `t1` SELECT `t2alt`.`c1` AS `c1alt`, `t2alt`.`c2` AS `c2alt` FROM `t2` AS `t2alt`", stmt)
	})

}
