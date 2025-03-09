package chunk

import (
	prom_chunk "metric/pkg/cortex/chunk/encoding"

	"github.com/golang/snappy"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/model/labels"
)

// Chunk contains encoded timeseries data
type Chunk struct {
	// These two fields will be missing from older chunks (as will the hash).
	// On fetch we will initialise these fields from the DynamoDB key.
	Fingerprint model.Fingerprint `json:"fingerprint"`
	UserID      string            `json:"userID"`

	// These fields will be in all chunks, including old ones.
	From    model.Time    `json:"from"`
	Through model.Time    `json:"through"`
	Metric  labels.Labels `json:"metric"`

	// The hash is not written to the external storage either.  We use
	// crc32, Castagnoli table.  See http://www.evanjones.ca/crc32c.html.
	// For old chunks, ChecksumSet will be false.
	ChecksumSet bool   `json:"-"`
	Checksum    uint32 `json:"-"`

	// We never use Delta encoding (the zero value), so if this entry is
	// missing, we default to DoubleDelta.
	Encoding prom_chunk.Encoding `json:"encoding"`
	Data     prom_chunk.Chunk    `json:"-"`

	// The encoded version of the chunk, held so we don't need to re-encode it
	encoded []byte
}

// DecodeContext holds data that can be re-used between decodes of different chunks
type DecodeContext struct {
	reader *snappy.Reader
}

// NewDecodeContext creates a new, blank, DecodeContext
func NewDecodeContext() *DecodeContext {
	return &DecodeContext{
		reader: snappy.NewReader(nil),
	}
}
