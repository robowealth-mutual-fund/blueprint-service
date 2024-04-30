package config

type Server struct {
	GRPCPort                 string `env:"GRPC_SERVER_PORT" envDefault:"3001"`
	HTTPPort                 string `env:"HTTP_SERVER_PORT" envDefault:"3002"`
	HTTPServerTimeout        int64  `env:"HTTP_SERVER_TIMEOUT" envDefault:"0"`
	GRPCMaxConcurrentStreams uint32 `env:"GRPC_MAX_CONCURRENT_STREAMS" envDefault:"500"`
	GRPCWriteBufferSize      int    `env:"GRPC_WRITE_BUFFER_SIZE" envDefault:"100"`
	GRPCReadBufferSize       int    `env:"GRPC_READ_BUFFER_SIZE" envDefault:"100"`
}
