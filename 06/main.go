package main

import (
	"fmt"
	"strings"

	"github.com/j-sv/advent-of-code-2020/input"
)

func countFirst(entry string) int {
	var (
		total int
		set   = make(map[byte]struct{})
	)

	for i := range entry {
		if entry[i] == '\n' {
			continue
		}

		if _, ok := set[entry[i]]; !ok {
			total++
		}

		set[entry[i]] = struct{}{}
	}

	return total
}

func countSecond(entry string) int {
	var (
		total   int
		set     = make(map[byte]int)
		persons = strings.Count(entry, "\n") + 1
	)

	for i := range entry {
		if entry[i] == '\n' {
			continue
		}

		if _, ok := set[entry[i]]; !ok {
			set[entry[i]] = 0
		}

		set[entry[i]]++
	}

	for key := range set {
		if set[key] != persons {
			continue
		}

		total += 1
	}

	return total
}

func parse(contents string) []string {
	entries := strings.Split(contents, "\n\n")

	return entries
}

func main() {
	contents, err := input.Read()
	if err != nil {
		panic(err)
	}

	var (
		totalFirst  int
		totalSecond int
		entries     = parse(strings.TrimSpace(string(contents)))
	)

	for i := range entries {
		totalFirst += countFirst(entries[i])
		totalSecond += countSecond(entries[i])
	}

	fmt.Println(totalFirst)  // 6565
	fmt.Println(totalSecond) // 3137
}
