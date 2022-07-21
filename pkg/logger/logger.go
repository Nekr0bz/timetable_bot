// Package logger from https://github.com/bigwhite/experiments/blob/master/uber-zap-advanced-usage/
package logger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const timeLayout = "2006-01-02 15:04:05.000 MST"

var l *zap.Logger

// NewAppLogger create a new logger (not support logger rotating).
func NewAppLogger() *zap.Logger {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(timeLayout))
	}

	logger, _ := cfg.Build()
	return logger
}

func InitAppLogger() {
	l = NewAppLogger()
}

func GetLogger() *zap.Logger {
	return l
}
