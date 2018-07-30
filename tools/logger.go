package tools

import (
	"log"
	"os"
	"io"
)

type LoggerInterface interface {
}

type Logger struct {
}

const (
	DEBUG   = "DEBUG"
	INFO    = "INFO"
	NOTICE  = "NOTICE"
	WARNING = "WARNING"
)

const LOGFILEPATH = "storage/log/gin.log"

var (
	logger *log.Logger
)

func (l *Logger) InitLogger() {
	file, err := os.OpenFile(LOGFILEPATH, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file:", err)
	}

	logger = log.New(io.MultiWriter(file), "INFO:", log.Ldate|log.Lmicroseconds)
}

func (l *Logger) Info(message string) {
	l.addRecord(INFO, message)
}

func (l *Logger) Error(message string) {
	l.addRecord(INFO, message)
}

func (l *Logger) addRecord(levle string, message string) {
	logger.Println(message)
}