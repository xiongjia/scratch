package chunk

import (
	"context"
	"io"
	"time"

	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/model/labels"

	"metric/pkg/cortex/util/flagext"
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
	Clients() (IndexClient, Client, TableClient, SchemaConfig, io.Closer, error)
}

// CloserFunc is to io.Closer as http.HandlerFunc is to http.Handler.
type CloserFunc func() error

// Close implements io.Closer.
func (f CloserFunc) Close() error {
	return f()
}

// DefaultSchemaConfig returns default schema for use in test fixtures
func TestUtilDefaultSchemaConfig(kind string) SchemaConfig {
	schemaConfig := DefaultSchemaConfig(kind, "v1", model.Now().Add(-time.Hour*2))
	return schemaConfig
}

// Setup a fixture with initial tables
func Setup(fixture Fixture, tableName string) (IndexClient, Client, io.Closer, error) {
	var tbmConfig TableManagerConfig
	flagext.DefaultValues(&tbmConfig)
	indexClient, chunkClient, tableClient, schemaConfig, closer, err := fixture.Clients()
	if err != nil {
		return nil, nil, nil, err
	}

	tableManager, err := NewTableManager(tbmConfig, schemaConfig, 12*time.Hour, tableClient, nil, nil, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	err = tableManager.SyncTables(context.Background())
	if err != nil {
		return nil, nil, nil, err
	}

	err = tableClient.CreateTable(context.Background(), TableDesc{
		Name: tableName,
	})

	return indexClient, chunkClient, closer, err
}

// CreateChunks creates some chunks for testing
func CreateChunks(startIndex, batchSize int, from model.Time, through model.Time) ([]string, []Chunk, error) {
	keys := []string{}
	chunks := []Chunk{}
	for j := 0; j < batchSize; j++ {
		// chunk := dummyChunkFor(from, through, labels.Labels{
		// 	{Name: model.MetricNameLabel, Value: "foo"},
		// 	{Name: "index", Value: strconv.Itoa(startIndex*batchSize + j)},
		// })
		// chunks = append(chunks, chunk)
		// keys = append(keys, chunk.ExternalKey())
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

// func dummyChunkFor(from, through model.Time, metric labels.Labels) Chunk {
// 	cs := promchunk.New()

// 	for ts := from; ts <= through; ts = ts.Add(15 * time.Second) {
// 		_, err := cs.Add(model.SamplePair{Timestamp: ts, Value: 0})
// 		if err != nil {
// 			panic(err)
// 		}
// 	}

// 	chunk := NewChunk(
// 		userID,
// 		Fingerprint(metric),
// 		metric,
// 		cs,
// 		from,
// 		through,
// 	)
// 	// Force checksum calculation.
// 	err := chunk.Encode()
// 	if err != nil {
// 		panic(err)
// 	}
// 	return chunk
// }

// // NewChunk creates a new chunk
// func NewChunk(userID string, fp model.Fingerprint, metric labels.Labels, c prom_chunk.Chunk, from, through model.Time) Chunk {
// 	return Chunk{
// 		Fingerprint: fp,
// 		UserID:      userID,
// 		From:        from,
// 		Through:     through,
// 		Metric:      metric,
// 		Encoding:    c.Encoding(),
// 		Data:        c,
// 	}
// }
