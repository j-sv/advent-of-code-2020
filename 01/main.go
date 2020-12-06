package main

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"

	"github.com/j-sv/advent-of-code-2020/input"
)

func findTwoNumbers(entries []int) (int, int) {
	var maxY int

	for x := range entries {
		maxY = 2020 - entries[x]

		for y := range entries[x:] {
			if entries[y] > maxY {
				break
			}

			if entries[x]+entries[y] == 2020 {
				return entries[x], entries[y]
			}
		}
	}

	return -1, -1
}

func findThreeNumbers(entries []int) (int, int, int) {
	var (
		maxY int
		maxZ int
	)

	for x := range entries {
		maxY = 2020 - entries[x]

		for y := range entries[x+1:] {
			if entries[y] > maxY {
				break
			}

			maxZ = maxY - entries[y]

			for z := range entries[x+y+1:] {
				if entries[z] > maxZ {
					break
				}

				if entries[x]+entries[y]+entries[z] == 2020 {
					return entries[x], entries[y], entries[z]
				}
			}
		}
	}

	return -1, -1, -1
}

func findSolutions(entries []int) {
	x, y := findTwoNumbers(entries)

	fmt.Println(x * y) // 927684

	x, y, z := findThreeNumbers(entries)

	fmt.Println(x * y * z) // 292093004
}

func main() {
	file, err := input.Open()
	if err != nil {
		panic(err)
	}
	defer file.Close()

	entries := make([]int, 0, 200)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}

		entries = append(entries, int(num))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	sort.Ints(entries)

	findSolutions(entries)
}
