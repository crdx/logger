package logger

type NullLogger struct{}

func NewNullLogger() *NullLogger {
	return &NullLogger{}
}

func (*NullLogger) Added(string, ...any)   {}
func (*NullLogger) Removed(string, ...any) {}

func (*NullLogger) Succeeded(string, ...any) {}
func (*NullLogger) Failed(string, ...any)    {}

func (*NullLogger) Enabled(string, ...any)  {}
func (*NullLogger) Disabled(string, ...any) {}

func (*NullLogger) Print(string, ...any) {}
func (*NullLogger) Info(string, ...any)  {}
func (*NullLogger) Warn(string, ...any)  {}
func (*NullLogger) Err(string, ...any)   {}
