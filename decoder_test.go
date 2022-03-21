package dcode

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimpleInt(t *testing.T) {
	r := require.New(t)
	ints := []int{1, 2, 3, 4, 5, 1000}
	decoder := Int()
	for i, expected := range ints {
		expectedBytes := []byte(strconv.Itoa(expected))
		actual, err := DecodeBytes(expectedBytes, decoder)
		r.NoError(err, "for iteration %d, int %d", i, expected)
		r.Equal(expected, actual, "expected int %d", actual)
	}

	notInts := []string{`"abc"`, `"dev"`}
	for _, notInt := range notInts {
		actual, err := DecodeBytes([]byte(notInt), decoder)
		r.True(err != nil)
		r.Equal(0, actual)
	}
}

func TestSimpleString(t *testing.T) {
	r := require.New(t)
	decoder := String()
	strings := []string{
		`"this is a thing"`,
		`"this is another thing"`,
	}
	for i, expected := range strings {
		b := []byte(expected)
		actual, err := DecodeBytes(b, decoder)
		r.NoError(err, "for iteration %d", i)
		actualJSONStr := fmt.Sprintf(`"%s"`, actual)
		r.Equal(expected, actualJSONStr, "for iteration %d", i)
	}
}
