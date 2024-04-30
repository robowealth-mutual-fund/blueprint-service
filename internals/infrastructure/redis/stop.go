package redis

import log "github.com/robowealth-mutual-fund/stdlog"

func (cache *Redis) Stop() {
	if err := cache.RedisClient.Close(); err != nil {
		log.Error("Error: Closing DB Connection", err)
	} else {
		log.Info("DB Connection Closed")
	}
}
