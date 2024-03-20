package berr

import "github.com/charmbracelet/log"

// Logger returns a *log.Logger with the error field set to the pretty error message
func Logger(err error, args ...any) *log.Logger {
	args = append(args, "error", Format(err))
	return log.With(args...)
}
