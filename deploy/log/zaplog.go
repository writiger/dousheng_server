package zaplog

import (
	hertzzap "github.com/hertz-contrib/logger/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var ZapLogger *hertzzap.Logger

func InitLogger() {
	ZapLogger = hertzzap.NewLogger(
		hertzzap.WithCores([]hertzzap.CoreConfig{
			{
				Enc: zapcore.NewJSONEncoder(simpleEncoderConfig()),
				Ws:  getWriteSyncer("./deploy/log/info/log.log"),
				Lvl: zap.NewAtomicLevelAt(zapcore.LevelOf(
					zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
						return lev == zap.InfoLevel
					}))),
			},
			{
				Enc: zapcore.NewJSONEncoder(simpleEncoderConfig()),
				Ws:  getWriteSyncer("./deploy/log/warn/log.log"),
				Lvl: zap.NewAtomicLevelAt(zapcore.LevelOf(
					zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
						return lev == zap.WarnLevel
					}))),
			},
			{
				Enc: zapcore.NewJSONEncoder(simpleEncoderConfig()),
				Ws:  getWriteSyncer("./deploy/log/error/log.log"),
				Lvl: zap.NewAtomicLevelAt(zapcore.LevelOf(
					zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
						return lev == zap.ErrorLevel
					}))),
			},
		}...))
	defer ZapLogger.Sync()
}

func simpleEncoderConfig() zapcore.EncoderConfig {
	cfg := basicEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncodeLevel = zapcore.CapitalLevelEncoder
	cfg.EncodeDuration = zapcore.StringDurationEncoder
	return cfg
}

func getWriteSyncer(file string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file,
		MaxSize:    10,
		MaxBackups: 50000,
		MaxAge:     1000,
		Compress:   true,
		LocalTime:  true,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func basicEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "name",
		TimeKey:        "ts",
		CallerKey:      "caller",
		FunctionKey:    "func",
		StacktraceKey:  "stacktrace",
		LineEnding:     "\n",
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}
