package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterByte(t *testing.T) {
	t.Run("must return valid results", func(t *testing.T) {
		cases := []struct {
			param1 []byte
			filter uint8
			res    []byte
			resCap int
		}{
			{
				param1: []byte{'a', 'v', '-', '1', '3'},
				filter: '-',
				res:    []byte{'a', 'v', '1', '3'},
				resCap: 4,
			},
			{
				param1: []byte{'1', '2'},
				filter: '|',
				res:    []byte{'1', '2'},
				resCap: 2,
			},
			{
				param1: []byte("do it! just do it!"),
				filter: '!',
				res:    []byte("do it just do it"),
				resCap: 16,
			},
		}

		for _, c := range cases {
			res := FilterByte(c.param1, c.filter)
			assert.Equal(t, c.res, res, "must equals")
			assert.Equal(t, c.resCap, cap(res), "capacity must equals")
		}
	})
}
