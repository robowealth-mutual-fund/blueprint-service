package oracle

import "github.com/robowealth-mutual-fund/blueprint-service/internals/infrastructure/database"

type Repository struct {
	db *database.DB
}

func New(db *database.DB) Interface {
	return &Repository{
		db: db,
	}
}
