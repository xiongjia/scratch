package bulbasaur

import (
	"github.com/prometheus/prometheus/scrape"
	"github.com/prometheus/prometheus/storage"
)

type Engine struct {
	storage *LocalStorageEngine
}

func NewEngine() (*Engine, error) {
	storageOpt := &LocalStorageOptions{Dir: "c:/wrk/tmp/tsdb1"}
	storageEng, err := MakeLocalStorageEngine(storageOpt)
	if err != nil {
		return nil, err
	}
	fanoutStorage := storage.NewFanout(nil, storageEng.db)
	scrape.NewManager(nil, nil, fanoutStorage, nil)
	return &Engine{storage: storageEng}, nil
}
