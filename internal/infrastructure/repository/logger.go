package repository

import (
	"context"

	usecase_repository "github.com/Goalt/FileSharer/internal/usecase/repository"
	"gorm.io/gorm/logger"
)

type Logger struct {
	prefix  string
	logCore logger.Interface
}

func NewLogger(logCore logger.Interface) *Logger {
	return &Logger{
		prefix:  "",
		logCore: logCore,
	}
}

func (l *Logger) Info(ctx context.Context, s string, vars ...interface{}) {
	l.logCore.Info(ctx, l.prefix+s, vars...)
}

func (l *Logger) Warn(ctx context.Context, s string, vars ...interface{}) {
	l.logCore.Info(ctx, l.prefix+s, vars...)
}

func (l *Logger) Error(ctx context.Context, s string, vars ...interface{}) {
	l.logCore.Info(ctx, l.prefix+s, vars...)
}

func (l *Logger) WithPrefix(prefix string) usecase_repository.Logger {
	return &Logger{
		prefix:  prefix + " ",
		logCore: l.logCore,
	}
}
