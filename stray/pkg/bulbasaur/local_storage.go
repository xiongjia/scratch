package bulbasaur

import (
	"log/slog"
	"stray/pkg/dugtrio"

	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/tsdb"
)

type LocalStorageEngine struct {
	logger *slog.Logger
	stats  *tsdb.DBStats
	db     storage.Storage
}

func (e *LocalStorageEngine) Test() {
	e.logger.Debug("test")

	// appender := e.db.Appender(context.Background())

}

type LocalStorageOptions struct {
	Logger *slog.Logger
	Dir    string
}

func MakeLocalStorageEngine(opts *LocalStorageOptions) (*LocalStorageEngine, error) {
	stats := tsdb.NewDBStats()
	db, err := tsdb.Open(opts.Dir,
		dugtrio.NewKitLoggerAdapterSlog(opts.Logger), nil, nil, stats)
	if err != nil {
		return nil, err
	}
	return &LocalStorageEngine{
		logger: opts.Logger,
		stats:  stats,
		db:     db,
	}, nil
}
