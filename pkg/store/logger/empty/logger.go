package empty

type emptyLogger struct{}

func NewLogger() *emptyLogger {
	return &emptyLogger{}
}

// It does not log any error messages or context.
func (l *emptyLogger) Error(err error, msg string, kvs ...any) {
	// No operation performed for logging errors
}
