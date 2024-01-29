package bulbasaur

import (
	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/tsdb"
)

type storageEngine struct {
	stats *tsdb.DBStats
	db    storage.Storage
}

type Engine struct {
	storage *storageEngine
}

func makeStorageEngine() (*storageEngine, error) {
	dbStates := tsdb.NewDBStats()
	db, err := tsdb.Open("c:/wrk/tmp/tsdb", nil, nil, nil, dbStates)
	if err != nil {
		return nil, err
	}

	return &storageEngine{
		stats: dbStates,
		db:    db,
	}, nil

}

func NewEngine() (*Engine, error) {
	storageEng, err := makeStorageEngine()
	if err != nil {
		return nil, err
	}
	return &Engine{storage: storageEng}, nil
}
