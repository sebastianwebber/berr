package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/charmbracelet/log"
	"github.com/sebastianwebber/berr"
)

var (
	simpleError = berr.New("simple error")

	// since berr.betterError implements the Error interface, it can
	// be used as a normal error.
	complexError        = fmt.Errorf("complex error: %w", simpleError)
	veryComplexError    = fmt.Errorf("very complex error: %w", complexError)
	ultraComplexError   = fmt.Errorf("ultra complex error: %w", veryComplexError)
	godLikeComplexError = fmt.Errorf("god like complex error: %w", ultraComplexError)
	abortError          = fmt.Errorf("need to abort: %w", complexError)

	examples = map[string]error{
		"simple error":           simpleError,
		"complex error":          complexError,
		"very complex error":     veryComplexError,
		"ultra complex error":    ultraComplexError,
		"god like complex error": godLikeComplexError,
		"join error":             errors.Join(simpleError, complexError),
	}
)

func main() {
	defer endFunc()
	for k, v := range examples {
		logger(v, "example", k).Info("message")
		time.Sleep(1 * time.Second)
	}
}

func endFunc() {
	berr.Options.PrintStack = true
	// you could only use the formatter if you want
	log.Fatal("finishing with an error and stack trace", "details", berr.Format(abortError))
}

func logger(err error, args ...any) *log.Logger {
	args = append(args, "error", berr.Format(err))
	return log.With(args...)
}
