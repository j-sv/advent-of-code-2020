package main

import (
	"bufio"
	"fmt"

	"github.com/j-sv/advent-of-code-2020/input"
)

type toboggan struct {
	right int
	down  int
}

func (t *toboggan) Traverse(lines []string) int {
	var (
		height = len(lines)
		width  = len(lines[0])
		hits   int
		x      int
		y      int
	)

	for ; x < height; x += t.down {
		if lines[x][y] == byte('#') {
			hits++
		}

		y = (y + t.right) % width
	}

	return hits
}

func main() {
	file, err := input.Open()
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := make([]string, 0, 323)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println((&toboggan{right: 3, down: 1}).Traverse(lines)) // 299

	toboggans := []*toboggan{
		{
			right: 1,
			down:  1,
		},
		{
			right: 3,
			down:  1,
		},
		{
			right: 5,
			down:  1,
		},
		{
			right: 7,
			down:  1,
		},
		{
			right: 1,
			down:  2,
		},
	}
	total := 1

	for i := range toboggans {
		total *= toboggans[i].Traverse(lines)
	}

	fmt.Println(total) // 3621285278
}
