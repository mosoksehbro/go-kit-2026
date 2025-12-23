package utils

import (
	"log"
	"os"
)

type Logger struct {
	info  *log.Logger
	error *log.Logger
}

func NewLogger() *Logger {
	infoFile, _ := os.OpenFile("logs/app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	errFile, _ := os.OpenFile("logs/error.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	return &Logger{
		info:  log.New(infoFile, "[INFO] ", log.LstdFlags|log.Lshortfile),
		error: log.New(errFile, "[ERROR] ", log.LstdFlags|log.Lshortfile),
	}
}

func (l *Logger) Info(msg string, data map[string]interface{}) {
	l.info.Printf("%s | %+v", msg, data)
}

func (l *Logger) Error(msg string, data map[string]interface{}) {
	l.error.Printf("%s | %+v", msg, data)
}
