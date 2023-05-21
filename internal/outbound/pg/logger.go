package pg

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/tracelog"
)

func NewTraceLog() *tracelog.TraceLog {
	return &tracelog.TraceLog{
		LogLevel: tracelog.LogLevelTrace,
		Logger:   &logger{},
	}
}

type logger struct {
}

func (lg *logger) Log(_ context.Context, level tracelog.LogLevel, msg string, data map[string]any) {
	dataString := ""

	for key, value := range data {
		dataString += fmt.Sprintf(" %s=%v", key, value)
	}

	log.Printf("[%s] %s%s", level.String(), msg, dataString)
}
