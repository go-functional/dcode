package dcode

import (
	"testing"

	tree "github.com/bmatsuo/go-jsontree"
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
	},
	"simple": 1
}`

type TheSuite struct {
	suite.Suite
}

func (t TheSuite) getTree(b []byte) (*tree.JsonTree, error) {
	tr := tree.New()
	if err := tr.UnmarshalJSON(b); err != nil {
		return nil, err
	}
	return tr, nil
}

func TestTheSuite(t *testing.T) {
	suite.Run(t, new(TheSuite))
}
