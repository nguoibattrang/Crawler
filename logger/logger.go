package logger

import (
	"go.uber.org/zap"
)

var Log *zap.Logger

// InitLogger initializes the zap logger.
func InitLogger() {
	var err error
	Log, err = zap.NewProduction() // Use zap.NewDevelopment() for development mode
	if err != nil {
		panic("failed to initialize logger: " + err.Error())
	}
	defer Log.Sync() // Flushes buffer, if any
}
