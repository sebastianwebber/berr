package berr

import (
	"errors"
	"fmt"
	"strings"
)

// Parse converts a error into a betterError
func Parse(err error) betterError {
	return parse(err)
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
