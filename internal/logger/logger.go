package logger

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

type LogLevel int

const (
	LogLevelTrace LogLevel = -8
	LogLevelDebug LogLevel = -4
	LogLevelInfo  LogLevel = 0
	LogLevelWarn  LogLevel = 4
	LogLevelError LogLevel = 8
)

type Logger struct {
	Level LogLevel
}

func NewLogger() Logger {
	// default level is 0
	return Logger{}
}

func (l LogLevel) String() string {
	var s string
	switch l {
	case LogLevelTrace:
		s = "TRACE"
	case LogLevelDebug:
		s = "DEBUG"
	case LogLevelInfo:
		s = "INFO"
	case LogLevelWarn:
		s = "WARN"
	case LogLevelError:
		s = "ERROR"
	default:
		s = ""
	}

	return s
}

func (l *Logger) Output(callDepth int, level LogLevel, msg string) {
	if level < l.Level {
		return
	}

	timeString := time.Now().Format("2006-01-02 03:04:05")

	levelString := level.String()

	// Info log do not include file
	if level == LogLevelInfo {
		fmt.Fprintf(os.Stdout, "[%s] %s %s\n", timeString, levelString, msg)
		return
	}

	var file string
	var line int
	_, file, line, ok := runtime.Caller(callDepth)
	if !ok {
		file = "???"
		line = 0
	}

	var short string
	// Steal from std log
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}

	if level > 0 {
		fmt.Fprintf(os.Stderr, "[%s] %s %s:%d %s\n", timeString, levelString, short, line, msg)
	} else {
		fmt.Fprintf(os.Stdout, "[%s] %s %s:%d %s\n", timeString, levelString, short, line, msg)
	}
}

func (l *Logger) Error(msg string) {
	l.Output(2, LogLevelError, msg)
}

func (l *Logger) Warn(msg string) {
	l.Output(2, LogLevelWarn, msg)
}

func (l *Logger) Info(msg string) {
	l.Output(2, LogLevelInfo, msg)
}

func (l *Logger) Debug(msg string) {
	l.Output(2, LogLevelDebug, msg)
}

func (l *Logger) Trace(msg string) {
	l.Output(2, LogLevelTrace, msg)
}
