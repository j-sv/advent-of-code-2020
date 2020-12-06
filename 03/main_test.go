package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_traversal(t *testing.T) {
	lines := []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}

	tests := []struct {
		t        *toboggan
		expected int
	}{
		{
			t: &toboggan{
				right: 1,
				down:  1,
			},
			expected: 2,
		},
		{
			t: &toboggan{
				right: 3,
				down:  1,
			},
			expected: 7,
		},
		{
			t: &toboggan{
				right: 5,
				down:  1,
			},
			expected: 3,
		},
		{
			t: &toboggan{
				right: 7,
				down:  1,
			},
			expected: 4,
		},
		{
			t: &toboggan{
				right: 1,
				down:  2,
			},
			expected: 2,
		},
	}

	total := 1

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			actual := test.t.Traverse(lines)
			require.Equal(t, test.expected, actual)

			total *= actual
		})
	}

	assert.Equal(t, 336, total)
}
