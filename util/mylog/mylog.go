package mylog

// intentionally poorly named to make clear that is an internal logging
// api

// for development, print to stdout.

import (
	"fmt"
	"time"
)

// log messages below thresh will not be let pass
const DEFAULT_SEV_THRESH = 0

// default internal severity levels
const (
	SEV_EMERGENCY     = 0
	SEV_ALERT         = 1
	SEV_CRITICAL      = 2
	SEV_ERROR         = 3
	SEV_WARNING       = 4
	SEV_NOTICE        = 5
	SEV_INFORMATIONAL = 6
	SEV_DEBUG         = 7
	SEV_TRACE         = 8
)

var severityStrings = [...]string{"emergency", "alert", "critical", "error",
	"warning", "notice", "info", "debug"}

type Logger interface {
	SetSeverityThresh(severity int)
	Emergency(format string, a ...interface{})
	Alert(format string, a ...interface{})
	Critical(format string, a ...interface{})
	Error(format string, a ...interface{})
	Warning(format string, a ...interface{})
	Notice(format string, a ...interface{})
	Informational(format string, a ...interface{})
	Debug(format string, a ...interface{})
	Trace(format string, a ...interface{})
}

func getSeverityString(severity int) string {
	return severityStrings[severity]
}

// always return the same and only implementation
func CreateLogger(name string) Logger {
	l := &simple{}
	l.SetSeverityThresh(DEFAULT_SEV_THRESH)
	l.SetName(name)
	return l
}

// simple is an implementation of Logger interface
type simple struct {
	severityThresh int
	name           string
}

func (l *simple) log(severity int, format string, a ...interface{}) {
	if severity > l.severityThresh {
		return // don't process this message
	}
	s := fmt.Sprintf(format, a...)
	severityString := getSeverityString(severity)
	tz := time.Now()
	timeStr := tz.Format(time.RFC3339Nano)
	fmt.Printf("log - %s - %s - %s - %s\n", l.name, timeStr, severityString, s)
}

func (l *simple) SetSeverityThresh(severity int) {
	l.severityThresh = severity
}

func (l *simple) SetName(name string) {
	l.name = name
}

// These interfacey methods just feed into the internal log method. simple.log
// is the internal method that actually does everything.
func (l *simple) Emergency(format string, a ...interface{}) {
	l.log(SEV_EMERGENCY, format, a...)
}
func (l *simple) Alert(format string, a ...interface{}) {
	l.log(SEV_ALERT, format, a...)
}
func (l *simple) Critical(format string, a ...interface{}) {
	l.log(SEV_CRITICAL, format, a...)
}
func (l *simple) Error(format string, a ...interface{}) {
	l.log(SEV_ERROR, format, a...)
}
func (l *simple) Warning(format string, a ...interface{}) {
	l.log(SEV_WARNING, format, a...)
}
func (l *simple) Notice(format string, a ...interface{}) {
	l.log(SEV_NOTICE, format, a...)
}
func (l *simple) Informational(format string, a ...interface{}) {
	l.log(SEV_INFORMATIONAL, format, a...)
}
func (l *simple) Debug(format string, a ...interface{}) {
	l.log(SEV_DEBUG, format, a...)
}
func (l *simple) Trace(format string, a ...interface{}) {
	l.log(SEV_TRACE, format, a...)
}
