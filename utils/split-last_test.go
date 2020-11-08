package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitLast(t *testing.T) {
	t.Run("must return valid results", func(t *testing.T) {
		cases := []struct {
			param1    string
			delimiter uint8
			res       string
		}{
			{
				param1:    "name3-name2-name1",
				delimiter: '-',
				res:       "name1",
			},
			{
				param1:    "what?",
				delimiter: '|',
				res:       "what?",
			},
			{
				param1:    "do it! just do it!",
				delimiter: '!',
				res:       "",
			},
		}

		for _, c := range cases {
			res := SplitLast(c.param1, c.delimiter)
			assert.Equal(t, c.res, res, "must equals")
		}
	})
}

func TestSplitLastByte(t *testing.T) {
	t.Run("must return valid results", func(t *testing.T) {
		cases := []struct {
			param1    []byte
			delimiter uint8
			res       []byte
			resCap    int
		}{
			{
				param1:    []byte{'a', 'v', '-', '1', '3'},
				delimiter: '-',
				res:       []byte{'1', '3'},
				resCap:    2,
			},
			{
				param1:    []byte{'1', '2'},
				delimiter: '|',
				res:       []byte{'1', '2'},
				resCap:    2,
			},
			{
				param1:    []byte("do it! just do it!"),
				delimiter: '!',
				res:       []byte{},
				resCap:    0,
			},
		}

		for _, c := range cases {
			res := SplitLastByte(c.param1, c.delimiter)
			assert.Equal(t, c.res, res, "must equals")
			assert.Equal(t, c.resCap, cap(res), "capacity must equals")
		}
	})
}
