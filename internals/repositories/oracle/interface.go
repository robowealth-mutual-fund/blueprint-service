package oracle

import (
	"context"

	"github.com/robowealth-mutual-fund/blueprint-service/internals/utils"
)

//go:generate mockery --dir=./ --name=Interface --filename=crud.go --output=mocks --outpkg=mocks
type Interface interface {
	Count(ctx context.Context, filters any, entity any) (int64, error)
	Create(ctx context.Context, entity interface{}) (interface{}, error)
	//CreateInBatch(ctx context.Context, value any, batchSize int) error
	Delete(ctx context.Context, filters, entity interface{}) error
	Find(ctx context.Context, orderBy string, selects []string, filters, entity interface{}) ([]any, error)
	First(ctx context.Context, orderBy string, selects []string, filters map[string]any, entity any) error
	//IsErrorRecordNotFound(err error) bool
	Last(ctx context.Context, orderBy string, selects []string, filters map[string]any, entity any) error
	List(ctx context.Context, offset, limit int64, orderBy string, selects []string, filters, entity interface{}) (*utils.Pagination, error)
	Raw(ctx context.Context, entity any, sql string, value ...any) ([]any, error)
	Update(ctx context.Context, filters, entity interface{}) (interface{}, error)
}
