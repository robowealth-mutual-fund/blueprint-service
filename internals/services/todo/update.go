package todo

import (
	"context"

	model "github.com/robowealth-mutual-fund/blueprint-service/internals/models/todo"
)

func (s *Service) Update(ctx context.Context, request *model.UpdateRequest) (*model.UpdateResponse, error) {
	return s.repository.Update(ctx, request)
}
