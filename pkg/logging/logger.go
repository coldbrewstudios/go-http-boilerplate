package logging

import (
	"go.uber.org/zap"
)

type AppLogger interface {
	Info(msg string, keysAndValues ...interface{})
	Error(msg string, keysAndValues ...interface{})
}

type logger struct {
	instance *zap.SugaredLogger
}

func CreateLogger(logName string) AppLogger {
	NewLogFile(logName)
	zapLogger := NewZapSugarLogger(logName)
	return &logger{
		instance: zapLogger,
	}
}

func (l *logger) Info(msg string, keysAndValues ...interface{}) {
	l.instance.Infow(msg, keysAndValues...)
}

func (l *logger) Error(msg string, keysAndValues ...interface{}) {
	l.instance.Errorw(msg, keysAndValues...)
}
