package logging

import (
	"time"

	"github.com/narendra121/pkghub/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func ConfigureLog() {
	cfg := zapcore.EncoderConfig{
		TimeKey:       "timestamp",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		FunctionKey:   zapcore.OmitKey,
		MessageKey:    "message",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		EncodeTime: func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
			pae.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	logCfg := logger.NewLoggerBuilder().
		SetEncoding("json").
		SetLogFilePath([]string{"info.log"}, []string{"error.log"}).
		SetLogLevel(zap.DebugLevel).
		SetEncoderConfig(cfg).
		// SetDisableCaller(true).
		SetDisableStacktrace(true).
		Build()

	Log = logCfg
	defer Log.Sync()
}
