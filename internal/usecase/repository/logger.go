package usecase_repository

type Logger interface {
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
	Errorf(string, ...interface{})
	Infof(string, ...interface{})
	WithField(string, interface{}) Logger
	Printf(string, ...interface{})
}
