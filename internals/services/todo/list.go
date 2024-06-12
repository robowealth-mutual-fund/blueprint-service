package todo

import (
	"context"

	model "github.com/robowealth-mutual-fund/blueprint-service/internals/models/todo"
)

func (s *Service) List(ctx context.Context) (*[]model.ListResponse, error) {
	return s.repository.List(ctx)
}
