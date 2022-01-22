package config

import (
	"errors"
	"facegram_file_server/pkg/utility"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func GetRequestLimiterConfigBuilder(key string) (*RequestLimiterItemDetail, error) {
	v := *utility.GetEnv(key, "")
	if v == "" {
		return nil, errors.New("error")
	}
	s := strings.Split(v, ",")
	if len(s) == 4 {
		cr, err := strconv.Atoi(s[1])
		if err != nil {
			return nil, err
		}

		lt, err := strconv.Atoi(s[3])
		if err != nil {
			return nil, err
		}

		return &RequestLimiterItemDetail{
			Key:          s[0],
			CountRequest: cr,
			TimeType:     s[2],
			LimitTime:    lt,
		}, nil
	}
	return nil, errors.New("error")
}

func GetCassandraConfig() *CassandraServer {
	servers := strings.Split(*utility.GetEnv("CASSANDRA_SERVERS", "localhost:9042"), ",")
	return &CassandraServer{
		Port:     *utility.GetEnv("CASSANDRA_PORT", ":7000"),
		Password: *utility.GetEnv("CASSANDRA_PASSWORD", ""),
		Keyspace: *utility.GetEnv("CASSANDRA_APP_KEYSPACE", "fileserver"),
		Username: *utility.GetEnv("CASSANDRA_USER_NAME", ""),
		Servers:  servers,
	}
}

func GetPostgresSqlConfig() *PostgresSQL {
	return &PostgresSQL{
		Address:  *utility.GetEnv("POSTGRESQL_ADDRESS", ""),
		Port:     *utility.GetEnv("POSTGRESQL_PORT", ""),
		Username: *utility.GetEnv("POSTGRESQL_USER_NAME", ""),
		Password: *utility.GetEnv("POSTGRESQL_PASSWORD", ""),
		Database: *utility.GetEnv("POSTGRESQL_DATABASE", ""),
		Schema:   *utility.GetEnv("POSTGRESQL_SCHEMA", ""),
	}
}

func GetRedisConfig() *RedisServer {
	db, err := strconv.Atoi(*utility.GetEnv("REDIS_DATABASE", "1"))
	if err != nil {
		db = 1
	}
	return &RedisServer{
		Port:     *utility.GetEnv("REDIS_PORT", ""),
		Password: *utility.GetEnv("REDIS_PASSWORD", ""),
		Username: *utility.GetEnv("REDIS_USER_NAME", ""),
		Address:  *utility.GetEnv("REDIS_ADDRESS", ""),
		Database: db,
	}
}

func GetRabbitMqtConfig() *RabbitMqt {
	return &RabbitMqt{
		Port:     *utility.GetEnv("RABBIT_MQT_PORT", ""),
		Password: *utility.GetEnv("RABBIT_MQT_PASSWORD", ""),
		Username: *utility.GetEnv("RABBIT_MQT_USER_NAME", ""),
		Address:  *utility.GetEnv("RABBIT_MQT_ADDRESS", ""),
	}
}

func GetStorageServerConfig() *StorageServer {
	ssl, err := strconv.ParseBool(*utility.GetEnv("STORAGE_USE_SSL", "true"))
	if err != nil {
		ssl = true
	}
	return &StorageServer{
		AccessKey: *utility.GetEnv("STORAGE_ACCESS_KEY", ""),
		SecretKey: *utility.GetEnv("STORAGE_SECRET_KEY", ""),
		Endpoint:  *utility.GetEnv("STORAGE_ENDPOINT", ""),
		UseSSL:    ssl,
	}
}

func GetInternalKeyConfig() *string {
	return utility.GetEnv("INTERNAL_KEY", "")
}

func GetExternalKeyConfig() *string {
	return utility.GetEnv("EXTERNAL_KEY", "")
}

func GetRpcKeyConfig() *string {
	return utility.GetEnv("RPC_KEY", "")
}
