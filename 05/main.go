package main

import (
	"bufio"
	"fmt"
	"sort"

	"github.com/j-sv/advent-of-code-2020/input"
)

const (
	maxRow    = 127
	maxColumn = 7
)

type seat struct {
	row    string
	column string
}

func New(input string) *seat {
	return &seat{
		row:    input[0:7],
		column: input[7:10],
	}
}

func (s *seat) Row() int {
	return search(s.row, 'F', 'B', maxRow)
}

func (s *seat) Column() int {
	return search(s.column, 'L', 'R', maxColumn)
}

func (s *seat) ID() int {
	return s.Row()*8 + s.Column()
}

func search(input string, keepLow, keepHigh byte, max int) int {
	var (
		min int
		mid int
	)

	for i := range input {
		mid = (max-min)/2 + 1

		switch input[i] {
		case keepLow:
			max -= mid
		case keepHigh:
			min += mid
		}
	}

	return min
}

func main() {
	file, err := input.Open()
	if err != nil {
		panic(err)
	}
	defer file.Close()

	seatIDs := make([]int, 0, 927)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		seat := New(scanner.Text())
		seatIDs = append(seatIDs, seat.ID())
	}

	sort.Ints(seatIDs)

	fmt.Println(seatIDs[len(seatIDs)-1]) // 933

	for i := 1; i < len(seatIDs); i++ {
		if seatIDs[i]-seatIDs[i-1] > 1 {
			fmt.Println(seatIDs[i] - 1) // 711
			break
		}
	}
}
