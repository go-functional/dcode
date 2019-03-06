package dcode

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

const jsonStr = `{
	"a":{
		"b": {
			"c": {
				"d": 200
			}
		},
		"otherb": {
			"otherc": 200
		}
	}
}`

type TheSuite struct {
	suite.Suite
}

func TestTheSuite(t *testing.T) {
	suite.Run(t, new(TheSuite))
}
