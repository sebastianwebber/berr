package berr

type Config struct {
	// PrintStack will print the stack trace of the error
	// but ignores the functions related to the berr package
	PrintStack bool

	// ShowCompleteStack will print the complete stack trace
	// including the functions related to the berr package
	ShowCompleteStack bool
}

var (
	Options = Config{
		PrintStack:        false,
		ShowCompleteStack: false,
	}
)

// Reset sets the default options
func Reset() {
	Options.PrintStack = false
	Options.ShowCompleteStack = false
}
