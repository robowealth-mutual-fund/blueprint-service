package todo

import (
	"context"
	"fmt"

	ent "github.com/robowealth-mutual-fund/blueprint-service/internals/entities/todo"
	model "github.com/robowealth-mutual-fund/blueprint-service/internals/models/todo"
)

func (r *Repository) Get(ctx context.Context, request *model.GetRequest) (*model.GetResponse, error) {
	selects := []string{"ID", "TASK_NAME", "STATUS", "CREATED_AT", "UPDATED_AT"}
	filters := map[string]interface{}{"ID": request.Id}
	entity := &ent.Todo{}

	err := r.oracle.First(ctx, "ID", selects, filters, entity)
	if err != nil {
		return nil, err
	}

	if *entity == (ent.Todo{}) {
		return nil, fmt.Errorf("todo not found")
	}

	return &model.GetResponse{
		Id:        entity.Id,
		TaskName:  entity.TaskName,
		Status:    entity.Status.String,
		CreatedAt: entity.CreatedAt.Int64,
		UpdatedAt: entity.UpdatedAt.Int64,
	}, nil
}
