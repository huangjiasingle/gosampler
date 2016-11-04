package logger

import (
	"fmt"
	l "log"
	"os"
	"runtime"
)

const (
	levelDebug = "[DEBUG]"
	levelInfo  = "[INFO]"
	levelWarn  = "[WARN]"
	levelError = "[ERROR]"
	levelFatal = "[FATAL]"
	levelExit  = "[EXIT]"
)

type new struct {
	name string
}

func New(name string) *new {
	return &new{name}
}

// ------------------
// Debug

func (l *new) Debug(v ...interface{}) {
	plog(levelDebug, l.name, v...)
}

func (l *new) Debugf(format string, v ...interface{}) {
	plogf(levelDebug, format, l.name, v...)
}

// ------------------
// Info

func (l *new) Info(v ...interface{}) {
	plog(levelInfo, l.name, v...)
}

func (l *new) Infof(format string, v ...interface{}) {
	plogf(levelInfo, format, l.name, v...)
}

// ------------------
// Warn

func (l *new) Warn(v ...interface{}) {
	plog(levelWarn, l.name, v...)
}

func (l *new) Warnf(format string, v ...interface{}) {
	plogf(levelWarn, format, l.name, v...)
}

// ------------------
// Error

func (l *new) Error(v ...interface{}) {
	plog(levelError, l.name, v...)
}

func (l *new) Errorf(format string, v ...interface{}) {
	plogf(levelError, format, l.name, v...)
}

// ------------------
// Fatal

func (l *new) Fatal(v ...interface{}) {
	plog(levelFatal, l.name, v...)
	os.Exit(1)
}

func (l *new) Fatalf(format string, v ...interface{}) {
	plogf(levelFatal, format, l.name, v...)
	os.Exit(1)
}

// ------------------
// Exit

func (l *new) Exit(v ...interface{}) {
	plog(levelExit, l.name, v...)
	os.Exit(0)
}

func (l *new) Exitf(format string, v ...interface{}) {
	plogf(levelExit, format, l.name, v...)
	os.Exit(0)
}

// ------------------
// plog

func plog(level, name string, v ...interface{}) {
	plogf(level, "", name, v...)
}

func plogf(level, format, name string, v ...interface{}) {
	if level == levelInfo {
		if format != "" {
			l.Print(level + " " + name + " " + fmt.Sprintf(format, v...))
		} else {
			l.Print(level + " " + name + " " + fmt.Sprintln(v...))
		}
	} else {
		_, file, line, _ := runtime.Caller(2)
		loc := fmt.Sprintf("[%s:%d]", file, line)
		msg := fmt.Sprintln(v...)
		if len(msg) > 0 {
			msg = msg[0 : len(msg)-1]
		}
		if format != "" {
			l.Print(level + " " + name + " " + fmt.Sprintf(format, v...) + " " + loc)
		} else {
			l.Print(level + " " + name + " " + msg + " " + loc)
		}
	}
}
