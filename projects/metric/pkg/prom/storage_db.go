package prom

import (
	"context"
	"database/sql"
	"sync"

	_ "github.com/mattn/go-sqlite3"
	"github.com/prometheus/prometheus/config"
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

		db *sql.DB

		rwMtx sync.RWMutex
		labs  map[uint64]rcdLabs
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
	dbQuerier struct {
		log     kitlog.Logger
		ctx     context.Context
		storage *storageDb

		mint int64
		maxt int64
	}
)

func makeStorageDb(dbPath string, log kitlog.Logger) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		_ = level.Error(log).Log("msg", "open db error", "err", err)
		return nil, err
	}
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS lab (id integer not null primary key, labs text, metadata text);

        CREATE TABLE IF NOT EXISTS rcd (id integer not null primary key, lid integer not null, ts integer not null, val REAL );
	`)
	if err != nil {
		db.Close()
		_ = level.Error(log).Log("msg", "init db schema", "err", err)
		return nil, err
	}
	return db, nil
}

func loadLabs(db *sql.DB, log kitlog.Logger) (map[uint64]rcdLabs, error) {
	sql := `SELECT id, ifnull(labs, "") as labs, ifnull(metadata, "") as metadata FROM lab`
	rows, err := db.Query(sql)
	if err != nil {
		_ = level.Error(log).Log("msg", "load cache", "err", err)
		return nil, err
	}
	defer rows.Close()

	labs := make(map[uint64]rcdLabs)
	for rows.Next() {
		var id uint64
		var content string
		var metadata string
		err = rows.Scan(&id, &content, &metadata)
		if err != nil {
			continue
		}

		item, err := labsFromJson(content)
		if err != nil {
			continue
		}
		itemHash := item.Hash()
		rcd := rcdLab{Id: id, Lab: item, Metadata: labMetadataFromJsonDefault(metadata, nil)}
		buckets, ok := labs[itemHash]
		if ok {
			labs[itemHash] = append(buckets, rcd)
		} else {
			labs[itemHash] = append(make(rcdLabs, 0), rcd)
		}
	}
	return labs, nil
}

func createDbStorage(dbPath string, log kitlog.Logger) (storage.Storage, error) {
	db, err := makeStorageDb(dbPath, log)
	if err != nil {
		_ = level.Error(log).Log("msg", "open db", "err", err)
		return nil, err
	}

	// load cache to memory
	labs, err := loadLabs(db, log)
	if err != nil {
		db.Close()
		_ = level.Error(log).Log("msg", "load db", "err", err)
		return nil, err
	}
	return &storageDb{log: log, db: db, labs: labs}, nil
}

func (s *storageDb) dbAddRcd(labsId uint64, t int64, v float64) error {
	sql := "INSERT INTO rcd (lid, ts, val) values (?, ?, ?)"
	statement, err := s.db.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = statement.Exec(labsId, t, v)
	if err != nil {
		return err
	}
	return nil
}

func (s *storageDb) dbAddRcdLabMetadata(labId uint64, m *metadata.Metadata) error {
	content, err := labMetadataToJson(m)
	if err != nil {
		return err
	}
	sql := "UPDATE lab set metadata = ? WHERE id = ?"
	statement, err := s.db.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = statement.Exec(content, labId)
	if err != nil {
		return err
	}
	return nil
}

func (s *storageDb) dbAddRcdLab(l labels.Labels) (uint64, error) {
	content, err := labsToJson(l)
	if err != nil {
		return 0, nil
	}

	sql := "INSERT INTO lab (labs) values (?)"
	statement, err := s.db.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(content)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return (uint64)(id), nil
}

func (s *storageDb) rcdFindLabs(labHash uint64, l labels.Labels) uint64 {
	s.rwMtx.RLock()
	defer s.rwMtx.RUnlock()

	buckets, ok := s.labs[labHash]
	if !ok {
		return 0
	}
	for _, bucket := range buckets {
		if labels.Equal(l, bucket.Lab) {
			return bucket.Id
		}
	}
	return 0
}

func (s *storageDb) rcdPutLabs(labHash uint64, l labels.Labels) (uint64, error) {
	labId := s.rcdFindLabs(labHash, l)
	if labId != 0 {
		return labId, nil
	}

	s.rwMtx.Lock()
	defer s.rwMtx.Unlock()

	labId, err := s.dbAddRcdLab(l)
	if err != nil {
		return 0, err
	}

	// Insert it to DB
	rcd := rcdLab{Id: labId, Lab: l.Copy()}
	buckets, ok := s.labs[labHash]
	if ok {
		buckets = append(buckets, rcd)
		s.labs[labHash] = buckets
	} else {
		labs := make(rcdLabs, 0)
		labs = append(labs, rcd)
		s.labs[labHash] = labs
	}
	return rcd.Id, nil
}

func (s *storageDb) rcdPut(labsId uint64, t int64, v float64) error {
	err := s.dbAddRcd(labsId, t, v)
	if err != nil {
		_ = level.Error(s.log).Log("msg", "rcd add", "err", err)
	}
	return err
}

func (s *storageDb) updateBucket(bucket *rcdLab, m *metadata.Metadata) {
	if bucket == nil || m == nil {
		return
	}

	if labMetadataEqual(m, bucket.Metadata) {
		return
	}

	bucket.Metadata = labMetadataCopy(m)
	err := s.dbAddRcdLabMetadata(bucket.Id, m)
	if err != nil {
		_ = level.Error(s.log).Log("msg", "lab metadata", "err", err)
	}
}

func (s *storageDb) rcdLabsUpdateMetadata(ref storage.SeriesRef, l labels.Labels, m *metadata.Metadata) (storage.SeriesRef, error) {
	if m == nil {
		return ref, nil
	}

	labHash := l.Hash()
	s.rwMtx.Lock()
	defer s.rwMtx.Unlock()
	buckets, ok := s.labs[labHash]
	if !ok {
		return ref, nil
	}
	for _, bucket := range buckets {
		if labels.Equal(l, bucket.Lab) {
			s.updateBucket(&bucket, m)
			return storage.SeriesRef(bucket.Id), nil
		}
	}
	return ref, nil
}

func (s *storageDb) rcdAppend(ref storage.SeriesRef, l labels.Labels, t int64, v float64) (storage.SeriesRef, error) {
	if v == 0.0 {
		_ = level.Debug(s.log).Log("msg", "0.0", "l", l)
	}

	if ref != 0 {
		return ref, s.rcdPut(uint64(ref), t, v)
	}
	srcLabs := labsSort(l)
	srcLabsHash := srcLabs.Hash()
	labsId, err := s.rcdPutLabs(srcLabsHash, srcLabs)
	if err != nil {
		return 0, err
	}
	_ = level.Debug(s.log).Log("msg", "rcd append", "labsId", labsId, "labsHash", srcLabsHash)
	return storage.SeriesRef(labsId), s.rcdPut(labsId, t, v)
}

func (s *storageDb) Querier(ctx context.Context, mint, maxt int64) (storage.Querier, error) {
	_ = level.Debug(s.log).Log("msg", "Querier", "mint", mint, "maxt", maxt)
	return &dbQuerier{log: s.log, ctx: ctx, storage: s, mint: mint, maxt: maxt}, nil
}

func (s *storageDb) ChunkQuerier(ctx context.Context, mint, maxt int64) (storage.ChunkQuerier, error) {
	_ = level.Debug(s.log).Log("msg", "ChunkQuerier", "mint", mint, "maxt", maxt)
	return nil, tsdb.ErrNotReady
}

func (s *storageDb) ExemplarQuerier(ctx context.Context) (storage.ExemplarQuerier, error) {
	_ = level.Debug(s.log).Log("msg", "ExemplarQuerier")
	return nil, tsdb.ErrNotReady
}

func (s *storageDb) Appender(ctx context.Context) storage.Appender {
	_ = level.Debug(s.log).Log("msg", "Appender")
	return &dbAppender{log: s.log, storage: s, ctx: ctx}
}

func (s *storageDb) ApplyConfig(conf *config.Config) error {
	_ = level.Debug(s.log).Log("msg", "ApplyConfig", "conf", conf)
	return nil
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

func (q *dbQuerier) Select(sortSeries bool, hints *storage.SelectHints, matchers ...*labels.Matcher) storage.SeriesSet {

	return nil
}

func (q *dbQuerier) LabelValues(name string, matchers ...*labels.Matcher) ([]string, storage.Warnings, error) {
	return nil, nil, nil
}

func (q *dbQuerier) LabelNames(matchers ...*labels.Matcher) ([]string, storage.Warnings, error) {
	return nil, nil, nil
}

func (q *dbQuerier) Close() error {
	return nil
}
