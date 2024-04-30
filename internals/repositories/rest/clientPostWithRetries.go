package rest

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	restyV2 "github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/robowealth-mutual-fund/blueprint-service/internals/constants"
	log "github.com/sirupsen/logrus"
)

func (r *Client) ClientPostWithRetries(ctx context.Context, host, path string, headers map[string]string, request, entity, errorEntity interface{}) error {
	restyClient := r.HttpClient.RestyClient
	rReq := restyClient.
		SetTLSClientConfig(&tls.Config{
			InsecureSkipVerify: r.Config.Connection.InsecureSkipVerify,
		}).
		SetRetryCount(3).
		SetRetryWaitTime(5 * time.Second).
		SetRetryMaxWaitTime(20 * time.Second).
		AddRetryCondition(
			func(r *restyV2.Response, err error) bool {
				return err != nil || r.StatusCode() == http.StatusInternalServerError || r.StatusCode() == http.StatusGatewayTimeout
			},
		).
		R().
		SetContext(ctx).
		ForceContentType("application/json").
		SetBody(request).
		SetResult(&entity).
		SetError(&errorEntity)

	if len(headers) > 0 {
		rReq.SetHeaders(headers)
	}

	response, err := rReq.Post(fmt.Sprintf("%s/%s", host, path))
	if err != nil {
		return errors.Wrap(err, constants.ErrPostHttpRequest)
	}

	statusCode := response.StatusCode()
	if statusCode != http.StatusOK {
		errorMessage := fmt.Sprintf("%s %s", response.Status(), string(response.Body()))
		return errors.Wrap(errors.New(errorMessage), constants.ErrIsNotOK)
	}

	if err != nil {
		return errors.Wrap(err, "Error Unmarshal Response Body")
	}

	log.Info("response with retry: " + fmt.Sprintf("%v", response))

	return err
}
