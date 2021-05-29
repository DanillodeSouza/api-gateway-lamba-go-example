package apigatewaylambdagoexample

import (
	"context"
	"runtime"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger returns an usable zap logger.
func NewLogger(logLevel LogLevelConfig) (*zap.Logger, error) {
	logger, err := configLog(zap.NewAtomicLevelAt(logLevel.Value)).Build()
	if err != nil {
		return nil, err
	}
	return logger, nil
}

// LogError logs a error to stderr including some internal information like
// transaction ID, internal ID, error code and message.
func LogError(ctx context.Context, logger *zap.Logger, msg, route string) {
	_, filename, _, _ := runtime.Caller(1)
	logger.Error("failed to process request",
		zap.String("error-message", msg),
		zap.String("route", route),
		zap.String("filename", filename),
	)
}

func configLog(level zap.AtomicLevel) zap.Config {
	return zap.Config{
		Level:         level,
		Development:   false,
		DisableCaller: true,
		Sampling:      nil,
		Encoding:      "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: millisDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

func millisDurationEncoder(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendInt(int(float64(d) / float64(time.Millisecond)))
}
