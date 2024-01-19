package infrastructure

import (
	"time"

	"github.com/lodashventure/nlp/model"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	Call *logrus.Logger
}

func NewLogger() *Logger {
	log := logrus.New()
	log.Formatter = &logrus.JSONFormatter{}

	return &Logger{Call: log}
}

func (l *Logger) Fatalln(logFiber *model.LogFiber, prefix, appName, detail string, err error) {
	l.Call.WithFields(logrus.Fields{
		"prefix":   prefix,
		"app_name": appName,
	}).WithError(err).Fatalln(detail)
	go l.saveLog(logFiber, prefix, appName, detail, err)
}

func (l *Logger) Panicln(logFiber *model.LogFiber, prefix, appName, detail string, err error) {
	l.Call.WithFields(logrus.Fields{
		"prefix":   prefix,
		"app_name": appName,
	}).WithError(err).Panicln(detail)
	go l.saveLog(logFiber, prefix, appName, detail, err)
}

func (l *Logger) Error(logFiber *model.LogFiber, prefix, appName, detail string, err error) {
	l.Call.WithFields(logrus.Fields{
		"prefix":   prefix,
		"app_name": appName,
	}).WithError(err).Error(detail)
	go l.saveLog(logFiber, prefix, appName, detail, err)
}

func (l *Logger) Info(prefix, appName, detail string) {
	l.Call.WithFields(logrus.Fields{
		"prefix":   prefix,
		"app_name": appName,
	}).Infoln(detail)
}

func (l *Logger) saveLog(logFiber *model.LogFiber, prefix, appName, detail string, err error) {
	timeNow := time.Now()
	logError := &model.LogError{
		Time:      timeNow.UTC().Format("2006-01-02 15:04:05"),
		Message:   detail,
		Error:     err.Error(),
		Path:      logFiber.Path,
		IP:        logFiber.IP,
		AppName:   logFiber.AppName,
		TokenID:   logFiber.TokenID,
		CreatedAt: timeNow.Format(time.RFC3339),
	}

	err = LogErrorServiceManage.Insert(logError)
	if err != nil {
		Log.Error(logFiber, "LogRestHTTP", logError.AppName, "Can't insert data log http in database", err)
	}
}
