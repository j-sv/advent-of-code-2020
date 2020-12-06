package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/j-sv/advent-of-code-2020/input"
)

var re = regexp.MustCompile(`^(\d+)-(\d+)\s(\w):\s(\w+)$`)

type Entry struct {
	Min      int
	Max      int
	Char     string
	Password string
}

func NewEntry(buf []byte) *Entry {
	e := new(Entry)
	e.Min, e.Max, e.Char, e.Password = parse(buf)

	return e
}

func mustParseInt(buf []byte) int {
	i, err := strconv.ParseInt(string(buf), 10, 64)
	if err != nil {
		panic(err)
	}

	return int(i)
}

func parse(buf []byte) (int, int, string, string) {
	matches := re.FindSubmatch(buf)

	return mustParseInt(matches[1]), mustParseInt(matches[2]), string(matches[3]), string(matches[4])
}

func count(entries []*Entry, policy func(*Entry) bool) int {
	valid := 0

	for i := range entries {
		if policy(entries[i]) {
			valid++
		}
	}

	return valid
}

func policyOne(e *Entry) bool {
	count := strings.Count(e.Password, e.Char)

	return count >= e.Min && count <= e.Max
}

func policyTwo(e *Entry) bool {
	fst := string(e.Password[e.Min-1])
	snd := string(e.Password[e.Max-1])

	switch {
	case fst == e.Char && snd == e.Char:
		return false
	case fst == e.Char:
		return true
	case snd == e.Char:
		return true
	default:
		return false
	}
}

func main() {
	file, err := input.Open()
	if err != nil {
		panic(err)
	}
	defer file.Close()

	entries := make([]*Entry, 0, 1000)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		entries = append(entries, NewEntry(scanner.Bytes()))
	}

	fmt.Println(count(entries, policyOne)) // 580
	fmt.Println(count(entries, policyTwo)) // 611
}
