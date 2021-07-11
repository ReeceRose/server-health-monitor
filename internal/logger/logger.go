package logger

import (
	"io"
	"log"
	"os"
)

type Logger interface {
	Info(string)
	Warning(string)
	Error(string)
}

type StandardLogger struct {
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
	GenericLogger *log.Logger
}

var (
	logger *StandardLogger
	_      Logger = (*StandardLogger)(nil)
)

func Instance() *StandardLogger {
	if logger != nil { // more than likely that the other loggers are also initialized
		return logger
	}
	file, err := os.OpenFile("server-health-monitor.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	mw := io.MultiWriter(os.Stdout, file)
	logger := StandardLogger{
		infoLogger:    log.New(mw, "INFO: ", log.Ldate|log.Ltime),
		warningLogger: log.New(mw, "WARNING: ", log.Ldate|log.Ltime),
		errorLogger:   log.New(mw, "ERROR: ", log.Ldate|log.Ltime),
		GenericLogger: log.New(mw, "", log.Ldate|log.Ltime),
	}
	return &logger

}

func (l *StandardLogger) Info(message string) {
	l.infoLogger.Println(message)
}

func (l *StandardLogger) Warning(message string) {
	l.warningLogger.Println(message)
}

func (l *StandardLogger) Error(message string) {
	l.errorLogger.Println(message)
}
