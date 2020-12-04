package aoc

import (
	"regexp"
	"strings"
)

func ValidatePassports(passports []map[string]string) int {
	requiredKeys := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	numValid := 0
outer:
	for _, passport := range passports {
		for _, k := range requiredKeys {
			if _, ok := passport[k]; !ok {
				continue outer
			}
		}
		numValid++
	}
	return numValid
}

func parsePassports(file string) []map[string]string {
	passports := make([]map[string]string, 0)
	var current map[string]string
	iterateLines(file, func(line string) {
		line = strings.TrimSpace(line)
		if current == nil {
			current = make(map[string]string, 0)
		}
		if line == "" {
			passports = append(passports, current)
			current = nil
			return
		}
		kvpairs := strings.Split(line, " ")
		for _, kvpair := range kvpairs {
			split := strings.Split(kvpair, ":")
			current[split[0]] = split[1]
		}
	})
	if current != nil {
		passports = append(passports, current)
	}
	return passports
}

func ValidatePassportsStrict(passports []map[string]string) int {
	requiredKeys := map[string]*regexp.Regexp{
		"byr": regexp.MustCompile("^(19[2-8][0-9]|199[0-9]|200[0-2])$"),
		"iyr": regexp.MustCompile("^(201[0-9]|2020)$"),
		"eyr": regexp.MustCompile("^(202[0-9]|2030)$"),
		"hgt": regexp.MustCompile("^((1[5-8][0-9]|19[0-3])cm)|((59|6[0-9]|7[0-6])in)$"),
		"hcl": regexp.MustCompile("^\\#[a-f0-9]{6}$"),
		"ecl": regexp.MustCompile("^amb|blu|brn|gry|grn|hzl|oth$"),
		"pid": regexp.MustCompile("^\\d{9}$"),
	}
	numValid := 0
outer:
	for _, passport := range passports {
		for k, r := range requiredKeys {
			v, ok := passport[k]
			if !ok {
				continue outer
			}
			if r.Match([]byte(v)) == false {
				continue outer
			}
		}
		numValid++
	}
	return numValid
}
