package http

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
	model "github.com/robowealth-mutual-fund/blueprint-service/internals/models/resty"
	log "github.com/robowealth-mutual-fund/stdlog"
)

func (c *HttpClient) RequestWithJSON(ctx context.Context, request *model.Request) (*resty.Response, error) {
	var (
		response *resty.Response
		err      error
	)

	client := c.RestyClient

	restyRequest := client.
		SetBaseURL(request.Host).
		SetLogger(&CustomLogger{Debug: c.Config.Connection.DebugMode}).
		SetDebug(c.Config.Connection.DebugMode).
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: c.Config.Connection.InsecureSkipVerify}).
		R().
		SetContext(ctx).
		ForceContentType("application/json").
		SetHeaders(request.Header).
		SetBody(request.Body)

	if request.PathParams != nil {
		restyRequest.SetPathParams(request.PathParams)
	}

	if request.QueryParams != nil {
		restyRequest.SetQueryParams(request.QueryParams)
	}

	switch request.Method {
	case resty.MethodGet:
		response, err = restyRequest.Get(request.Path)
	case resty.MethodPost:
		response, err = restyRequest.Post(request.Path)
	case resty.MethodPut:
		response, err = restyRequest.Put(request.Path)
	case resty.MethodPatch:
		response, err = restyRequest.Patch(request.Path)
	case resty.MethodDelete:
		response, err = restyRequest.Delete(request.Path)
	default:
		return nil, errors.New("method not match")
	}

	if err != nil {
		err.Error()
		return response, err
	}

	log.Info(fmt.Sprintf("url : %s/%s", request.Host, request.Path))
	log.Info(fmt.Sprintf("response body : %s", string(response.Body())))
	log.Info(fmt.Sprintf("statusCode : %d", response.StatusCode()))

	return response, nil
}
