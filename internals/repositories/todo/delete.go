package todo

import (
	"context"

	ent "github.com/robowealth-mutual-fund/blueprint-service/internals/entities/todo"
	model "github.com/robowealth-mutual-fund/blueprint-service/internals/models/todo"
)

func (r *Repository) Delete(ctx context.Context, request *model.DeleteRequest) (*model.DeleteResponse, error) {
	err := r.oracle.Delete(ctx, map[string]interface{}{"ID": request.Id}, &ent.Todo{})
	if err != nil {
		return nil, err
	}

	return &model.DeleteResponse{
		Message: "delete successfully",
	}, nil
}
