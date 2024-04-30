package config

type Connection struct {
	DebugMode          bool `env:"RESTY_DEBUG_MODE" envDefault:"false"`
	InsecureSkipVerify bool `env:"INSECURE_SKIP_VERIFY" envDefault:"true"`

	//custom http client
	HttpTimeout             int `env:"HTTP_TIMEOUT" envDefault:"20"`
	HttpMaxIdleConns        int `env:"HTTP_MAX_IDLE_CONNS" envDefault:"100"`
	HttpMaxConnsPerHost     int `env:"HTTP_MAX_CONNS_PER_HOST" envDefault:"100"`
	HttpMaxIdleConnsPerHost int `env:"HTTP_MAX_IDLE_CONNS_PER_HOST" envDefault:"100"`
}
