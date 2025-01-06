package logger

import (
	"os"
	"tudo-thrfting-clothes-server/pkg/setting"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config setting.LoggerSetting) *LoggerZap {
	logLevel := config.LogLevel
	// debug -> info -> warn -> error -> dpanic -> panic -> fatal
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
	encoder := configEncode()
	hook := lumberjack.Logger{
		Filename:   config.FileName,
		MaxSize:    config.MaxSize, // megabytes
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,   //days
		Compress:   config.Compress, // disabled by default
	}

	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout),
			zapcore.AddSync(&hook),
		),
		level,
	)

	return &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))}

}

func configEncode() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	// Parse time to dd-mm-yyyy hh:mm:ss
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// Info to INFO with color
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// Caller to full path
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	// change time key to time
	encodeConfig.TimeKey = "time"

	return zapcore.NewJSONEncoder(encodeConfig)
}
