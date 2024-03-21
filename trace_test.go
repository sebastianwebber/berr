package berr

import (
	"strings"

	"github.com/stretchr/testify/assert"
)

func (suite *betterErrorsTestSuite) TestTraceSupport() {
	for _, c := range traceTestCases {
		suite.Run(c.name, func() {
			Options.PrintStack = c.showTrace
			Options.ShowCompleteStack = c.showCompleteTrace

			out := Format(c.args)
			assert.Equal(suite.T(), c.want, out)

			Reset()
		})
	}
}

var (
	traceTestCases = []struct {
		name              string
		args              error
		want              string
		showTrace         bool
		showCompleteTrace bool
	}{
		{
			"should print trace ignoring berr functions",
			complexError,
			strings.Join([]string{
				"complex error",
				"",
				"caused by:",
				"   0: simple error",
				"",
				"Stack trace (without berr functions):",
				"  github.com/sebastianwebber/berr.(*betterErrorsTestSuite).TestTraceSupport.func1:",
				"    /Users/seba/projetos/github.com/sebastianwebber/berr/trace_test.go:15",
				"  github.com/stretchr/testify/suite.(*Suite).Run.func1:",
				"    /Users/seba/go/pkg/mod/github.com/stretchr/testify@v1.9.0/suite/suite.go:115",
				"  testing.tRunner:",
				"    /nix/store/k9srp8ngvblscg68fdpcyqkydh86429k-go-1.22.1/share/go/src/testing/testing.go:1689",
				"  runtime.goexit:",
				"    /nix/store/k9srp8ngvblscg68fdpcyqkydh86429k-go-1.22.1/share/go/src/runtime/asm_arm64.s:1222",
			}, "\n"),
			true,
			false,
		},
		{
			"should print complete trace",
			complexError,
			strings.Join([]string{
				"complex error",
				"",
				"caused by:",
				"   0: simple error",
				"",
				"Stack trace:",
				"  github.com/sebastianwebber/berr.getStack:",
				"    /Users/seba/projetos/github.com/sebastianwebber/berr/trace.go:53",
				"  github.com/sebastianwebber/berr.collectStackTrace:",
				"    /Users/seba/projetos/github.com/sebastianwebber/berr/trace.go:38",
				"  github.com/sebastianwebber/berr.betterError.Error:",
				"    /Users/seba/projetos/github.com/sebastianwebber/berr/berr.go:40",
				"  github.com/sebastianwebber/berr.Format:",
				"    /Users/seba/projetos/github.com/sebastianwebber/berr/format.go:24",
				"  github.com/sebastianwebber/berr.(*betterErrorsTestSuite).TestTraceSupport.func1:",
				"    /Users/seba/projetos/github.com/sebastianwebber/berr/trace_test.go:15",
				"  github.com/stretchr/testify/suite.(*Suite).Run.func1:",
				"    /Users/seba/go/pkg/mod/github.com/stretchr/testify@v1.9.0/suite/suite.go:115",
				"  testing.tRunner:",
				"    /nix/store/k9srp8ngvblscg68fdpcyqkydh86429k-go-1.22.1/share/go/src/testing/testing.go:1689",
				"  runtime.goexit:",
				"    /nix/store/k9srp8ngvblscg68fdpcyqkydh86429k-go-1.22.1/share/go/src/runtime/asm_arm64.s:1222",
			}, "\n"),
			true,
			true,
		},
	}
)
