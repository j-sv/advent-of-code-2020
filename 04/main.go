package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/j-sv/advent-of-code-2020/input"
)

var (
	reSpace      = regexp.MustCompile(`\s+`)
	reHeight     = regexp.MustCompile(`^(\d{2,3})(cm|in)$`)
	reHairColor  = regexp.MustCompile(`^#[a-f0-9]{6}$`)
	rePassportID = regexp.MustCompile(`^[0-9]{9}$`)
)

type passport map[string]string

func (p passport) Validate(strict bool) bool {
	if !p.checkMandatoryKeys() {
		return false
	}

	if !strict {
		return true
	}

	if !p.validIntegerRange("byr", 1920, 2002) {
		return false
	}

	if !p.validIntegerRange("iyr", 2010, 2020) {
		return false
	}

	if !p.validIntegerRange("eyr", 2020, 2030) {
		return false
	}

	if !p.validHeight() {
		return false
	}

	if !p.validHairColor() {
		return false
	}

	if !p.validEyeColor() {
		return false
	}

	return p.validPassportID()
}

func (p passport) checkMandatoryKeys() bool {
	keys := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for i := range keys {
		if _, ok := p[keys[i]]; !ok {
			return false
		}
	}

	return true
}

func (p passport) validIntegerRange(key string, min, max int) bool {
	value := mustParseInt(p[key])

	return value >= min && value <= max
}

func (p passport) validHeight() bool {
	matches := reHeight.FindAllStringSubmatch(p["hgt"], -1)
	if matches == nil {
		return false
	}

	height, unit := mustParseInt(matches[0][1]), matches[0][2]

	switch unit {
	case "cm":
		return height >= 150 && height <= 193
	case "in":
		return height >= 59 && height <= 76
	default:
		return false
	}
}

func (p passport) validHairColor() bool {
	return reHairColor.MatchString(p["hcl"])
}

func (p passport) validEyeColor() bool {
	switch p["ecl"] {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	default:
		return false
	}
}

func (p passport) validPassportID() bool {
	return rePassportID.MatchString(p["pid"])
}

func parse(contents string) []passport {
	entries := strings.Split(contents, "\n\n")

	passports := make([]passport, len(entries))

	for i := range entries {
		passports[i] = make(map[string]string)

		fields := reSpace.Split(entries[i], -1)

		for j := range fields {
			if fields[j] == "" {
				continue
			}

			values := strings.Split(fields[j], ":")
			passports[i][values[0]] = values[1]
		}
	}

	return passports
}

func mustParseInt(value string) int {
	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(err)
	}

	return int(i)
}

func main() {
	contents, err := input.Read()
	if err != nil {
		panic(err)
	}

	passports := parse(string(contents))

	fst := 0
	snd := 0

	for i := range passports {
		if passports[i].Validate(false) {
			fst++
		}

		if passports[i].Validate(true) {
			snd++
		}
	}

	fmt.Println(fst) // 182
	fmt.Println(snd) // 109
}
