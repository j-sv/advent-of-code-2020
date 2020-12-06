package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testInput = `abc

a
b
c

ab
ac

a
a
a
a

b`

func Test_countFirst_input(t *testing.T) {
	entries := parse(testInput)

	require.Len(t, entries, 5)

	total := 0

	for i := range entries {
		c := countFirst(entries[i])
		total += c
	}

	assert.Equal(t, 11, total)
}

func Test_countFirst_entries(t *testing.T) {
	tests := []struct {
		entry    string
		expected int
	}{
		{
			entry:    "abc",
			expected: 3,
		},
		{
			entry:    "a\nb\nc",
			expected: 3,
		},
		{
			entry:    "ab\nac",
			expected: 3,
		},
		{
			entry:    "a\na\na\na",
			expected: 1,
		},
		{
			entry:    "b",
			expected: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.entry, func(t *testing.T) {
			assert.Equal(t, test.expected, countFirst(test.entry))
		})
	}
}

func Test_countSecond_input(t *testing.T) {
	entries := parse(testInput)

	require.Len(t, entries, 5)

	total := 0

	for i := range entries {
		c := countSecond(entries[i])
		total += c
	}

	assert.Equal(t, 6, total)
}

func Test_countSecond_entries(t *testing.T) {
	tests := []struct {
		entry    string
		expected int
	}{
		{
			entry:    "abc",
			expected: 3,
		},
		{
			entry:    "a\nb\nc",
			expected: 0,
		},
		{
			entry:    "ab\nac",
			expected: 1,
		},
		{
			entry:    "a\na\na\na",
			expected: 1,
		},
		{
			entry:    "b",
			expected: 1,
		},
		{
			entry: `ju
uj
uj
ju`,
			expected: 2,
		},
		{
			entry: `pyg
gpy`,
			expected: 3,
		},
		{
			entry: `e
			xwv
e`,
			expected: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.entry, func(t *testing.T) {
			assert.Equal(t, test.expected, countSecond(test.entry))
		})
	}
}
