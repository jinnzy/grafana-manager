package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)



var log *zap.SugaredLogger
//var log *zap.Logger

type Field = zapcore.Field
type Config struct {
	RunMode string
}

// 时间格式
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02 15:04:05.000"))
}

// zap日志配置信息
func newEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "app",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     timeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func Init(conf *Config) {

	var level zapcore.Level

	switch conf.RunMode {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	config := zap.Config{
		Level: zap.NewAtomicLevelAt(level),
		Development: false,
		DisableStacktrace: true,
		Encoding: "console", // json / console
		EncoderConfig: newEncoderConfig(),
		InitialFields: map[string]interface{}{"MyName": "kainhuck"},
		OutputPaths: []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
	}
	logger, err := config.Build(zap.AddCaller(), zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	log = logger.Sugar()

}

// 使用printf 格式化
func Debug(args ...interface{})  {
	log.Debug(args)
}
func Debugf(template string, args ...interface{})  {
	log.Debugf(template, args...)
}
func Debugw(template string, args ...interface{})  {
	log.Debugw(template, args...)
}
func Info(args ...interface{}) {
	log.Info(args...)
}

func Infof(template string, args ...interface{}) {
	log.Infof(template, args...)
}

func Infow(template string, args ...interface{}) {
	log.Infow(template, args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	log.Warnf(template, args...)
}

func Warnw(template string, args ...interface{}) {
	log.Warnw(template, args...)
}

func Error(args ... interface{}) {
	log.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	log.Errorf(template, args...)
}

func Errorw(template string, args ...interface{}) {
	log.Errorw(template, args...)
}

func Panic(args ...interface{}) {
	log.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	log.Panicf(template, args...)
}

func Panicw(template string, args ...interface{}) {
	log.Panicw(template, args...)
}

func DPanic(args ...interface{}) {
	log.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	log.DPanicf(template, args...)
}

func DPanicw(template string, args ...interface{}) {
	log.DPanicw(template, args...)
}

func Fatal(args ...interface{})  {
	log.Fatal(args)
}

func Fatalf(template string, args ...interface{})  {
	log.Fatalf(template, args...)
}

func Fatalw(template string, args ...interface{})  {
	log.Fatalw(template, args...)
}
