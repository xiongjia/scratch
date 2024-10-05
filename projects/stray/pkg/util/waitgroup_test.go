package util_test

import (
	"testing"
	"time"

	"stray/pkg/util"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

func TestWg(t *testing.T) {
	defer goleak.VerifyNone(t)

	wg := util.NewWaitGroup()

	var testData int
	testData = 1
	wg.Go(func() {
		t.Log("Test1 is running")
		time.Sleep(2 * time.Second)
		testData = 2
		t.Log("Test2 finished")
	})

	var testData2 int
	testData2 = 1
	wg.Go(func() {
		t.Log("Test2 is running")
		time.Sleep(1 * time.Second)
		testData2 = 2
		t.Log("Test1 finished")
	})

	wg.Wait()
	t.Log("All tests finished")
	assert.Equal(t, testData, 2)
	assert.Equal(t, testData2, 2)
}
