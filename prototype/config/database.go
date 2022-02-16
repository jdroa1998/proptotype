package config

type PostgresConnection struct {
	DatabaseHost     string
	DatabasePort     string
	DatabaseName     string
	DatabaseUser     string
	DatabasePassword string
}

type PostgresConfig struct {
	PostgresConnection PostgresConnection
}

var DB_HOST = "localhost"
var DB_PORT = "5432"
var DB_NAME = "test"
var DB_USER = "postgres"
var DB_PASSWORD = "12345"

func GetPostgresConnection() PostgresConfig {
	return PostgresConfig{
		PostgresConnection: PostgresConnection{
			DatabaseHost:     DB_HOST,
			DatabasePort:     DB_PORT,
			DatabaseName:     DB_NAME,
			DatabaseUser:     DB_USER,
			DatabasePassword: DB_PASSWORD,
		},
	}
}
