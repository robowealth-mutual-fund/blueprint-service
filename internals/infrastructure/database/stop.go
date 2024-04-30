package database

import (
	log "github.com/robowealth-mutual-fund/stdlog"
)

func (db *DB) Stop() {
	if err := db.Sql.Close(); err != nil {
		log.Error("Error: Closing DB Connection", err)
	} else {
		log.Info("DB Connection Closed")
	}
}
