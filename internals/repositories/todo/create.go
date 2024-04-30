package todo

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	ent "github.com/robowealth-mutual-fund/blueprint-service/internals/entities/todo"
	model "github.com/robowealth-mutual-fund/blueprint-service/internals/models/todo"
)

func (r *Repository) Create(ctx context.Context, request *model.CreateRequest) (*model.CreateResponse, error) {
	entity := ent.Todo{
		Id:       uuid.NewString(),
		TaskName: request.TaskName,
		Status: sql.NullString{
			String: request.Status,
			Valid:  true,
		},
		CreatedAt: sql.NullInt64{
			Int64: time.Now().Unix(),
			Valid: true,
		},
		UpdatedAt: sql.NullInt64{
			Int64: time.Now().Unix(),
			Valid: true,
		},
	}

	_, err := r.oracle.Create(ctx, &entity)
	if err != nil {
		return nil, err
	}

	return &model.CreateResponse{
		TaskName:  entity.TaskName,
		Status:    entity.Status.String,
		CreatedAt: entity.CreatedAt.Int64,
		UpdatedAt: entity.UpdatedAt.Int64,
	}, nil
}
