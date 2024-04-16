package logger

import (
	"go.uber.org/zap"
)

type Logger interface {
	Named(string) Logger
	Debug(fields ...any)
	Debugf(msg string, fields ...any)
	Info(fields ...any)
	Infof(msg string, fields ...any)
	Warn(fields ...any)
	Warnf(msg string, fields ...any)
	Error(fields ...any)
	Errorf(msg string, fields ...any)
	Fatal(fields ...any)
	Fatalf(msg string, fields ...any)
	Sync() error
}

type logger struct {
	*zap.SugaredLogger
}

func (l *logger) Named(name string) Logger {
	return &logger{
		SugaredLogger: l.SugaredLogger.Named(name),
	}
}

func New() (Logger, error) {
	l, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return &logger{
		SugaredLogger: l.Sugar(),
	}, nil
}
