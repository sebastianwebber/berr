package main

import (
	"errors"
	"fmt"

	"github.com/sebastianwebber/berr"
)

func main() {

	var (
		simpleError         = errors.New("simple error")
		complexError        = fmt.Errorf("complex error: %w", simpleError)
		veryComplexError    = fmt.Errorf("very complex error: %w", complexError)
		ultraComplexError   = fmt.Errorf("ultra complex error: %w", veryComplexError)
		godLikeComplexError = fmt.Errorf("god like complex error: %w", ultraComplexError)

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
	}
}
