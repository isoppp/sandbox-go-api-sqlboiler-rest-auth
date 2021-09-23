package logger

import (
	"sandbox-go-api-sqlboiler-rest-auth/internal/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(cfg *config.Config) (*zap.Logger, error) {
	var logger *zap.Logger

	if cfg.IsDev {
		zapConfig := zap.NewDevelopmentConfig()
		encConfig := zap.NewDevelopmentEncoderConfig()
		encConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		zapConfig.EncoderConfig = encConfig
		logger, err := zapConfig.Build()
		if err != nil {
			return logger, err
		}
		return logger, nil
	}

	logger, err := zap.NewProduction()
	if err != nil {
		return logger, err
	}
	return logger, nil
}
