package config

type Database struct {
	Driver            string   `env:"DB_DRIVER" envDefault:"oracle"`
	Host              []string `env:"DB_HOST" envDefault:"localhost"  envSeparator:","`
	Port              string   `env:"DB_PORT" envDefault:"1521"`
	DatabaseName      string   `env:"DB_NAME" envDefault:"FREEPDB1"`
	User              string   `env:"DB_USER" envDefault:"TESTUSR"`
	Password          string   `env:"DB_PASSWORD" envDefault:"PASSWORD"`
	ConnMaxLifetime   int      `env:"DB_CONN_MAX_LIFETIME" envDefault:"300"`
	MaxConnection     int      `env:"DB_MAX_CONN" envDefault:"7"`
	MaxIdleConnection int      `env:"DB_MAX_IDLE_CONN" envDefault:"5"`
	HostCluster       string   `env:"DB_HOST_CLUSTER" envDefault:""`
}
