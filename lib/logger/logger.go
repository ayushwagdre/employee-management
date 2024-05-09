package logger

import (
	"context"
	"fmt"
	"os"
	"practice/lib/environment"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger
var jsonEncoder zapcore.Encoder

type LoggingFunc func(message string, fields ...zapcore.Field)

const (
	INFO  = 1
	ERROR = 2
)

func Init(mode int, env environment.Environment) {
	var logLevel zapcore.Level
	switch mode {
	case INFO:
		logLevel = zapcore.InfoLevel
	case ERROR:
		logLevel = zapcore.ErrorLevel
	}

	cfg := zap.Config{
		Encoding: "json",
		Level:    zap.NewAtomicLevelAt(logLevel),
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,
		},
	}

	logger, _ = cfg.Build()
	jsonEncoder = zapcore.NewJSONEncoder(cfg.EncoderConfig)

	if env == environment.DevEnv {
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		logger = logger.WithOptions(
			zap.WrapCore(
				func(zapcore.Core) zapcore.Core {
					return zapcore.NewCore(zapcore.NewConsoleEncoder(cfg.EncoderConfig), zapcore.AddSync(os.Stderr), zapcore.DebugLevel)
				}))
	} else {
		logger = logger.WithOptions(
			zap.WrapCore(
				func(zapcore.Core) zapcore.Core {
					return zapcore.NewCore(jsonEncoder, zapcore.AddSync(os.Stderr), logLevel)
				}))
	}

}

func Get() *zap.Logger {
	return logger
}

func Field(key string, value interface{}) zapcore.Field {
	return zap.Any(key, value)
}

func I(ctx context.Context, message string, fields ...zapcore.Field) {
	logger.Info(message, fields...)
}

func E(ctx context.Context, err error, message string, fields ...zapcore.Field) {
	fields = append(fields, Field("error", err))
}

func Sync() {
	logger.Info("SYNCING LOGGER....")
	err := logger.Sync()
	if err != nil {
		fmt.Println("FAILED TO SYNC LOGGER...")
	}
}
