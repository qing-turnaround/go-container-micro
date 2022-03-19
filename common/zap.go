package common

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	fileName  = "micro.log" // 日志文件名
	maxSize   = 512         // 日志的最大为200MB
	maxAge    = 30          // 日志保留最大天数
	maxBackup = 2           // 最大备份日志文件数
)

// ZapInit 初始化zap
func ZapInit() {
	writeSyncer := getLogWriter(fileName, maxSize, maxBackup, maxAge)
	encoder := getEncoder()
	level := zap.NewAtomicLevelAt(zap.DebugLevel)
	core := zapcore.NewCore(encoder, writeSyncer, level)
	log := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	zap.ReplaceGlobals(log) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}
