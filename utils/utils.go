package utils

import (
	"log"
	"os"

	"github.com/gocql/gocql"
)

func GetCassandraSession() (*gocql.Session, error) {
	log.Printf("Connecting to %s", os.Getenv("CASS_HOST"))
	cluster := gocql.NewCluster(os.Getenv("CASS_HOST"))
	return cluster.CreateSession()

}
