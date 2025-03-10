package prom

import (
	"testing"

	"fmt"
)

const (
	metric_keyspace      = "metric"
	metric_role          = "metric"
	metric_role_password = "metric"
)

func TestDb(t *testing.T) {
	cqlSession, err := MakeCqlSession(CassandraClientOptions{
		Hosts:    []string{"192.168.71.237"},
		Port:     9042,
		Keyspace: "system",
		Username: "cassandra",
		Password: "cassandra",
	})
	if err != nil {
		return
	}
	defer cqlSession.Close()

	err = CqlMultiExec(&cqlSession, []string{
		fmt.Sprintf("CREATE ROLE IF NOT EXISTS %s WITH LOGIN = true AND PASSWORD = '%s'", metric_role, metric_role_password),
		fmt.Sprintf("CREATE KEYSPACE IF NOT EXISTS %s WITH replication = {'class' : 'SimpleStrategy', 'replication_factor' : %d}", metric_keyspace, 1),
		fmt.Sprintf("GRANT all PERMISSIONS ON KEYSPACE %s TO %s", metric_keyspace, metric_role),
	})
	if err != nil {
		t.Logf("err: %v", err)
		return
	}

	// batch := session.NewBatch(gocql.UnloggedBatch)
	// batch.Query(fmt.Sprintf("GRANT all PERMISSIONS ON KEYSPACE %s TO %s", metric_keyspace, metric_role))
	// err = session.ExecuteBatch(batch)
	// if err != nil {
	// 	t.Logf("err: %v", err)
	// 	return
	// }
}
