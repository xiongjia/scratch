package prom

import (
	"context"
	"database/sql"
	"sync"

	_ "github.com/mattn/go-sqlite3"
	"github.com/prometheus/prometheus/config"
	"github.com/prometheus/prometheus/model/exemplar"
	"github.com/prometheus/prometheus/model/histogram"
	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/model/metadata"
	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/tsdb"

	kitlog "github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

type (
	storageDb struct {
		log kitlog.Logger

		db    *sql.DB
		rwMtx sync.RWMutex
	}

	rcdLab struct {
		Id       uint64
		Lab      labels.Labels
		Metadata *metadata.Metadata
	}

	rcdLabs  []rcdLab
	rcdValue struct {
		Id    uint64
		LabId uint64
	}

	tsdbHdr struct {
		labs map[uint64]rcdLabs
	}

	dbAppender struct {
		log kitlog.Logger
		ctx context.Context
		db  *storageDb
	}
)

func makeStorageDb(log kitlog.Logger) (*storageDb, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}
	// Init DB schema

	_, err = db.Exec(`
		create table label (id integer not null primary key, labs text, metadata text);

		create table rcd (id integer not null primary key, ts integer not null, val REAL );
	`)
	if err != nil {
		db.Close()
		return nil, err
	}
	return &storageDb{log: log, db: db}, nil
}

func createDbStorage(l kitlog.Logger) (storage.Storage, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		_ = level.Debug(l).Log("msg", "WALReplayStatus")
		return nil, err
	}

	return &storageDb{log: l, db: db}, nil
}

func (s *storageDb) Querier(ctx context.Context, mint, maxt int64) (storage.Querier, error) {
	_ = level.Debug(s.log).Log("msg", "Querier", "mint", mint, "maxt", maxt)
	// TODO
	return nil, tsdb.ErrNotReady
}

func (s *storageDb) ChunkQuerier(ctx context.Context, mint, maxt int64) (storage.ChunkQuerier, error) {
	_ = level.Debug(s.log).Log("msg", "ChunkQuerier", "mint", mint, "maxt", maxt)
	// TODO
	return nil, tsdb.ErrNotReady
}

func (s *storageDb) Appender(ctx context.Context) storage.Appender {
	_ = level.Debug(s.log).Log("msg", "Appender")
	return &dbAppender{log: s.log, db: s, ctx: ctx}
}

func (s *storageDb) ApplyConfig(conf *config.Config) error {
	_ = level.Debug(s.log).Log("msg", "ApplyConfig", "conf", conf)
	return nil
}

func (s *storageDb) ExemplarQuerier(ctx context.Context) (storage.ExemplarQuerier, error) {
	_ = level.Debug(s.log).Log("msg", "ExemplarQuerier")
	return nil, tsdb.ErrNotReady
}

func (s *storageDb) Close() error {
	_ = level.Debug(s.log).Log("msg", "Close")
	return nil
}

func (s *storageDb) CleanTombstones() error {
	_ = level.Debug(s.log).Log("msg", "CleanTombstones")
	return nil

}

func (s *storageDb) Delete(mint, maxt int64, ms ...*labels.Matcher) error {
	_ = level.Debug(s.log).Log("msg", "delete", "mint", mint, "maxt", maxt, "ms", ms)
	return tsdb.ErrNotReady
}

func (s *storageDb) Snapshot(dir string, withHead bool) error {
	_ = level.Debug(s.log).Log("msg", "Snapshot", "dir", dir, "withHead", withHead)
	return nil
}

func (s *storageDb) Stats(statsByLabelName string, limit int) (*tsdb.Stats, error) {
	_ = level.Debug(s.log).Log("msg", "Snapshot", "statsByLabelName", statsByLabelName, "limit", limit)
	return &tsdb.Stats{}, nil
}

func (s *storageDb) WALReplayStatus() (tsdb.WALReplayStatus, error) {
	_ = level.Debug(s.log).Log("msg", "WALReplayStatus")
	return tsdb.WALReplayStatus{}, nil
}

func (s *storageDb) StartTime() (int64, error) {
	_ = level.Debug(s.log).Log("msg", "StartTime")
	return 0, nil
}

func (db *dbAppender) Append(ref storage.SeriesRef, l labels.Labels, t int64, v float64) (storage.SeriesRef, error) {
	// TODO
	return 0, nil
}

func (db *dbAppender) UpdateMetadata(ref storage.SeriesRef, l labels.Labels, m metadata.Metadata) (storage.SeriesRef, error) {
	return 0, nil
}

func (db *dbAppender) Commit() error {
	// Don't support commit & rollback
	_ = level.Debug(db.log).Log("msg", "db appender commit")
	return nil
}

func (db *dbAppender) Rollback() error {
	// Don't support commit & rollback
	_ = level.Debug(db.log).Log("msg", "db appender rollback")
	return nil
}

func (dbAppender) AppendExemplar(ref storage.SeriesRef, l labels.Labels, v exemplar.Exemplar) (storage.SeriesRef, error) {
	// TODO
	return 0, nil
}

func (dbAppender) AppendHistogram(storage.SeriesRef, labels.Labels, int64, *histogram.Histogram, *histogram.FloatHistogram) (storage.SeriesRef, error) {
	// TODO
	return 0, nil
}
