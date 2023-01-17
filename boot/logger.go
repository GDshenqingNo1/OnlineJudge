package boot

import (
	g "OnlineJudge/app/global"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

func LoggerSetup() {
	dynamicLevel := zap.NewAtomicLevel()
	switch g.Config.Logger.LogLevel {
	case "debug":
		dynamicLevel.SetLevel(zap.DebugLevel)
	case "info":
		dynamicLevel.SetLevel(zap.InfoLevel)

	case "warn":
		dynamicLevel.SetLevel(zap.WarnLevel)
	case "error":
		dynamicLevel.SetLevel(zap.ErrorLevel)
	}

	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	})
	cores := [...]zapcore.Core{
		zapcore.NewCore(encoder, os.Stdout, dynamicLevel),
		zapcore.NewCore(
			encoder,
			zapcore.AddSync(&lumberjack.Logger{
				Filename:   g.Config.Logger.SavePath,
				MaxSize:    g.Config.Logger.MaxSize,
				MaxBackups: g.Config.Logger.MaxBackups,
				LocalTime:  true,
				Compress:   g.Config.Logger.IsCompress,
			}),
			dynamicLevel,
		),
	}
	g.Logger = zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())
	defer func(Logger *zap.Logger) {
		_ = Logger.Sync()
	}(g.Logger)

	g.Logger.Info("initialize logger successfully!")

}

func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("[2006-01-02 15:04:05.000"))
}

func getWriteSyncer(file string) zapcore.WriteSyncer {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   file,
		MaxSize:    g.Config.Logger.MaxSize,
		MaxAge:     g.Config.Logger.MaxAge,
		MaxBackups: g.Config.Logger.MaxBackups,
		Compress:   true,
	}
	return zapcore.AddSync(lumberjackLogger)
}
