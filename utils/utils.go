package utils

import (
	"os"

	"github.com/gocql/gocql"
)

func GetCassandraSession() (*gocql.Session, error) {
	cluster := gocql.NewCluster(os.Getenv("CASS_HOST"))
	return cluster.CreateSession()

}
