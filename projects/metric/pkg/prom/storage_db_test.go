package prom

import (
	"testing"
	"time"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

func TestDB(t *testing.T) {
	cluster := &gocql.ClusterConfig{
		Hosts:   []string{"192.168.71.237"},
		Timeout: time.Duration(30 * time.Second),
		Port:    9042,
	}
	t.Logf("cluster %s", cluster)

	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		return
	}
	t.Logf("session %s", session)
}
