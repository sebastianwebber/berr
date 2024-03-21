package berr

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type betterErrorsTestSuite struct {
	suite.Suite
}

func TestBetterErrors(t *testing.T) {
	suite.Run(t, new(betterErrorsTestSuite))
}
