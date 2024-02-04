package bulbasaur

import (
	"log/slog"

	"github.com/prometheus/prometheus/scrape"
	"github.com/prometheus/prometheus/storage"
)

type EngineOptions struct {
	Logger *slog.Logger
}

type Engine struct {
	logger  *slog.Logger
	storage *LocalStorageEngine
}

func NewEngine(opts EngineOptions) (*Engine, error) {
	storageOpt := &LocalStorageOptions{
		Dir:    "c:/wrk/tmp/tsdb1",
		Logger: opts.Logger,
	}
	storageEng, err := MakeLocalStorageEngine(storageOpt)
	if err != nil {
		return nil, err
	}
	fanoutStorage := storage.NewFanout(nil, storageEng.db)
	scrape.NewManager(nil, nil, fanoutStorage, nil)
	return &Engine{
		logger:  opts.Logger,
		storage: storageEng,
	}, nil
}
