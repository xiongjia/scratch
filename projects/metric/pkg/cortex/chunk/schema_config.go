package chunk

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/prometheus/common/model"
)

var (
	errInvalidSchemaVersion     = errors.New("invalid schema version")
	errInvalidTablePeriod       = errors.New("the table period must be a multiple of 24h (1h for schema v1)")
	errConfigFileNotSet         = errors.New("schema config file needs to be set")
	errConfigChunkPrefixNotSet  = errors.New("schema config for chunks is missing the 'prefix' setting")
	errSchemaIncreasingFromTime = errors.New("from time in schemas must be distinct and in increasing order")
)

// Bucket describes a range of time with a tableName and hashKey
type Bucket struct {
	from       uint32
	through    uint32
	tableName  string
	hashKey    string
	bucketSize uint32 // helps with deletion of series ids in series store. Size in milliseconds.
}

// PeriodConfig defines the schema and tables to use for a period of time
type PeriodConfig struct {
	From        DayTime             `yaml:"from"`         // used when working with config
	IndexType   string              `yaml:"store"`        // type of index client to use.
	ObjectType  string              `yaml:"object_store"` // type of object client to use; if omitted, defaults to store.
	Schema      string              `yaml:"schema"`
	IndexTables PeriodicTableConfig `yaml:"index"`
	ChunkTables PeriodicTableConfig `yaml:"chunks"`
	RowShards   uint32              `yaml:"row_shards"`
}

// DayTime is a model.Time what holds day-aligned values, and marshals to/from
// YAML in YYYY-MM-DD format.
type DayTime struct {
	model.Time
}

// SchemaConfig contains the config for our chunk index schemas
type SchemaConfig struct {
	Configs []PeriodConfig `yaml:"configs"`

	fileName string
}

func defaultRowShards(schema string) uint32 {
	switch schema {
	case "v1", "v2", "v3", "v4", "v5", "v6", "v9":
		return 0
	default:
		return 16
	}
}

func validateChunks(cfg PeriodConfig) error {
	objectStore := cfg.IndexType
	if cfg.ObjectType != "" {
		objectStore = cfg.ObjectType
	}
	switch objectStore {
	case "cassandra", "aws-dynamo", "bigtable-hashed", "gcp", "gcp-columnkey", "bigtable", "grpc-store":
		if cfg.ChunkTables.Prefix == "" {
			return errConfigChunkPrefixNotSet
		}
		return nil
	default:
		return nil
	}
}

// PeriodicTableConfig is configuration for a set of time-sharded tables.
type PeriodicTableConfig struct {
	Prefix string
	Period time.Duration
	Tags   Tags
}

func (cfg *PeriodConfig) applyDefaults() {
	if cfg.RowShards == 0 {
		cfg.RowShards = defaultRowShards(cfg.Schema)
	}
}

// Validate the period config.
func (cfg PeriodConfig) validate() error {
	validateError := validateChunks(cfg)
	if validateError != nil {
		return validateError
	}

	_, err := cfg.CreateSchema()
	return err
}

// CreateSchema returns the schema defined by the PeriodConfig
func (cfg PeriodConfig) CreateSchema() (BaseSchema, error) {
	buckets, bucketsPeriod := cfg.createBucketsFunc()

	// Ensure the tables period is a multiple of the bucket period
	if cfg.IndexTables.Period > 0 && cfg.IndexTables.Period%bucketsPeriod != 0 {
		return nil, errInvalidTablePeriod
	}

	if cfg.ChunkTables.Period > 0 && cfg.ChunkTables.Period%bucketsPeriod != 0 {
		return nil, errInvalidTablePeriod
	}

	switch cfg.Schema {
	case "v1":
		return newStoreSchema(buckets, originalEntries{}), nil
	case "v2":
		return newStoreSchema(buckets, originalEntries{}), nil
	case "v3":
		return newStoreSchema(buckets, base64Entries{originalEntries{}}), nil
	case "v4":
		return newStoreSchema(buckets, labelNameInHashKeyEntries{}), nil
	case "v5":
		return newStoreSchema(buckets, v5Entries{}), nil
	case "v6":
		return newStoreSchema(buckets, v6Entries{}), nil
	case "v9":
		return newSeriesStoreSchema(buckets, v9Entries{}), nil
	case "v10", "v11":
		if cfg.RowShards == 0 {
			return nil, fmt.Errorf("Must have row_shards > 0 (current: %d) for schema (%s)", cfg.RowShards, cfg.Schema)
		}

		v10 := v10Entries{rowShards: cfg.RowShards}
		if cfg.Schema == "v10" {
			return newSeriesStoreSchema(buckets, v10), nil
		}

		return newSeriesStoreSchema(buckets, v11Entries{v10}), nil
	default:
		return nil, errInvalidSchemaVersion
	}
}

// Validate the schema config and returns an error if the validation
// doesn't pass
func (cfg *SchemaConfig) Validate() error {
	for i := range cfg.Configs {
		periodCfg := &cfg.Configs[i]
		periodCfg.applyDefaults()
		if err := periodCfg.validate(); err != nil {
			return err
		}

		if i+1 < len(cfg.Configs) {
			if cfg.Configs[i].From.Time.Unix() >= cfg.Configs[i+1].From.Time.Unix() {
				return errSchemaIncreasingFromTime
			}
		}
	}
	return nil
}
