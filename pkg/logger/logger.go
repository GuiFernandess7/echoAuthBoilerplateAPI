package logger

import (
	"io"
	"log"
	"os"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

type Logger struct {
	*log.Logger
	level LogLevel
}

func NewLogger(path string, level LogLevel) (*Logger, error) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	multiWriter := io.MultiWriter(os.Stdout, file)
	logger := log.New(multiWriter, "", log.LstdFlags)

	return &Logger{
		Logger: logger,
		level:  level,
	}, nil
}

func (l *Logger) Print(format string, v ...interface{}) {
	if l.level <= INFO {
		l.Logger.Printf(format, v...)
	}
}

func (l *Logger) Info(format string, v ...interface{}) {
	if l.level <= INFO {
		l.Logger.Printf(format, v...)
	}
}

func (l *Logger) Error(format string, v ...interface{}) {
	if l.level <= ERROR {
		l.Logger.Printf("[ERROR] "+format, v...)
	}
}

func (l *Logger) Fatal(format string, v ...interface{}) {
	if l.level <= FATAL {
		l.Logger.Printf("[FATAL] "+format, v...)
		os.Exit(1)
	}
}
