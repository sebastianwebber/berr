package berr

import (
	"errors"
	"fmt"
	"strings"
)

// New returns a Better Error with a interface
// compatible with the errors.New()
func New(text string) error {
	return parse(errors.New(text))
}

// betterError is a struct to hold the error message and the cause of the error
type betterError struct {
	fullMsg     string
	errMsg      string
	originalErr error
	cause       []error
}

// Error implements the error interface
func (p betterError) Error() string {
	if len(p.cause) == 0 {
		return p.errMsg
	}

	output := []string{
		p.errMsg,
		"",
		"caused by:",
	}

	for l := 0; l < len(p.cause); l++ {
		c := parse(p.cause[l])
		output = append(output, fmt.Sprintf("  %2d: %s", l, c.errMsg))
	}

	output = append(output, collectStackTrace()...)

	return strings.Join(output, "\n")
}

func (p betterError) Unwrap() []error {
	// if the error is a join error, we want to join
	// the formatted errors
	if e, fromJoin := p.originalErr.(interface {
		Unwrap() []error
	}); fromJoin {

		allErrors := e.Unwrap()
		var out []error

		for _, joinErr := range allErrors {
			out = append(out, parse(joinErr))
		}

		return out
	}

	return []error{p.originalErr}
}
