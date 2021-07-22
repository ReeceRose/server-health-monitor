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
	Info(string)
	Warning(string)
	Error(string)
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

// Warning should be used to log events of concern
func (l *standardLogger) Warning(message string) {
	l.warningLogger.Println(message)
}

// Error should be used to log unexpected behaviour
func (l *standardLogger) Error(message string) {
	l.errorLogger.Println(message)
}

// Logger returns a generic instance of a default logger
func (l *standardLogger) Logger() *log.Logger {
	return l.genericLogger
}
