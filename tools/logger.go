package tools

import (
	"log"
	"os"
	"io/ioutil"
	"io"
)

type Logger struct {
}

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func (l *Logger) InitLogger() {
	file, err := os.OpenFile("storage/log/gin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalln("Failed to open log file:", err)
	}

	Trace = log.New(ioutil.Discard, "TRACK:", log.Ldate|log.Lmicroseconds)
	Info = log.New(os.Stdout, "INFO:", log.Ldate|log.Lmicroseconds)
	Warning = log.New(os.Stdout, "WARING:", log.Ldate|log.Lmicroseconds)
	Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR:", log.Ldate|log.Lmicroseconds)
}

func (l *Logger) Error(message string) {
	Error.Println(message)
}