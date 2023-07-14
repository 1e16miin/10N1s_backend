package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger methods interface

type Logger interface {
	InitLogger()
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	DPanic(args ...interface{})
	DPanicf(template string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
}

// Logger
type zapLogger struct {
	cfg         zap.Config
	sugarLogger *zap.SugaredLogger
}

// App Logger constructor
func NewZapLogger() *zapLogger {
	zapConfig := zap.NewDevelopmentConfig()
	return &zapLogger{cfg: zapConfig}
}

// Init logger
func (l *zapLogger) InitLogger() {
	logWriter := zapcore.AddSync(os.Stderr)

	var encoderCfg zapcore.EncoderConfig

	if l.cfg.Development {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
	}

	var encoder zapcore.Encoder
	encoderCfg.LevelKey = "LEVEL"
	encoderCfg.CallerKey = "CALLER"
	encoderCfg.TimeKey = "TIME"
	encoderCfg.NameKey = "NAME"
	encoderCfg.MessageKey = "MESSAGE"

	if l.cfg.Encoding == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}

	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(encoder, logWriter, l.cfg.Level)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	l.sugarLogger = logger.Sugar()
	if err := l.sugarLogger.Sync(); err != nil {
		l.sugarLogger.Error(err)
	}
}

// Logger methods

func (l *zapLogger) Debug(args ...interface{}) {
	l.sugarLogger.Debug(args...)
}

func (l *zapLogger) Debugf(template string, args ...interface{}) {
	l.sugarLogger.Debugf(template, args...)
}

func (l *zapLogger) Info(args ...interface{}) {
	l.sugarLogger.Info(args...)
}

func (l *zapLogger) Infof(template string, args ...interface{}) {
	l.sugarLogger.Infof(template, args...)
}

func (l *zapLogger) Warn(args ...interface{}) {
	l.sugarLogger.Warn(args...)
}

func (l *zapLogger) Warnf(template string, args ...interface{}) {
	l.sugarLogger.Warnf(template, args...)
}

func (l *zapLogger) Error(args ...interface{}) {
	l.sugarLogger.Error(args...)
}

func (l *zapLogger) Errorf(template string, args ...interface{}) {
	l.sugarLogger.Errorf(template, args...)
}

func (l *zapLogger) DPanic(args ...interface{}) {
	l.sugarLogger.DPanic(args...)
}

func (l *zapLogger) DPanicf(template string, args ...interface{}) {
	l.sugarLogger.DPanicf(template, args...)
}

func (l *zapLogger) Panic(args ...interface{}) {
	l.sugarLogger.Panic(args...)
}

func (l *zapLogger) Panicf(template string, args ...interface{}) {
	l.sugarLogger.Panicf(template, args...)
}

func (l *zapLogger) Fatal(args ...interface{}) {
	l.sugarLogger.Fatal(args...)
}

func (l *zapLogger) Fatalf(template string, args ...interface{}) {
	l.sugarLogger.Fatalf(template, args...)
}
