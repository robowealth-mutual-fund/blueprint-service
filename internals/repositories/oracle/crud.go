package oracle

import (
	"context"

	"github.com/kisielk/sqlstruct"
	"github.com/robowealth-mutual-fund/blueprint-service/internals/utils"
)

func (r *Repository) Count(ctx context.Context, filters any, entity any) (int64, error) {
	var count int64

	query, args, err := GetCountBuilder(filters, entity)
	if err != nil {
		return count, err
	}

	err = r.db.Sql.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return count, err
	}

	return count, nil
}

func (r *Repository) Create(ctx context.Context, entity any) (any, error) {
	query, args, err := GetInsertBuilder(entity)
	if err != nil {
		return nil, err
	}

	stmt, err := r.db.Sql.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = stmt.Close()
		if err != nil {
			return
		}
	}()

	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return nil, err
	}

	return result, err
}

//func (r *Repository) CreateInBatch(ctx context.Context, value any, batchSize int) error {
//
//}

func (r *Repository) Delete(ctx context.Context, filters, entity any) error {
	query, args, err := GetDeleteBuilder(filters, entity)
	if err != nil {
		return err
	}

	_, err = r.db.Sql.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) First(ctx context.Context, orderBy string, selects []string, filters map[string]any, entity any) error {
	query, args, err := GetFirstBuilder(selects, orderBy, filters, entity)
	if err != nil {
		return err
	}

	rows, err := r.db.Sql.QueryContext(ctx, query, args...)
	if err != nil {
		return err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			return
		}
	}()

	for rows.Next() {
		err = sqlstruct.Scan(entity, rows)

		if err != nil {
			return err
		}
	}

	return err
}

func (r *Repository) Last(ctx context.Context, orderBy string, selects []string, filters map[string]any, entity any) error {
	query, args, err := GetLastBuilder(selects, orderBy, filters, entity)
	if err != nil {
		return err
	}

	rows, err := r.db.Sql.QueryContext(ctx, query, args...)
	if err != nil {
		return err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			return
		}
	}()

	for rows.Next() {
		err = sqlstruct.Scan(entity, rows)

		if err != nil {
			return err
		}
	}

	return err
}

func (r *Repository) List(ctx context.Context, offset, limit int64, orderBy string, selects []string, filters, entity any) (*utils.Pagination, error) {
	var total int64

	query, args, err := GetCountBuilder(filters, entity)
	if err != nil {
		return nil, err
	}

	err = r.db.Sql.QueryRowContext(ctx, query, args...).Scan(&total)
	if err != nil {
		return nil, err
	}

	query, args, err = GetListBuilder(offset, limit, orderBy, selects, filters, entity)
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Sql.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			return
		}
	}()

	items := make([]any, 0)

	for rows.Next() {
		item, err := Clone(entity)
		if err != nil {
			return nil, err
		}

		err = sqlstruct.Scan(item, rows)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return utils.FormatPagination(items, limit, total), nil
}

func (r *Repository) Raw(ctx context.Context, entity any, query string, args ...any) ([]any, error) {
	items := make([]any, 0)

	rows, err := r.db.Sql.QueryContext(ctx, query, args...)
	if err != nil {
		return items, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			return
		}
	}()

	for rows.Next() {
		item, err := Clone(entity)
		if err != nil {
			return items, err
		}

		err = sqlstruct.Scan(item, rows)
		if err != nil {
			return items, err
		}

		items = append(items, item)
	}

	return items, nil
}

func (r *Repository) Find(ctx context.Context, orderBy string, selects []string, filters, entity any) ([]any, error) {
	items := make([]any, 0)

	query, args, err := GetFindBuilder(orderBy, selects, filters, entity)
	if err != nil {
		return items, err
	}

	rows, err := r.db.Sql.QueryContext(ctx, query, args...)
	if err != nil {
		return items, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			return
		}
	}()

	for rows.Next() {
		item, err := Clone(entity)
		if err != nil {
			return items, err
		}

		err = sqlstruct.Scan(item, rows)
		if err != nil {
			return items, err
		}

		items = append(items, item)
	}

	return items, nil
}

func (r *Repository) Update(ctx context.Context, filters, entity any) (any, error) {
	query, args, err := GetUpdateBuilder(filters, entity)
	if err != nil {
		return nil, err
	}

	stmt, err := r.db.Sql.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = stmt.Close()
		if err != nil {
			return
		}
	}()

	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return nil, err
	}

	return result, err
}

//func (r *Repository) IsErrorRecordNotFound(err error) bool {
//	return r.db.IsErrorRecordNotFound(err)
//}
