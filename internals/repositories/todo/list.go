package todo

import (
	"context"
	"fmt"

	ent "github.com/robowealth-mutual-fund/blueprint-service/internals/entities/todo"
	model "github.com/robowealth-mutual-fund/blueprint-service/internals/models/todo"
)

func (r *Repository) List(ctx context.Context) (*[]model.ListResponse, error) {
	filters := map[string]interface{}{}

	res, err := r.oracle.Find(ctx, "ID", []string{}, filters, &ent.Todo{})
	if err != nil {
		return nil, err
	}

	var response []model.ListResponse
	for _, d := range res {
		todo, ok := d.(*ent.Todo)
		if !ok {
			return nil, fmt.Errorf("unexpected element type: %T", d)
		}
		setDt := model.ListResponse{
			Id:        todo.Id,
			TaskName:  todo.TaskName,
			Status:    map[bool]string{true: todo.Status.String, false: ""}[todo.Status.Valid],
			CreatedAt: map[bool]int64{true: todo.CreatedAt.Int64, false: 0}[todo.CreatedAt.Valid],
			UpdatedAt: map[bool]int64{true: todo.UpdatedAt.Int64, false: 0}[todo.UpdatedAt.Valid],
		}

		response = append(response, setDt)
	}

	return &response, nil
}
