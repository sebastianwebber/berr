package berr

import (
	"bufio"
	"fmt"
	"runtime"
	"strings"
)

var (
	functionsToIgnore = []string{
		"github.com/sebastianwebber/berr.getStack",
		"github.com/sebastianwebber/berr.collectStackTrace",
		"github.com/sebastianwebber/berr.betterError.Error",
		"github.com/sebastianwebber/berr.Format",
	}
)

const (
	// MaxStackDepth is the maximum number of stack frames to retrieve
	// when collecting a stack trace.
	MaxStackDepth = 100
)

// collectStackTrace returns a slice of strings representing the
// current stack trace ignore the calls in the berr package
func collectStackTrace() []string {
	var out []string

	if Options.PrintStack == false {
		return out
	}

	traceLine := "Stack trace (without berr functions):"
	if Options.ShowCompleteStack {
		traceLine = "Stack trace:"
	}

	out = append(out, []string{
		"",
		traceLine,
	}...)

	trace := strings.NewReader(getStack())
	scanner := bufio.NewScanner(trace)
	for scanner.Scan() {
		line := strings.ReplaceAll(scanner.Text(), "\t", "  ")
		out = append(out, fmt.Sprintf("  %s", line))
	}

	return out
}

func getStack() string {
	var out []string
	pc := make([]uintptr, MaxStackDepth)
	n := runtime.Callers(1, pc)
	if n == 0 {
		// Return now to avoid processing the zero Frame that would
		// otherwise be returned by frames.Next below.
		return ""
	}

	pc = pc[:n] // pass only valid pcs to runtime.CallersFrames
	frames := runtime.CallersFrames(pc)

	for {
		frame, more := frames.Next()

		skip := false
		for _, ignore := range functionsToIgnore {
			if strings.Contains(frame.Function, ignore) {
				skip = true
				break
			}
		}

		if !Options.ShowCompleteStack && skip {
			continue
		}

		out = append(out, fmt.Sprintf("%s:\n\t%s:%d", frame.Function, frame.File, frame.Line))

		if !more {
			break
		}
	}

	return strings.Join(out, "\n")
}
