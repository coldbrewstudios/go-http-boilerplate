package logging

import (
	"go.uber.org/zap"
)

type AppLogger interface {
	Info(msg string, keysAndValues ...interface{})
}

type logInstance struct {
	instance *zap.SugaredLogger
}

func CreateLogger(env string) *logInstance {
	zapLogger := NewZapSugarLogger(env)
	return &logInstance{
		instance: zapLogger,
	}
}

func (l *logInstance) Info(msg string, keysAndValues ...interface{}) {
	l.instance.Infow(msg, keysAndValues...)
}

func (l *logInstance) Error(msg string, keysAndValues ...interface{}) {
	l.instance.Errorw(msg, keysAndValues...)
}

// Mocks
type FakeLogger struct {
	instance fakeZapLogger
}

type fakeZapLogger interface {
	Infow(msg string, keysAndValues ...interface{})
}

func (l *FakeLogger) Info(msg string, keysAndValues ...interface{}) {
	l.instance.Infow(msg, keysAndValues...)
}
