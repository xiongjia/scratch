package util_test

import (
	"stray/pkg/util"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

func TestJsonParser(t *testing.T) {
	defer goleak.VerifyNone(t)

	type (
		mockObj struct {
			Name string
		}
	)

	jsonb, err := util.JsonStringFromObject[*mockObj](&mockObj{Name: "name1"})
	assert.NoError(t, err)
	assert.NotEmpty(t, jsonb)
	input := strings.NewReader(string(jsonb))
	t.Logf("mock input json %s", string(jsonb))

	output, err := util.JsonReaderToObject[mockObj](input)
	assert.NoError(t, err)
	t.Logf("output object = %+v", output)
	assert.Equal(t, output.Name, "name1")
}
