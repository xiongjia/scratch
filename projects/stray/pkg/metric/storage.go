package metric

import (
	"context"
	"errors"
	"math"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/prometheus/config"
	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/tsdb"
)

type (
	PromStorage struct {
		mtx             sync.RWMutex
		stats           *tsdb.DBStats
		startTimeMargin int64
		db              storage.Storage
	}
)

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

func NewPromStorage(dbPath string) (*PromStorage, error) {
	dbStats := tsdb.NewDBStats()
	db, err := OpenLocalTsdb(dbPath, prometheus.DefaultRegisterer, tsdb.DefaultOptions(), dbStats)
	if err != nil {
		return nil, err
	}

	return &PromStorage{
		stats: dbStats,
		db:    db,
	}, nil
}

func (s *PromStorage) Set(db storage.Storage, startTimeMargin int64) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.db = db
	s.startTimeMargin = startTimeMargin
}

func (s *PromStorage) StartTime() (int64, error) {
	if x := s.get(); x != nil {
		switch db := x.(type) {
		case *tsdb.DB:
			var startTime int64
			if len(db.Blocks()) > 0 {
				startTime = db.Blocks()[0].Meta().MinTime
			} else {
				startTime = time.Now().Unix() * 1000
			}
			// Add a safety margin as it may take a few minutes for everything to spin up.
			return startTime + s.startTimeMargin, nil
		default:
			return 0, errors.New("TSDB not implement")
		}
	}
	return math.MaxInt64, tsdb.ErrNotReady
}

func (s *PromStorage) Querier(mint, maxt int64) (storage.Querier, error) {
	if x := s.get(); x != nil {
		return x.Querier(mint, maxt)
	}
	return nil, tsdb.ErrNotReady
}

func (s *PromStorage) ChunkQuerier(mint, maxt int64) (storage.ChunkQuerier, error) {
	if x := s.get(); x != nil {
		return x.ChunkQuerier(mint, maxt)
	}
	return nil, tsdb.ErrNotReady
}

func (s *PromStorage) Appender(ctx context.Context) storage.Appender {
	if x := s.get(); x != nil {
		return x.Appender(ctx)
	}
	return notReadyAppend{}
}

func (s *PromStorage) ApplyConfig(conf *config.Config) error {
	db := s.get()
	if db, ok := db.(*tsdb.DB); ok {
		return db.ApplyConfig(conf)
	}
	return nil
}

func (s *PromStorage) ExemplarQuerier(ctx context.Context) (storage.ExemplarQuerier, error) {
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
	if x := s.get(); x != nil {
		return x.Close()
	}
	return nil
}

func (s *PromStorage) CleanTombstones() error {
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

func (s *PromStorage) Delete(ctx context.Context, mint, maxt int64, ms ...*labels.Matcher) error {
	if x := s.get(); x != nil {
		switch db := x.(type) {
		case *tsdb.DB:
			return db.Delete(ctx, mint, maxt, ms...)
		default:
			return errors.New("TSDB not implement")
		}
	}
	return tsdb.ErrNotReady
}

func (s *PromStorage) Snapshot(dir string, withHead bool) error {
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
	if x := s.getStats(); x != nil {
		return x.Head.WALReplayStatus.GetWALReplayStatus(), nil
	}
	return tsdb.WALReplayStatus{}, tsdb.ErrNotReady
}
