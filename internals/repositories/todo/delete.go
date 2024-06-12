package todo

import (
	"context"

	ent "github.com/robowealth-mutual-fund/blueprint-service/internals/entities/todo"
	model "github.com/robowealth-mutual-fund/blueprint-service/internals/models/todo"
)

func (r *Repository) Delete(ctx context.Context, request *model.DeleteRequest) (*model.DeleteResponse, error) {
	filter := map[string]interface{}{"ID": request.Id}
	entity := ent.Todo{}

	err := r.oracle.Delete(ctx, filter, &entity)
	if err != nil {
		return nil, err
	}

	return &model.DeleteResponse{Message: "delete successfully"}, nil
}
