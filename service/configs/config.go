package configs

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Environment      string
	UserServiceHost  string
	UserServicePort  int
	CtxTimeout       int
	LogLevel         string
	HTTPHost         string
	HTTPPort         string
	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string
	RPCPort          string
	PostServiceHost  string
	PostServicePort  int
}

func Load() Config {
	c := Config{}
	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPHost = cast.ToString(getOrReturnDefault("HTTP_HOST", "localhost"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", "8080"))
	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "userss"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "1"))
	c.CtxTimeout = cast.ToInt("CTX_TIMEOUT")
	c.PostServiceHost= cast.ToString(getOrReturnDefault("POST_SERVICE_HOST","localhost"))
	c.PostServicePort=cast.ToInt(getOrReturnDefault("POST_SERVICE_PORT",8000))

	c.RPCPort = cast.ToString(getOrReturnDefault("RPC_PORT", ":9000"))
	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
