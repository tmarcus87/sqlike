package statement

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFieldValueMap(t *testing.T) {
	type Value struct {
		Key1 string
		Key2 int `sqlike:"key2alt"`
	}

	v := Value{
		Key1: "key1",
		Key2: 2,
	}

	t.Run("WithNonPtr", func(t *testing.T) {
		fvm, err := getColumnName2FieldValueMap(v)

		asserts := assert.New(t)
		asserts.Nil(err)
		asserts.Contains(fvm, "key1")
		asserts.NotContains(fvm, "key2")
		asserts.Contains(fvm, "key2alt")
		asserts.Equal("key1", fvm["key1"].Interface())
		asserts.Equal(2, fvm["key2alt"].Interface())
	})

	t.Run("WithPtr", func(t *testing.T) {
		fvm, err := getColumnName2FieldValueMap(&v)

		asserts := assert.New(t)
		asserts.Nil(err)
		asserts.Contains(fvm, "key1")
		asserts.NotContains(fvm, "key2")
		asserts.Contains(fvm, "key2alt")
		asserts.Equal("key1", fvm["key1"].Interface())
		asserts.Equal(2, fvm["key2alt"].Interface())
	})
}
