package initialize

import (
	"gin_skeleton/g"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() {

	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	g.Logger = logger
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {

	// 日志写入文件
	writer := GetLogWriter("server.serverLog")

	return zapcore.AddSync(writer)
}
