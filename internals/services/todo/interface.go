package todo

import (
	"context"

	model "github.com/robowealth-mutual-fund/blueprint-service/internals/models/todo"
)

//go:generate mockery --name=Interface
type Interface interface {
	Create(ctx context.Context, request *model.CreateRequest) (*model.CreateResponse, error)
	Delete(ctx context.Context, request *model.DeleteRequest) (*model.DeleteResponse, error)
	Update(ctx context.Context, request *model.UpdateRequest) (*model.UpdateResponse, error)
	List(ctx context.Context) (*[]model.ListResponse, error)
}
