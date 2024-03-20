package berr

import "strings"

// Format returns a pretty formatted error message.
// Heavily inspired by anyhow output: https://docs.rs/anyhow/latest/anyhow/
func Format(err error) string {
	// if the error is a join error, we want to join the formatted
	// errors and return them as a single string
	if e, fromJoin := err.(interface {
		Unwrap() []error
	}); fromJoin {

		allErrors := e.Unwrap()
		var out []string

		for _, joinErr := range allErrors {
			out = append(out, parse(joinErr).Error())
		}

		return strings.Join(out, "\n\n")
	}

	return parse(err).Error()
}
