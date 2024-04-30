package rest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/robowealth-mutual-fund/blueprint-service/internals/constants"
	"github.com/robowealth-mutual-fund/blueprint-service/internals/models/resty"
)

func (r *Client) RestyClientPost(ctx context.Context, host, path string, request interface{}) ([]byte, error) {

	response, err := r.HttpClient.RequestWithJSON(
		ctx,
		&resty.Request{
			Host:   host,
			Method: resty.MethodPost,
			Path:   path,
			Body:   request,
			Header: map[string]string{
				"Content-Type": "application/json",
			},
		},
	)

	if err != nil {
		return response.Body(), errors.Wrap(err, constants.ErrGetData)
	}

	body := response.Body()
	statusCode := response.StatusCode()

	if (statusCode != http.StatusOK) && (statusCode != http.StatusCreated) {
		errorMessage := fmt.Sprintf("%s %s", response.Status(), string(response.Body()))
		return body, errors.Wrap(errors.New(errorMessage), constants.ErrIsNotOK)
	}
	return body, err
}
