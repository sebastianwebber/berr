package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/charmbracelet/log"
	"github.com/sebastianwebber/berr"
)

func main() {

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

	for k, v := range examples {
		berr.Logger(v, "example", k).Info("message")
		time.Sleep(1 * time.Second)
	}

	// you could only use the formatter if you want
	log.Fatal("finishing with an error", "details", berr.Format(abortError))
}
