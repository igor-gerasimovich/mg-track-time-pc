package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytesCoveredRune(t *testing.T) {
	t.Run("must return valid results", func(t *testing.T) {
		cases := []struct {
			param1 []byte
			filter uint8
			res    [][]byte
		}{
			{
				param1: []byte(`my name "is", "igor"`),
				filter: '"',
				res:    [][]byte{[]byte("is"), []byte("igor")},
			},
			{
				param1: []byte(`my name is, igor`),
				filter: '"',
				res:    [][]byte{},
			},
			{
				param1: []byte(`my name "is", igor "`),
				filter: '"',
				res:    [][]byte{[]byte("is")},
			},
		}

		for _, c := range cases {
			res := BytesCoveredRune(c.param1, c.filter)
			assert.Equal(t, c.res, res, "must equals")
		}
	})
}
