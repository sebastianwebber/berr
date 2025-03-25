package berr

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os/exec"
	"strings"
)

var (
	simpleError         = errors.New("simple error")
	complexError        = fmt.Errorf("complex error: %w", simpleError)
	veryComplexError    = fmt.Errorf("very complex error: %w", complexError)
	ultraComplexError   = fmt.Errorf("ultra complex error: %w", veryComplexError)
	godLikeComplexError = fmt.Errorf("god like complex error: %w", ultraComplexError)
	combinedOutput      bytes.Buffer
	cmdExitError        = func() error {
		cmd := exec.Command("bash", "-c", `echo "message from stdout" && echo "error msg from stderr" >&2 && exit 1 `)
		cmd.Stdout = &combinedOutput
		cmd.Stderr = &combinedOutput

		err := cmd.Run()

		return fmt.Errorf("%s: %w", combinedOutput.String(), err)
	}()
	spacesOutputExitError = func() error {
		cmd := exec.Command("bash", "-c", `echo "      \t message from stdout" && echo "\t \t    error msg from stderr     \n" >&2 && exit 1 `)
		cmd.Stdout = &combinedOutput
		cmd.Stderr = &combinedOutput

		err := cmd.Run()

		return fmt.Errorf("%s: %w", combinedOutput.String(), err)
	}()
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
		{
			"with errors from cmd.Execute",
			fmt.Errorf("error running command: %w", cmdExitError),
			strings.Join([]string{
				"error running command",
				"",
				"caused by:",
				"   0: message from stdout",
				"      error msg from stderr",
				"   1: exit status 1",
			}, "\n"),
		},
		{
			"with spaces and output from cmd.Execute",
			fmt.Errorf("error running command: %w", cmdExitError),
			strings.Join([]string{
				"error running command",
				"",
				"caused by:",
				"   0: message from stdout",
				"      error msg from stderr",
				"   1: exit status 1",
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
