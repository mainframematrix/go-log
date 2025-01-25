package logger

import (
	"fmt"
	"log"
	"os"
)

const (
	DEBUG = iota
	INFO
	ERROR
)

type Logger struct {
	*log.Logger
	level int

	timeEnable bool
}

func CreateLogger(level int, output string, timeEnable bool) (*Logger, error) {
	var outputFile *os.File
	if output == "stdout" {
		outputFile = os.Stdout
	} else if output == "stderr" {
		outputFile = os.Stderr
	} else {
		var err error
		outputFile, err = os.OpenFile(output, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			return nil, fmt.Errorf("error opening output file %s: %v", output, err)
		}
	}
	out := log.New(outputFile, "", 0)
	return &Logger{
		out,
		level,

		timeEnable,
	}, nil
}

func (l *Logger) Log(level int, format string, args ...interface{}) {
	if level >= l.level {
		levelStr := ""
		switch level {
		case INFO:
			levelStr = "INFO"
		case DEBUG:
			levelStr = "DEBUG"
		case ERROR:
			levelStr = "ERROR"
		}
		if l.timeEnable {
			l.Logger.SetFlags(log.LstdFlags)
		}
		msg := fmt.Sprintf(format, args...)
		l.Logger.Printf("[%s] %s\n", levelStr, msg)
	}
}
