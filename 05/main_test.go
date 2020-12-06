package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		input  string
		row    int
		column int
		seatID int
	}{
		{
			input:  "FBFBBFFRLR",
			row:    44,
			column: 5,
			seatID: 357,
		},
		{
			input:  "BFFFBBFRRR",
			row:    70,
			column: 7,
			seatID: 567,
		},
		{
			input:  "FFFBBBFRRR",
			row:    14,
			column: 7,
			seatID: 119,
		},
		{
			input:  "BBFFBBFRLL",
			row:    102,
			column: 4,
			seatID: 820,
		},
		{
			input:  "BBBBBBBRRR",
			row:    127,
			column: 7,
			seatID: 1023,
		},
		{
			input:  "FFFFFFFLLL",
			row:    0,
			column: 0,
			seatID: 0,
		},
		{
			input:  "FFFFFFFRRR",
			row:    0,
			column: 7,
			seatID: 7,
		},
		{
			input:  "BBBBBBBLLL",
			row:    127,
			column: 0,
			seatID: 1016,
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			s := New(test.input)
			require.NotNil(t, s)

			assert.Equal(t, test.row, s.Row())
			assert.Equal(t, test.column, s.Column())
			assert.Equal(t, test.seatID, s.ID())
		})
	}
}
