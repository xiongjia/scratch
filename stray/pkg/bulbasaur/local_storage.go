package bulbasaur

import (
	"context"
	"fmt"
	"log/slog"
	"stray/pkg/dugtrio"

	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/tsdb"
	"github.com/prometheus/prometheus/tsdb/chunkenc"
)

type LocalStorageEngine struct {
	logger *slog.Logger
	DB     storage.Storage
}

func (e *LocalStorageEngine) Test() {
	e.logger.Debug("test")

	ctx := context.Background()
	var maxt int64
	maxt = 0
	for i := 0; i < 121; i++ {
		app := e.DB.Appender(ctx)
		_, err := app.Append(0, labels.FromStrings("foo", "bar"), maxt, float64(i+3))
		if err != nil {
			continue
		}
		maxt++
		app.Commit()
	}

	querier, _ := e.DB.Querier(0, 200)

	sets := querier.Select(context.TODO(), false, nil, labels.MustNewMatcher(labels.MatchEqual, "foo", "bar"))
	fmt.Printf("sets: %v\n", sets)

	var series chunkenc.Iterator
	idx := 0
	for sets.Next() {
		series = sets.At().Iterator(series)
		for series.Next() == chunkenc.ValFloat {
			_, v := series.At()
			fmt.Printf("i = %d , v = %v\n", idx, v)
			idx += 1
		}
	}

}

type LocalStorageOptions struct {
	Logger *slog.Logger
	Dir    string
}

func MakeLocalStorageEngine(opts *LocalStorageOptions) (*LocalStorageEngine, error) {
	kitlogger := dugtrio.NewKitLoggerAdapterSlog(opts.Logger)
	db, err := tsdb.Open(opts.Dir, kitlogger, nil, tsdb.DefaultOptions(), nil)
	if err != nil {
		return nil, err
	}
	return &LocalStorageEngine{
		logger: opts.Logger,
		DB:     db,
	}, nil
}
