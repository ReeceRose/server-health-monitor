package logger

import (
	"io"
	"log"
	"os"

	"github.com/PR-Developers/server-health-monitor/internal/consts"
	"github.com/PR-Developers/server-health-monitor/internal/utils"
	"github.com/PR-Developers/server-health-monitor/internal/wrapper"
)

// Logger is an interface which provides method signatures for logging to files/console
type Logger interface {
	Info(msg string)
	Infof(msg string, args ...interface{})
	Warning(msg string)
	Warningf(msg string, args ...interface{})
	Error(msg string)
	Errorf(msg string, args ...interface{})
	Logger() *log.Logger
}

type standardLogger struct {
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
	genericLogger *log.Logger
}

var (
	logger    *standardLogger
	_         Logger                  = (*standardLogger)(nil)
	osWrapper wrapper.OperatingSystem = &wrapper.DefaultOS{}
)

// Instance returns the active instance of the logger
func Instance() Logger {
	if logger != nil {
		return logger
	}
	file, err := osWrapper.OpenFile(utils.GetVariable(consts.LOG_FILE), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil
	}
	mw := io.MultiWriter(os.Stdout, file)
	logger = &standardLogger{
		infoLogger:    log.New(mw, "INFO: ", log.Ldate|log.Ltime),
		warningLogger: log.New(mw, "WARNING: ", log.Ldate|log.Ltime),
		errorLogger:   log.New(mw, "ERROR: ", log.Ldate|log.Ltime),
		genericLogger: log.New(mw, "", log.Ldate|log.Ltime),
	}

	return logger
}

// Info should be used to log generic log messages
func (l *standardLogger) Info(message string) {
	l.infoLogger.Println(message)
}

// Infof should be used to log generic log messages with special formatting
func (l *standardLogger) Infof(format string, args ...interface{}) {
	l.infoLogger.Printf(format+"\n", args...)
}

// Warning should be used to log events of concern
func (l *standardLogger) Warning(message string) {
	l.warningLogger.Println(message)
}

// Warningf should be used to log events of concern with special formatting
func (l *standardLogger) Warningf(format string, args ...interface{}) {
	l.warningLogger.Printf(format+"\n", args...)
}

// Error should be used to log unexpected behaviour
func (l *standardLogger) Error(message string) {
	l.errorLogger.Println(message)
}

// Errorf should be used to log unexpected behaviour with special formatting
func (l *standardLogger) Errorf(format string, args ...interface{}) {
	l.errorLogger.Printf(format+"\n", args...)
}

// Logger returns a generic instance of a default logger
func (l *standardLogger) Logger() *log.Logger {
	return l.genericLogger
}
