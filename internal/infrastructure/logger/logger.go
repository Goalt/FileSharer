package logger

import (
	"github.com/Goalt/FileSharer/internal/config"
	usecase_repository "github.com/Goalt/FileSharer/internal/usecase/repository"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Entry
}

func NewLogger(cnf config.Logger) *Logger {
	logrusLog := logrus.New()
	if cnf.SetReportCaller {
		logrusLog.ReportCaller = true
	}

	switch cnf.Level {
	case config.ErrorLevel:
		logrusLog.SetLevel(logrus.ErrorLevel)
	case config.WarnLevel:
		logrusLog.SetLevel(logrus.WarnLevel)
	case config.InfoLevel:
		logrusLog.SetLevel(logrus.InfoLevel)
	default:
		panic("wrong log level")
	}

	entry := logrus.NewEntry(logrusLog)

	return &Logger{entry}
}

func (l *Logger) WithField(key string, value interface{}) usecase_repository.Logger {
	newEntry := l.Entry.WithField(key, value)

	return &Logger{newEntry}
}
