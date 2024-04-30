package database

import (
	"errors"

	"gorm.io/gorm"
)

func (db *DB) IsErrorRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
