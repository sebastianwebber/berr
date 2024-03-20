package berr

import (
	"errors"
	"fmt"
	"strings"
)

// New returns a Better Error
func New(err error) error {
	return parse(err)
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

func parse(err error) betterError {
	if err == nil {
		return betterError{
			errMsg:      "",
			originalErr: err,
		}
	}

	return parseUnwrap(err)
}

// parseUnwrap is a helper function to convert an Unwrap error into a prettyError
func parseUnwrap(err error) betterError {
	if _, supportsUnwrap := err.(interface {
		Unwrap() error
	}); !supportsUnwrap {
		return betterError{
			originalErr: err,
			errMsg:      err.Error(),
		}
	}

	var causes []error

	e := errors.Unwrap(err)
	for {
		if e == nil {
			break
		}

		causes = append(causes, e)
		e = errors.Unwrap(e)
	}

	errMsg := err.Error()

	for _, c := range causes {
		errMsg = strings.ReplaceAll(errMsg, fmt.Sprintf(": %s", c.Error()), "")
	}

	return betterError{
		originalErr: err,
		fullMsg:     err.Error(),
		errMsg:      errMsg,
		cause:       causes,
	}
}
