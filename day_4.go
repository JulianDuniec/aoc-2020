package aoc

import (
	"regexp"
	"strings"
)

func validatePassports(reader lineReader, strict bool) int {
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
	var current map[string]string
	for {
		line, eof := reader.readLine()
		if current == nil {
			current = make(map[string]string, 0)
		}
		kvpairs := strings.Split(line, " ")
		for _, kvpair := range kvpairs {
			split := strings.Split(kvpair, ":")
			if len(split) == 1 {
				continue
			}
			current[split[0]] = split[1]
		}
		if eof || line == "" {
			valid := true
			for k, r := range requiredKeys {
				v, ok := current[k]
				if !ok {
					valid = false
					break
				}
				if r.Match([]byte(v)) == false && strict {
					valid = false
					break
				}
			}
			if valid {
				numValid++
			}
			current = nil
		}
		if eof {
			break
		}
	}
	return numValid
}
