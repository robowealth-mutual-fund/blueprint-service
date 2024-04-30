package http

import (
	"encoding/json"
	"fmt"

	log "github.com/robowealth-mutual-fund/stdlog"
)

type CustomLogger struct {
	Debug bool
}

type LogEntry struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}

func (l *CustomLogger) Debugf(format string, args ...interface{}) {
	if l.Debug {
		l.logEntry("debug", format, args...)
	}
}

func (l *CustomLogger) Errorf(format string, args ...interface{}) {
	l.logEntry("error", format, args...)
}

func (l *CustomLogger) Warnf(format string, args ...interface{}) {
	l.logEntry("warn", format, args...)
}

func (l *CustomLogger) logEntry(level string, format string, args ...interface{}) {
	entry := LogEntry{
		Level:   level,
		Message: fmt.Sprintf(format, args...),
	}

	logJSON, err := json.Marshal(entry)
	if err != nil {
		log.Error("Error encoding log entry:", err)
		return
	}

	log.Info(string(logJSON))
}
