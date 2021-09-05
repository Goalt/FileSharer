package usecase_repository

import "context"

type Logger interface {
	Info(context.Context, string, ...interface{})
	Warn(context.Context, string, ...interface{})
	Error(context.Context, string, ...interface{})
	WithPrefix(string) Logger
}
