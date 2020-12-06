package main

import (
	"math/rand"
	"os"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	rand.Seed(time.Now().Unix())

	os.Exit(m.Run())
}

func Test_findTwoNumbers(t *testing.T) {
	input := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}

	sort.Ints(input)
	x, y := findTwoNumbers(input)

	assert.Equal(t, 514579, x*y)
}

func Test_findThreeNumbers(t *testing.T) {
	input := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}

	sort.Ints(input)
	x, y, z := findThreeNumbers(input)

	assert.Equal(t, 241861950, x*y*z)
}

func Benchmark_findTwoNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		entries := rand.Perm(2020)
		sort.Ints(entries)
		findTwoNumbers(entries)
	}
}

func Benchmark_findThreeNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		entries := rand.Perm(2020)
		sort.Ints(entries)
		findThreeNumbers(entries)
	}
}
