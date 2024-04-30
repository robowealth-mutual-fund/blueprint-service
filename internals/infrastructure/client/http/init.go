package http

import (
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/robowealth-mutual-fund/blueprint-service/internals/config"
)

type HttpClient struct {
	Config      config.Config
	RestyClient *resty.Client
}

func NewHttpClient(config config.Config) *HttpClient {

	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = config.Connection.HttpMaxIdleConns
	t.MaxConnsPerHost = config.Connection.HttpMaxConnsPerHost
	t.MaxIdleConnsPerHost = config.Connection.HttpMaxIdleConnsPerHost

	httpClient := &http.Client{
		Timeout:   time.Duration(config.Connection.HttpTimeout) * time.Second,
		Transport: t,
	}

	restyClient := resty.NewWithClient(httpClient)

	//otelresty.TraceClient(restyClient, []otelresty.Option{
	//	otelresty.WithTracerName("resty-tracer"),
	//	otelresty.WithIsTraceRequestEnabled(true),
	//	otelresty.WithIsTraceResponseEnabled(true),
	//	otelresty.WithIsLogRequestEnabled(true),
	//	otelresty.WithLogRequestName("curl request"),
	//}...)

	return &HttpClient{
		Config:      config,
		RestyClient: restyClient,
	}
}
