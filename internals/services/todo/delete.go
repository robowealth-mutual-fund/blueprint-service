package todo

import (
	"context"

	model "github.com/robowealth-mutual-fund/blueprint-service/internals/models/todo"
)

func (s *Service) Delete(ctx context.Context, request *model.DeleteRequest) (*model.DeleteResponse, error) {
	return s.repository.Delete(ctx, request)
}
