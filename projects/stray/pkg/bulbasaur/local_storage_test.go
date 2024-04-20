package bulbasaur_test

import (
	"log/slog"
	"stray/pkg/bulbasaur"
	"stray/pkg/dugtrio"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorage(t *testing.T) {
	logger := dugtrio.NewSlogForTesting(t, dugtrio.SLogBaseOptions{Level: slog.LevelDebug})
	storageOpt := &bulbasaur.LocalStorageOptions{
		Dir:    "c:/wrk/tmp/tsdb1",
		Logger: logger,
	}
	eng, err := bulbasaur.MakeLocalStorageEngine(storageOpt)
	assert.NoError(t, err)

	eng.Test()
}
