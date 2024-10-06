package collection_test

import (
	"stray/pkg/collection"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

func TestFind(t *testing.T) {
	defer goleak.VerifyNone(t)

	_, ok := collection.First([]int{})
	assert.False(t, ok)

	num, ok := collection.First([]int{1, 2, 3})
	assert.True(t, ok)
	assert.Equal(t, num, 1)

	msg, ok := collection.First([]string{"1", "2", "3"})
	assert.True(t, ok)
	assert.Equal(t, msg, "1")
}

func TestEmpty(t *testing.T) {
	defer goleak.VerifyNone(t)

	type (
		Mock struct {
			Num int `default:"123"`
		}
	)

	data := collection.FirstOrEmpty([]Mock{})
	assert.Equal(t, data.Num, 123)
	data = collection.FirstOrEmpty([]Mock{{Num: 2}})
	assert.Equal(t, data.Num, 2)

	data2 := collection.FirstOrEmpty([]*Mock{})
	assert.Nil(t, data2)

	data2 = collection.FirstOrEmpty([]*Mock{{Num: 3}})
	assert.NotNil(t, data2)
	assert.Equal(t, data2.Num, 3)
}
