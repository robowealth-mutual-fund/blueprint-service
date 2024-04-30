package rest

import (
	"github.com/robowealth-mutual-fund/blueprint-service/internals/config"
	"github.com/robowealth-mutual-fund/blueprint-service/internals/infrastructure/client/http"
)

type Client struct {
	HttpClient *http.HttpClient
	Config     config.Config
}

func New(httpClient *http.HttpClient, config config.Config) Interface {
	return &Client{
		HttpClient: httpClient,
		Config:     config,
	}
}
