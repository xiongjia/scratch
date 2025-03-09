package testutils

import (
	"context"
	"io"
	"strconv"
	"time"

	"metric/pkg/cortex/chunk/cache"

	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/model/labels"

	"metric/pkg/cortex/chunk"
	promchunk "metric/pkg/cortex/chunk/encoding"
	"metric/pkg/cortex/util/flagext"
	"metric/pkg/cortex/util/validation"
)

const (
	userID = "userID"
)

const (
	offset64 = 14695981039346656037
	prime64  = 1099511628211
	offset32 = 2166136261
	prime32  = 16777619
)

// Fixture type for per-backend testing.
type Fixture interface {
	Name() string
	Clients() (chunk.IndexClient, chunk.Client, chunk.TableClient, chunk.SchemaConfig, io.Closer, error)
}

// CloserFunc is to io.Closer as http.HandlerFunc is to http.Handler.
type CloserFunc func() error

// Close implements io.Closer.
func (f CloserFunc) Close() error {
	return f()
}

// DefaultSchemaConfig returns default schema for use in test fixtures
func DefaultSchemaConfig(kind string) chunk.SchemaConfig {
	schemaConfig := chunk.DefaultSchemaConfig(kind, "v1", model.Now().Add(-time.Hour*2))
	return schemaConfig
}

// Setup a fixture with initial tables
func Setup(fixture Fixture, tableName string) (chunk.IndexClient, chunk.Client, io.Closer, error) {
	var tbmConfig chunk.TableManagerConfig
	flagext.DefaultValues(&tbmConfig)
	indexClient, chunkClient, tableClient, schemaConfig, closer, err := fixture.Clients()
	if err != nil {
		return nil, nil, nil, err
	}

	tableManager, err := chunk.NewTableManager(tbmConfig, schemaConfig, 12*time.Hour, tableClient, nil, nil, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	err = tableManager.SyncTables(context.Background())
	if err != nil {
		return nil, nil, nil, err
	}

	err = tableClient.CreateTable(context.Background(), chunk.TableDesc{
		Name: tableName,
	})

	return indexClient, chunkClient, closer, err
}

// CreateChunks creates some chunks for testing
func CreateChunks(startIndex, batchSize int, from model.Time, through model.Time) ([]string, []chunk.Chunk, error) {
	keys := []string{}
	chunks := []chunk.Chunk{}
	for j := 0; j < batchSize; j++ {
		chunk := dummyChunkFor(from, through, labels.Labels{
			{Name: model.MetricNameLabel, Value: "foo"},
			{Name: "index", Value: strconv.Itoa(startIndex*batchSize + j)},
		})
		chunks = append(chunks, chunk)
		keys = append(keys, chunk.ExternalKey())
	}
	return keys, chunks, nil
}

// hashNew initializes a new fnv64a hash value.
func hashNew() uint64 {
	return offset64
}

// hashAdd adds a string to a fnv64a hash value, returning the updated hash.
// Note this is the same algorithm as Go stdlib `sum64a.Write()`
func hashAdd(h uint64, s string) uint64 {
	for i := 0; i < len(s); i++ {
		h ^= uint64(s[i])
		h *= prime64
	}
	return h
}

// hashAdd adds a string to a fnv64a hash value, returning the updated hash.
func hashAddString(h uint64, s string) uint64 {
	for i := 0; i < len(s); i++ {
		h ^= uint64(s[i])
		h *= prime64
	}
	return h
}

// hashAddByte adds a byte to a fnv64a hash value, returning the updated hash.
func hashAddByte(h uint64, b byte) uint64 {
	h ^= uint64(b)
	h *= prime64
	return h
}

// Fingerprint runs the same algorithm as Prometheus labelSetToFingerprint()
func Fingerprint(labels labels.Labels) model.Fingerprint {
	sum := hashNew()
	for _, label := range labels {
		sum = hashAddString(sum, label.Name)
		sum = hashAddByte(sum, model.SeparatorByte)
		sum = hashAddString(sum, label.Value)
		sum = hashAddByte(sum, model.SeparatorByte)
	}
	return model.Fingerprint(sum)
}

func dummyChunkFor(from, through model.Time, metric labels.Labels) chunk.Chunk {
	cs := promchunk.New()

	for ts := from; ts <= through; ts = ts.Add(15 * time.Second) {
		_, err := cs.Add(model.SamplePair{Timestamp: ts, Value: 0})
		if err != nil {
			panic(err)
		}
	}

	chunk := chunk.NewChunk(
		userID,
		Fingerprint(metric),
		metric,
		cs,
		from,
		through,
	)
	// Force checksum calculation.
	err := chunk.Encode()
	if err != nil {
		panic(err)
	}
	return chunk
}

func SetupTestChunkStoreWithClients(indexClient chunk.IndexClient, chunksClient chunk.Client, tableClient chunk.TableClient) (chunk.Store, error) {
	var (
		tbmConfig chunk.TableManagerConfig
		schemaCfg = chunk.DefaultSchemaConfig("", "v10", 0)
	)
	flagext.DefaultValues(&tbmConfig)
	tableManager, err := chunk.NewTableManager(tbmConfig, schemaCfg, 12*time.Hour, tableClient, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	err = tableManager.SyncTables(context.Background())
	if err != nil {
		return nil, err
	}

	var limits validation.Limits
	flagext.DefaultValues(&limits)
	limits.MaxQueryLength = model.Duration(30 * 24 * time.Hour)
	overrides, err := validation.NewOverrides(limits, nil)
	if err != nil {
		return nil, err
	}

	var storeCfg chunk.StoreConfig
	flagext.DefaultValues(&storeCfg)

	store := chunk.NewCompositeStore(nil)
	err = store.AddPeriod(storeCfg, schemaCfg.Configs[0], indexClient, chunksClient, overrides, cache.NewNoopCache(), cache.NewNoopCache())
	if err != nil {
		return nil, err
	}

	return store, nil
}

func SetupTestChunkStore() (chunk.Store, error) {
	storage := chunk.NewMockStorage()
	return SetupTestChunkStoreWithClients(storage, storage, storage)
}

func SetupTestObjectStore() (chunk.ObjectClient, error) {
	return chunk.NewMockStorage(), nil
}
