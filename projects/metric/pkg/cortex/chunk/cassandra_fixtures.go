package chunk

import (
	"context"
	"io"
	"os"

	"metric/pkg/cortex/util/flagext"
)

// GOCQL doesn't provide nice mocks, so we use a real Cassandra instance.
// To enable these tests:
// $ docker run -d --name cassandra --rm -p 9042:9042 cassandra:3.11
// $ CASSANDRA_TEST_ADDRESSES=localhost:9042 go test ./pkg/chunk/storage

type fixture struct {
	name      string
	addresses string
}

func (f *fixture) Name() string {
	return f.name
}

func (f *fixture) Clients() (IndexClient, Client, TableClient, SchemaConfig, io.Closer, error) {
	var cfg CassandraConfig
	flagext.DefaultValues(&cfg)
	cfg.Addresses = f.addresses
	cfg.Keyspace = "test"
	cfg.Consistency = "QUORUM"
	cfg.ReplicationFactor = 1

	// Get a SchemaConfig with the defaults.
	schemaConfig := TestUtilDefaultSchemaConfig("cassandra")

	storageClient, err := NewStorageClient(cfg, schemaConfig, nil)
	if err != nil {
		return nil, nil, nil, schemaConfig, nil, err
	}

	objectClient, err := NewObjectClient(cfg, schemaConfig, nil)
	if err != nil {
		return nil, nil, nil, schemaConfig, nil, err
	}

	tableClient, err := NewTableClient(context.Background(), cfg, nil)
	if err != nil {
		return nil, nil, nil, schemaConfig, nil, err
	}

	closer := CloserFunc(func() error {
		storageClient.Stop()
		objectClient.Stop()
		tableClient.Stop()
		return nil
	})

	return storageClient, objectClient, tableClient, schemaConfig, closer, nil
}

// Fixtures for unit testing Cassandra integration.
func Fixtures() []Fixture {
	addresses := os.Getenv("CASSANDRA_TEST_ADDRESSES")
	if addresses == "" {
		return nil
	}

	return []Fixture{
		&fixture{
			name:      "Cassandra",
			addresses: addresses,
		},
	}
}
