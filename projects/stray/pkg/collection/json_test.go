package collection_test

import (
	"stray/pkg/collection"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonParser(t *testing.T) {
	type (
		mockObj struct {
			Name string
		}
	)

	jsonb, err := collection.JsonStringFromObject[*mockObj](&mockObj{Name: "name1"})
	assert.NoError(t, err)
	assert.NotEmpty(t, jsonb)
	input := strings.NewReader(string(jsonb))
	t.Logf("mock input json %s", string(jsonb))

	output, err := collection.JsonReaderToObject[mockObj](input)
	assert.NoError(t, err)
	t.Logf("output object = %+v", output)
	assert.Equal(t, output.Name, "name1")
}
