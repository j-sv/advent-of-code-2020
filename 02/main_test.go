package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countValidPolicyOne(t *testing.T) {
	input := [][]byte{
		[]byte("1-3 a: abcde"),
		[]byte("1-3 b: cdefg"),
		[]byte("2-9 c: ccccccccc"),
	}

	entries := make([]*Entry, 3)

	for i := range entries {
		entries[i] = NewEntry(input[i])
	}

	assert.Equal(t, 2, count(entries, policyOne))
}

func Test_countValidPolicyTwo(t *testing.T) {
	input := [][]byte{
		[]byte("1-3 a: abcde"),
		[]byte("1-3 b: cdefg"),
		[]byte("2-9 c: ccccccccc"),
	}

	entries := make([]*Entry, 3)

	for i := range entries {
		entries[i] = NewEntry(input[i])
	}

	assert.Equal(t, 1, count(entries, policyTwo))
}
