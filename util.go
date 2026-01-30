package logger

import (
	"fmt"
	"os"
	"path"
)

func check(s string) {
	if logger == nil {
		panic(fmt.Sprintf("crdx.org/logger.%s called before init", s))
	}
}

func getProgramName() string {
	return path.Base(os.Args[0])
}

// stringify converts a value to a string. It accepts error and string types,
// and panics with a clear message for anything else.
func stringify(value any) string {
	switch value := value.(type) {
	case error:
		return value.Error()
	case string:
		return value
	default:
		panic(fmt.Sprintf("crdx.org/logger: expected error or string, got %T", value))
	}
}
