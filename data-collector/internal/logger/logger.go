package logger

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
}

var (
	logger *Logger
)

func Instance() *Logger {
	if logger != nil { // more than likely that the other loggers are also initialized
		return logger
	}
	file, err := os.OpenFile("data-collector.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	mw := io.MultiWriter(os.Stdout, file)
	logger := Logger{
		infoLogger:    log.New(mw, "INFO: ", log.Ldate|log.Ltime),
		warningLogger: log.New(mw, "WARNING: ", log.Ldate|log.Ltime),
		errorLogger:   log.New(mw, "ERROR: ", log.Ldate|log.Ltime),
	}
	return &logger

}

func (l *Logger) Info(message string) {
	l.infoLogger.Println(message)
}

func (l *Logger) Warning(message string) {
	l.warningLogger.Println(message)
}

func (l *Logger) Error(message string) {
	l.errorLogger.Println(message)
}
