package utils

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zapLog *zap.Logger
var zapLogger *zap.SugaredLogger

func InitLogger(env string) {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)

	logFile, _ := os.OpenFile("logs/log.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	writer := zapcore.AddSync(logFile)
	defaultLogLevel := zapcore.DebugLevel
	core := zapcore.NewTee(zapcore.NewCore(fileEncoder, writer, defaultLogLevel))
	if env == "d" {
		consoleEncoder := zapcore.NewConsoleEncoder(config)
		core = zapcore.NewTee(
			zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
			zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
		)
	}
	zapLog = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	zapLogger = zapLog.Sugar()
}

func Logger() *zap.SugaredLogger {
	return zapLogger
}

func FastLogger() *zap.Logger {
	return zapLog
}
