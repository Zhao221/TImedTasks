package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const DateTimeFormat = "2006-01-02 15:04:05"

// 原始日志
var zc *zap.Logger

// sugar日志
var z *zap.SugaredLogger

// ZapC 原始日志对象
func ZapC() *zap.Logger {
	return zc
}

// Zap sugar日志对象
func Zap() *zap.SugaredLogger {
	return z
}

// WithContext  从context中取得带request-id的sugar日志对象，必须使用middleware_gin.RequestId()
func WithContext(c context.Context) *zap.SugaredLogger {
	zt := c.Value("zap_trace")
	if zt == nil {
		return z
	}

	if v, ok := zt.(*zap.SugaredLogger); ok {
		return v
	}

	return z
}

// 初始化zap日志
func init() {
	// 是否打印堆栈信息
	stacktraceKey := "stacktrace"
	// if env.InProd() {
	//     stacktraceKey = ""
	// }

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  stacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout(DateTimeFormat),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    encoderConfig,
		InitialFields:    map[string]interface{}{},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	// 构建日志对象
	l, err := config.Build()
	if err != nil {
		panic(err)
	}

	// 原始日志
	zc = l

	// sugar日志
	z = l.Sugar()
}
