package metric

import (
	"context"
	"errors"
	"sync"
	"time"

	kitlog "github.com/go-kit/log"
	"github.com/go-kit/log/level"

	"github.com/prometheus/prometheus/config"
	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/tsdb"
)

type (
	PromStorageOpts struct {
		Log        kitlog.Logger
		FsTsdbPath string
	}

	PromStorage struct {
		log kitlog.Logger
		mtx sync.RWMutex

		storageOpts PromStorageOpts

		stats           *tsdb.DBStats
		db              storage.Storage
		startTimeMargin int64
	}
)

func NewPromStorage(opts PromStorageOpts) (*PromStorage, error) {
	log := kitlog.With(opts.Log, LOG_COMPONENT_KEY, COMPONENT_STORAGE)
	dbStats := tsdb.NewDBStats()
	db, err := createFsStorage(opts.FsTsdbPath, dbStats, log)
	if err != nil {
		return nil, err
	}
	return &PromStorage{storageOpts: opts, stats: dbStats, db: db, log: log}, nil
}

func (s *PromStorage) get() storage.Storage {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.db
}

func (s *PromStorage) getStats() *tsdb.DBStats {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.stats
}

func (s *PromStorage) Set(db storage.Storage, startTimeMargin int64) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.db = db
	s.startTimeMargin = startTimeMargin
}

func (s *PromStorage) StartTime() (int64, error) {
	_ = level.Debug(s.log).Log("msg", "StartTime")
	if x := s.get(); x != nil {
		switch db := x.(type) {
		case *tsdb.DB:
			var startTime int64
			if len(db.Blocks()) > 0 {
				startTime = db.Blocks()[0].Meta().MinTime
			} else {
				startTime = time.Now().Unix() * 1000
			}
			return startTime + s.startTimeMargin, nil
		}
	}
	return 0, errors.New("TSDB not implement")

}

func (s *PromStorage) Querier(ctx context.Context, mint, maxt int64) (storage.Querier, error) {
	_ = level.Debug(s.log).Log("msg", "Querier", "mint", mint, "maxt", maxt)
	if x := s.get(); x != nil {
		q, err := x.Querier(ctx, mint, maxt)
		if err != nil {
			return nil, err
		}
		return makeStorageQuerierAdapter(s.log, q), nil
	}
	return nil, tsdb.ErrNotReady
}

func (s *PromStorage) ChunkQuerier(ctx context.Context, mint, maxt int64) (storage.ChunkQuerier, error) {
	_ = level.Debug(s.log).Log("msg", "ChunkQuerier", "mint", mint, "maxt", maxt)
	if x := s.get(); x != nil {
		return x.ChunkQuerier(ctx, mint, maxt)
	}
	return nil, tsdb.ErrNotReady
}

func (s *PromStorage) Appender(ctx context.Context) storage.Appender {
	_ = level.Debug(s.log).Log("msg", "Appender")
	if x := s.get(); x != nil {
		return NewFanoutAppender(x.Appender(ctx), s.log)
	}
	return notReadyAppend{}
}

func (s *PromStorage) ApplyConfig(conf *config.Config) error {
	_ = level.Debug(s.log).Log("msg", "ApplyConfig", "conf", conf)
	db := s.get()
	if db, ok := db.(*tsdb.DB); ok {
		return db.ApplyConfig(conf)
	}
	return nil
}

func (s *PromStorage) ExemplarQuerier(ctx context.Context) (storage.ExemplarQuerier, error) {
	_ = level.Debug(s.log).Log("msg", "ExemplarQuerier")
	if x := s.get(); x != nil {
		switch db := x.(type) {
		case *tsdb.DB:
			return db.ExemplarQuerier(ctx)
		default:
			return nil, errors.New("TSDB not implement")
		}
	}
	return nil, tsdb.ErrNotReady
}

func (s *PromStorage) Close() error {
	_ = level.Debug(s.log).Log("msg", "Close")
	if x := s.get(); x != nil {
		return x.Close()
	}
	return nil
}

func (s *PromStorage) CleanTombstones() error {
	_ = level.Debug(s.log).Log("msg", "CleanTombstones")
	if x := s.get(); x != nil {
		switch db := x.(type) {
		case *tsdb.DB:
			return db.CleanTombstones()
		default:
			return errors.New("TSDB not implement")
		}
	}
	return tsdb.ErrNotReady

}

func (s *PromStorage) Delete(mint, maxt int64, ms ...*labels.Matcher) error {
	_ = level.Debug(s.log).Log("msg", "delete", "mint", mint, "maxt", maxt, "ms", ms)
	if x := s.get(); x != nil {
		switch db := x.(type) {
		case *tsdb.DB:
			return db.Delete(mint, maxt, ms...)
		default:
			return errors.New("TSDB not implement")
		}
	}
	return tsdb.ErrNotReady
}

func (s *PromStorage) Snapshot(dir string, withHead bool) error {
	_ = level.Debug(s.log).Log("msg", "Snapshot", "dir", dir, "withHead", withHead)
	if x := s.get(); x != nil {
		switch db := x.(type) {
		case *tsdb.DB:
			return db.Snapshot(dir, withHead)
		default:
			return errors.New("TSDB not implement")
		}
	}
	return tsdb.ErrNotReady
}

func (s *PromStorage) Stats(statsByLabelName string, limit int) (*tsdb.Stats, error) {
	_ = level.Debug(s.log).Log("msg", "Snapshot", "statsByLabelName", statsByLabelName, "limit", limit)
	if x := s.get(); x != nil {
		switch db := x.(type) {
		case *tsdb.DB:
			return db.Head().Stats(statsByLabelName, limit), nil
		default:
			return nil, errors.New("TSDB not implement")
		}
	}
	return nil, tsdb.ErrNotReady
}

func (s *PromStorage) WALReplayStatus() (tsdb.WALReplayStatus, error) {
	_ = level.Debug(s.log).Log("msg", "WALReplayStatus")
	if x := s.getStats(); x != nil {
		return x.Head.WALReplayStatus.GetWALReplayStatus(), nil
	}
	return tsdb.WALReplayStatus{}, tsdb.ErrNotReady
}
