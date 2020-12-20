package exercise4

import (
	"adventofcode/common/solution"
	"adventofcode/common/utils"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Exercise4 struct {
	input []Passport
}

type Passport struct {
	contents map[string]string
}

/**
  byr (Birth Year)
  iyr (Issue Year)
  eyr (Expiration Year)
  hgt (Height)
  hcl (Hair Color)
  ecl (Eye Color)
  pid (Passport ID)
  cid (Country ID)
*/
var neededAttributes = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
	//"cid", not needed
}


var hairColorRegex, regexErr = regexp.Compile("^#[0-9a-f]{6}$")

func (p *Passport) isPresent() bool {
	for _, attribute := range neededAttributes {
		if _, found := p.contents[attribute]; !found {
			return false
		}
	}

	return true
}

func isBetween(str string, min int, max int) bool {
	year, err := strconv.Atoi(str)

	if err != nil {
		return false
	}

	return year >= min && year <= max
}

func parseHeight(str string) (string, string) {
	return str[:len(str) - 2], str[len(str) - 2:]
}

func (p *Passport) isValid() bool {
	c := p.contents

	if !isBetween(c["byr"], 1920, 2002) ||
	   !isBetween(c["iyr"], 2010, 2020) ||
	   !isBetween(c["eyr"], 2020, 2030) {
		return false
	}

	height, unit := parseHeight(c["hgt"])
	if unit == "cm" {
		if !isBetween(height, 150, 193) {
			return false
		}
	} else if unit == "in" {
		if !isBetween(height, 59, 76) {
			return false
		}
	} else {
		return false
	}

	if !hairColorRegex.Match([]byte(c["hcl"])) {
		return false
	}

	validEyeColors := map[string]bool {
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}
	if _, found := validEyeColors[c["ecl"]]; !found {
		return false
	}

	_, err := strconv.Atoi(c["pid"])
	if err != nil || len(c["pid"]) != 9 {
		return false
	}

	return true
}

func (p *Passport) String() string {
	ret := strings.Builder{}

	for key, value := range p.contents {
		ret.WriteString(fmt.Sprintf("%s:%s ", key, value))
	}

	return ret.String()
}

func (e *Exercise4) Prepare() error {
	if regexErr != nil {
		return errors.New("could not compile regex")
	}

	input := utils.ReadInput(4)

	contents := make(map[string]string)

	for line := range input {
		if line == "" {
			e.input = append(e.input, Passport{contents: contents})
			contents = make(map[string]string)
		} else {
			parts := strings.Split(line, " ")
			for _, part := range parts {
				keyAndValue := strings.Split(part, ":")
				if len(keyAndValue) != 2 {
					return errors.New("Could not parse input line: " + line)
				}
				contents[keyAndValue[0]] = keyAndValue[1]
			}
		}
	}

	e.input = append(e.input, Passport{contents: contents})

	return nil
}

func (e *Exercise4) Solution1() (solution.Solution, error) {
	validPassports := 0
	for _, passport := range e.input {
		if passport.isPresent() {
			validPassports++
		}
	}
	
	return solution.New(strconv.Itoa(validPassports)), nil
}

func (e *Exercise4) Solution2() (solution.Solution, error) {
	validPassports := 0
	for _, passport := range e.input {
		if passport.isPresent() && passport.isValid() {
			validPassports++
		}
	}

	return solution.New(strconv.Itoa(validPassports)), nil
}
