package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error

	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""
	config.EncoderConfig = encoderConfig

	// config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)

	// log, err = zap.NewProduction()
	// defer log.Sync()

	log, err = config.Build(zap.AddCallerSkip(1))
	// log.Error("zap Error test")

	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) {
	log.With(
		zap.Namespace("data"),
	).Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}
func Fatal(message string, fields ...zap.Field) {
	log.Fatal(message, fields...)
}
