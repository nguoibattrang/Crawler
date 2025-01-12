package logger

import (
	"errors"
	"go.uber.org/zap"
)

// InitLogger initializes the zap logger.
func InitLogger(mode string) (*zap.Logger, error) {
	switch mode {
	case "development":
		return zap.NewDevelopment()
	case "production":
		return zap.NewProduction()
	default:
		return nil, errors.New("invalid log mode")
	}
}
