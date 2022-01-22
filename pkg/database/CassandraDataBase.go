package database

import (
	"facegram_file_server/config"
	"github.com/fatih/color"
	"github.com/gocql/gocql"
	"time"
)

var cassandraConnection *gocql.Session

func GetCassandraConnection() (*gocql.Session, error) {

	if cassandraConnection != nil {
		return cassandraConnection, nil
	}

	cfg := config.GetCassandraConfig()
	cluster := gocql.NewCluster(cfg.Servers...)
	cluster.Timeout = time.Second * 10
	cluster.Keyspace = cfg.Keyspace
	session, err := cluster.CreateSession()

	if err != nil {
		return nil, err
	}

	cassandraConnection = session
	return cassandraConnection, nil
}

func init() {
	_, err := GetCassandraConnection()
	if err == nil {
		c := color.New(color.FgGreen).PrintfFunc()
		c("Successfully connected to cassandra\n")
	} else {
		c := color.New(color.FgRed).PrintfFunc()
		c("Failed Initializing cassandra Connection!\n")
	}
}
