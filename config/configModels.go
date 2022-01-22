package config

type CassandraServer struct {
	Port     string
	Password string
	Keyspace string
	Username string
	Servers  []string
}

type RedisServer struct {
	Port     string
	Password string
	Username string
	Address  string
	Database int
}

type PostgresSQL struct {
	Address  string
	Port     string
	Username string
	Password string
	Database string
	Schema   string
}

type RabbitMqt struct {
	Port     string
	Password string
	Username string
	Address  string
}

type RequestLimiterItemDetail struct {
	Key          string
	CountRequest int
	TimeType     string
	LimitTime    int
}

type StorageServer struct {
	AccessKey string
	SecretKey string
	Endpoint  string
	UseSSL    bool
}
