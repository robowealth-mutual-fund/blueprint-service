package todo

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	ent "github.com/robowealth-mutual-fund/blueprint-service/internals/entities/todo"
	model "github.com/robowealth-mutual-fund/blueprint-service/internals/models/todo"
)

func (r *Repository) Update(ctx context.Context, request *model.UpdateRequest) (*model.UpdateResponse, error) {
	filters := map[string]interface{}{"ID": request.Id}

	entity := ent.Todo{
		Id: request.Id,
		UpdatedAt: sql.NullInt64{
			Int64: time.Now().Unix(),
			Valid: true,
		},
	}

	if request.TaskName != "" {
		entity.TaskName = request.TaskName
	}

	if request.Status != "" {
		entity.Status = sql.NullString{
			String: request.Status,
			Valid:  true,
		}
	}

	todo, err := r.oracle.Count(ctx, filters, &ent.Todo{})
	if err != nil {
		return nil, err
	}

	if todo != 1 {
		return nil, fmt.Errorf("todo not found")
	}

	_, err = r.oracle.Update(ctx, filters, &entity)
	if err != nil {
		return nil, err
	}

	return &model.UpdateResponse{Id: request.Id}, nil
}
