package chunk

import (
	"time"

	"github.com/prometheus/common/model"
)

// DefaultSchemaConfig creates a simple schema config for testing
func DefaultSchemaConfig(store, schema string, from model.Time) SchemaConfig {
	s := SchemaConfig{
		Configs: []PeriodConfig{{
			IndexType: store,
			Schema:    schema,
			From:      DayTime{from},
			ChunkTables: PeriodicTableConfig{
				Prefix: "cortex",
				Period: 7 * 24 * time.Hour,
			},
			IndexTables: PeriodicTableConfig{
				Prefix: "cortex_chunks",
				Period: 7 * 24 * time.Hour,
			},
		}},
	}
	if err := s.Validate(); err != nil {
		panic(err)
	}
	return s
}
