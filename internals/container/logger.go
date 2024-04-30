package container

import (
	"github.com/robowealth-mutual-fund/blueprint-service/internals/config"
	log "github.com/robowealth-mutual-fund/stdlog"
)

func (c *Container) NewLogger(appConfig config.Config) {
	level := log.DEBUG_LEVEL

	if appConfig.Environment == "PRODUCTION" {
		level = log.INFO_LEVEL
	}

	log.SetGlobalLogLevel(level)
	log.SetGlobalPlatformName(appConfig.PlatformName)
}
