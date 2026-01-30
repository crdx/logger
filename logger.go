package logger

import "os"

type Logger interface {
	Added(v string, args ...any)
	Removed(v string, args ...any)
	Enabled(v string, args ...any)
	Disabled(v string, args ...any)
	Succeeded(v string, args ...any)
	Failed(v string, args ...any)
	Print(v string, args ...any)
	Info(v string, args ...any)
	Warn(v string, args ...any)
	Err(v string, args ...any)
}

var logger Logger

type Destination int

const (
	DestinationSyslog Destination = iota + 1
	DestinationStderr
	DestinationNull
)

func InitStderr() { Init(DestinationStderr) }
func InitSyslog() { Init(DestinationSyslog) }
func InitNull()   { Init(DestinationNull) }

func Init(destination Destination) {
	switch destination {
	case DestinationSyslog:
		logger = NewSyslogLogger()
	case DestinationStderr:
		logger = NewStderrLogger()
	case DestinationNull:
		logger = NewNullLogger()
	default:
		panic("crdx.org/logger.Init called with invalid destination")
	}
}

// Added logs a message prefixed with [+].
func Added(v string, args ...any) {
	check("Added")
	logger.Added(v, args...)
}

// Removed logs a message prefixed with [-].
func Removed(v string, args ...any) {
	check("Removed")
	logger.Removed(v, args...)
}

// Enabled logs a message prefixed with [✓].
func Enabled(v string, args ...any) {
	check("Enabled")
	logger.Enabled(v, args...)
}

// Disabled logs a message prefixed with [𐄂].
func Disabled(v string, args ...any) {
	check("Disabled")
	logger.Disabled(v, args...)
}

// Succeeded logs a message prefixed with [✓].
func Succeeded(v string, args ...any) {
	check("Succeeded")
	logger.Succeeded(v, args...)
}

// Failed logs a message prefixed with [𐄂].
func Failed(v string, args ...any) {
	check("Failed")
	logger.Failed(v, args...)
}

// Print logs a general message.
func Print(v string, args ...any) {
	check("Print")
	logger.Print(v, args...)
}

// Info logs an informative message.
func Info(v string, args ...any) {
	check("Info")
	logger.Info(v, args...)
}

// Warn logs a warning message.
func Warn(v string, args ...any) {
	check("Warn")
	logger.Warn(v, args...)
}

// Err logs an error message.
func Err(v any, args ...any) {
	check("Err")
	logger.Err(stringify(v), args...)
}

// Fatal logs an error message and then exits.
func Fatal(v any, args ...any) {
	check("Fatal")
	logger.Err(stringify(v), args...)
	os.Exit(1)
}
