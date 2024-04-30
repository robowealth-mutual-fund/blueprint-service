package config

type Redis struct {
	SentinelHost    string `env:"SENTINEL_HOST" envDefault:":6379"`
	RedisPassword   string `env:"REDIS_PASSWORD" envDefault:"password"`
	RedisExpiration string `env:"REDIS_EXPIRATION" envDefault:"5m"`
	MasterName      string `env:"MASTER_NAME" envDefault:"mymaster"`
}
