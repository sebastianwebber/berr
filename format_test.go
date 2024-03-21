package berr

import (
	"errors"
	"fmt"
	"strings"

	"github.com/stretchr/testify/assert"
)

var (
	simpleError         = errors.New("simple error")
	complexError        = fmt.Errorf("complex error: %w", simpleError)
	veryComplexError    = fmt.Errorf("very complex error: %w", complexError)
	ultraComplexError   = fmt.Errorf("ultra complex error: %w", veryComplexError)
	godLikeComplexError = fmt.Errorf("god like complex error: %w", ultraComplexError)
)

func (suite *betterErrorsTestSuite) TestFormat() {

	cases := []struct {
		name string
		args error
		want string
	}{
		{"simple error", simpleError, "simple error"},
		{
			"nested in two levels",
			complexError,
			strings.Join([]string{
				"complex error",
				"",
				"caused by:",
				"   0: simple error",
			}, "\n"),
		},
		{
			"nested in three levels",
			veryComplexError,
			strings.Join([]string{
				"very complex error",
				"",
				"caused by:",
				"   0: complex error",
				"   1: simple error",
			}, "\n"),
		},
		{
			"nested in four levels",
			ultraComplexError,
			strings.Join([]string{
				"ultra complex error",
				"",
				"caused by:",
				"   0: very complex error",
				"   1: complex error",
				"   2: simple error",
			}, "\n"),
		},
		{
			"nested in five levels",
			godLikeComplexError,
			strings.Join([]string{
				"god like complex error",
				"",
				"caused by:",
				"   0: ultra complex error",
				"   1: very complex error",
				"   2: complex error",
				"   3: simple error",
			}, "\n"),
		},
		{
			"join errors",
			errors.Join(simpleError, complexError),
			strings.Join([]string{
				"simple error",
				"",
				"complex error",
				"",
				"caused by:",
				"   0: simple error",
			}, "\n"),
		},
	}
	for _, c := range cases {
		out := Format(c.args)
		suite.Run(c.name, func() {
			assert.Equal(suite.T(), c.want, out)
		})
	}
}
