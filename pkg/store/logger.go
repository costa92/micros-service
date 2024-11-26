package store

type Logger interface {
	// Log logs a message at the specified level.
	Error(err error, message string, kvs ...any)
}
