package todo

import (
	"context"

	model "github.com/robowealth-mutual-fund/blueprint-service/internals/models/todo"
)

func (s *Service) Get(ctx context.Context, request *model.GetRequest) (*model.GetResponse, error) {
	return s.repository.Get(ctx, request)
}
