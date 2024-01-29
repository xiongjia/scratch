package bulbasaur_test

import (
	"fmt"
	"stray/pkg/bulbasaur"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorage(t *testing.T) {
	fmt.Printf("tmp dir = %s\n", t.TempDir())

	storageOpt := &bulbasaur.LocalStorageOptions{Dir: "c:/wrk/tmp/tsdb1"}
	_, err := bulbasaur.MakeLocalStorageEngine(storageOpt)
	assert.NoError(t, err)

	// Test data appending
}
