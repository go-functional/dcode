package dcode

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TheSuite struct {
	suite.Suite
}

func TestTheSuite(t *testing.T) {
	suite.Run(t, new(TheSuite))
}
