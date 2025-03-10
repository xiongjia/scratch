package prom

import (
	"time"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

type (
	CassandraClientOptions struct {
		Hosts []string
		Port  int

		Keyspace string

		Username string
		Password string
	}
)

func MakeCqlSession(opts CassandraClientOptions) (gocqlx.Session, error) {
	clusterConfig := gocql.NewCluster(opts.Hosts...)
	clusterConfig.Port = opts.Port
	clusterConfig.Keyspace = opts.Keyspace
	clusterConfig.Consistency = gocql.Quorum
	clusterConfig.Timeout = 30 * time.Second
	clusterConfig.RetryPolicy = &gocql.SimpleRetryPolicy{NumRetries: 3}
	clusterConfig.MaxWaitSchemaAgreement = 2 * time.Minute
	clusterConfig.Authenticator = gocql.PasswordAuthenticator{
		Username: opts.Username,
		Password: opts.Password,
	}
	return gocqlx.WrapSession(clusterConfig.CreateSession())
}

func CqlMultiExec(session *gocqlx.Session, stmt []string) error {
	for _, s := range stmt {
		err := session.ExecStmt(s)
		if err != nil {
			return err
		}
	}
	return nil
}
