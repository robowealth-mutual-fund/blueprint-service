package todo

import (
	"context"

	model "github.com/robowealth-mutual-fund/blueprint-service/internals/models/todo"
)

func (s *Service) Create(ctx context.Context, request *model.CreateRequest) (*model.CreateResponse, error) {
	return s.repository.Create(ctx, request)
}
