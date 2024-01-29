package bulbasaur

import (
	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/tsdb"
)

type LocalStorageEngine struct {
	stats *tsdb.DBStats
	db    storage.Storage
}

type LocalStorageOptions struct {
	Dir string
}

func MakeLocalStorageEngine(opts *LocalStorageOptions) (*LocalStorageEngine, error) {
	stats := tsdb.NewDBStats()
	db, err := tsdb.Open(opts.Dir, makeLogger(), nil, nil, stats)
	if err != nil {
		return nil, err
	}
	return &LocalStorageEngine{
		stats: stats,
		db:    db,
	}, nil
}
