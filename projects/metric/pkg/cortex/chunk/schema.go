package chunk

import (
	"metric/pkg/cortex/querier/astmapper"

	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/model/labels"
)

// IndexQuery describes a query for entries
type IndexQuery struct {
	TableName string
	HashValue string

	// One of RangeValuePrefix or RangeValueStart might be set:
	// - If RangeValuePrefix is not nil, must read all keys with that prefix.
	// - If RangeValueStart is not nil, must read all keys from there onwards.
	// - If neither is set, must read all keys for that row.
	RangeValuePrefix []byte
	RangeValueStart  []byte

	// Filters for querying
	ValueEqual []byte

	// If the result of this lookup is immutable or not (for caching).
	Immutable bool
}

type baseEntries interface {
	GetReadMetricQueries(bucket Bucket, metricName string) ([]IndexQuery, error)
	GetReadMetricLabelQueries(bucket Bucket, metricName string, labelName string) ([]IndexQuery, error)
	GetReadMetricLabelValueQueries(bucket Bucket, metricName string, labelName string, labelValue string) ([]IndexQuery, error)
	FilterReadQueries(queries []IndexQuery, shard *astmapper.ShardAnnotation) []IndexQuery
}

type schemaBucketsFunc func(from, through model.Time, userID string) []Bucket

// baseSchema implements BaseSchema given a bucketing function and and set of range key callbacks
type baseSchema struct {
	buckets schemaBucketsFunc
	entries baseEntries
}

// IndexEntry describes an entry in the chunk index
type IndexEntry struct {
	TableName string
	HashValue string

	// For writes, RangeValue will always be set.
	RangeValue []byte

	// New for v6 schema, label value is not written as part of the range key.
	Value []byte
}

// used by storeSchema
type storeEntries interface {
	baseEntries

	GetWriteEntries(bucket Bucket, metricName string, labels labels.Labels, chunkID string) ([]IndexEntry, error)
}

// storeSchema implements StoreSchema given a bucketing function and and set of range key callbacks
type storeSchema struct {
	baseSchema
	entries storeEntries
}

func newStoreSchema(buckets schemaBucketsFunc, entries storeEntries) storeSchema {
	return storeSchema{
		baseSchema: baseSchema{buckets: buckets, entries: entries},
		entries:    entries,
	}
}
