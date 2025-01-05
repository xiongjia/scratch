package main

import (
	"context"
	"fmt"
	"log/slog"
	"math"
	"sync"
	"time"

	"stray/internal/server"
	"stray/pkg/metric"
	metricUtil "stray/pkg/metric/util"
	"stray/pkg/util"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/prometheus/config"
	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/scrape"
	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/tsdb"
	"github.com/prometheus/prometheus/util/logging"
)

type readyStorage struct {
	mtx sync.RWMutex

	stats           *tsdb.DBStats
	startTimeMargin int64
	db              storage.Storage
}

func (s *readyStorage) ApplyConfig(conf *config.Config) error {
	db := s.get()
	if db, ok := db.(*tsdb.DB); ok {
		return db.ApplyConfig(conf)
	}
	return nil
}

func (s *readyStorage) getStats() *tsdb.DBStats {
	s.mtx.RLock()
	x := s.stats
	s.mtx.RUnlock()
	return x
}

func (s *readyStorage) get() storage.Storage {
	s.mtx.RLock()
	x := s.db
	s.mtx.RUnlock()
	return x
}

func (s *readyStorage) Set(db storage.Storage, startTimeMargin int64) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.db = db
	s.startTimeMargin = startTimeMargin
}

func (s *readyStorage) StartTime() (int64, error) {
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
			panic(fmt.Sprintf("unknown storage type %T", db))
		}
	}
	return math.MaxInt64, tsdb.ErrNotReady
}

// Querier implements the Storage interface.
func (s *readyStorage) Querier(mint, maxt int64) (storage.Querier, error) {
	if x := s.get(); x != nil {
		return x.Querier(mint, maxt)
	}
	return nil, tsdb.ErrNotReady
}

// ChunkQuerier implements the Storage interface.
func (s *readyStorage) ChunkQuerier(mint, maxt int64) (storage.ChunkQuerier, error) {
	if x := s.get(); x != nil {
		return x.ChunkQuerier(mint, maxt)
	}
	return nil, tsdb.ErrNotReady
}

func (s *readyStorage) Appender(ctx context.Context) storage.Appender {
	if x := s.get(); x != nil {
		return x.Appender(ctx)
	}
	return metricUtil.NewNotReadyAppender()
}

func (s *readyStorage) ExemplarQuerier(ctx context.Context) (storage.ExemplarQuerier, error) {
	if x := s.get(); x != nil {
		switch db := x.(type) {
		case *tsdb.DB:
			return db.ExemplarQuerier(ctx)
		default:
			panic(fmt.Sprintf("unknown storage type %T", db))
		}
	}
	return nil, tsdb.ErrNotReady
}

func (s *readyStorage) Close() error {
	if x := s.get(); x != nil {
		return x.Close()
	}
	return nil
}

func (s *readyStorage) CleanTombstones() error {
	if x := s.get(); x != nil {
		switch db := x.(type) {
		case *tsdb.DB:
			return db.CleanTombstones()
		default:
			panic(fmt.Sprintf("unknown storage type %T", db))
		}
	}
	return tsdb.ErrNotReady
}

func (s *readyStorage) Delete(ctx context.Context, mint, maxt int64, ms ...*labels.Matcher) error {
	if x := s.get(); x != nil {
		switch db := x.(type) {
		case *tsdb.DB:
			return db.Delete(ctx, mint, maxt, ms...)
		default:
			panic(fmt.Sprintf("unknown storage type %T", db))
		}
	}
	return tsdb.ErrNotReady
}

func (s *readyStorage) Snapshot(dir string, withHead bool) error {
	if x := s.get(); x != nil {
		switch db := x.(type) {
		case *tsdb.DB:
			return db.Snapshot(dir, withHead)
		default:
			panic(fmt.Sprintf("unknown storage type %T", db))
		}
	}
	return tsdb.ErrNotReady
}

func (s *readyStorage) Stats(statsByLabelName string, limit int) (*tsdb.Stats, error) {
	if x := s.get(); x != nil {
		switch db := x.(type) {
		case *tsdb.DB:
			return db.Head().Stats(statsByLabelName, limit), nil
		default:
			panic(fmt.Sprintf("unknown storage type %T", db))
		}
	}
	return nil, tsdb.ErrNotReady
}

func (s *readyStorage) WALReplayStatus() (tsdb.WALReplayStatus, error) {
	if x := s.getStats(); x != nil {
		return x.Head.WALReplayStatus.GetWALReplayStatus(), nil
	}
	return tsdb.WALReplayStatus{}, tsdb.ErrNotReady
}

func procInit() {
	util.InitDefaultLog(&util.LogOption{
		Level:     slog.LevelDebug,
		AddSource: false,
	})
	util.DebugDumpDeps()

	log := slog.Default()
	// logLvl := &promslog.AllowedLevel{}
	// logLvl.Set("debug")
	// promLog := promslog.New(&promslog.Config{Level: logLvl})
	// slog.SetDefault(promLog)
	// promLog.Debug("test")

	promLog := log.With("component", "scrape manager")

	localStorage := &readyStorage{stats: tsdb.NewDBStats()}

	// remoteStorage = remote.NewStorage(logger.With("component", "remote"), prometheus.DefaultRegisterer, localStorage.StartTime, localStoragePath, time.Duration(cfg.RemoteFlushDeadline), scraper, cfg.scrape.AppendMetadata)

	fanout := storage.NewFanout(promLog, localStorage)

	scrape.NewManager(
		&scrape.Options{},
		promLog,
		logging.NewJSONFileLogger,
		fanout,
		prometheus.DefaultRegisterer)

}

func main() {
	procInit()

	serv, err := server.NewServer(server.ServerConfig{
		EnableApiDoc: true,
	})
	if err != nil {
		slog.Error("Server init error", slog.Any("err", err))
		return
	}

	eng, err := metric.NewEngine()
	if err != nil {
		slog.Error("Engine init error", slog.Any("err", err))
		return
	}

	wg := util.NewWaitGroup()
	wg.Go(func() {
		slog.Info("API Server Started (0.0.0.0:8897)")
		err = server.StartServer("0.0.0.0", 8897, serv)
		if err != nil {
			slog.Error("Server error", slog.Any("err", err))
		}
	})

	wg.Go(func() {
		eng.Run()
	})

	wg.Wait()
}
