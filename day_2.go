package aoc

import (
	"log"
	"strconv"
	"strings"
)

type passwordPolicy struct {
	letter    rune
	low, high int
}

type password struct {
	policy   passwordPolicy
	password string
}

// CountValidPasswords https://adventofcode.com/2020/day/2
func CountValidPasswords(passwords []password) int {
	count := 0
	for _, pwd := range passwords {
		if isPasswordValid(pwd) {
			count++
		}
	}
	return count
}

func isPasswordValid(password password) bool {
	runeCount := 0
	for _, r := range password.password {
		if r == password.policy.letter {
			runeCount++
		}
	}
	return runeCount >= password.policy.low && runeCount <= password.policy.high
}

// CountValidPasswordsAlternate https://adventofcode.com/2020/day/2#part2
func CountValidPasswordsAlternate(passwords []password) int {
	count := 0
	for _, pwd := range passwords {
		if isPasswordValidAlternate(pwd) {
			count++
		}
	}
	return count
}

func isPasswordValidAlternate(password password) bool {
	runeIndex := 1
	var low rune
	var high rune
	for _, r := range password.password {
		if runeIndex == password.policy.low {
			low = r
		}
		if runeIndex == password.policy.high {
			high = r
			break
		}
		runeIndex++
	}
	l := password.policy.letter
	return (l == low) != (l == high)
}

func parseInputAsPasswords(file string) []password {
	res := make([]password, 0)
	iterateLines(file, func(line string) {
		line = strings.TrimSpace(line)
		split := strings.Split(line, ":")
		policyPart := split[0]
		passwordPart := split[1]

		policySplit := strings.Split(policyPart, " ")
		lowHighPart := policySplit[0]
		letterPart := policySplit[1]

		lowHighSplit := strings.Split(lowHighPart, "-")
		low, err := strconv.ParseInt(strings.TrimSpace(lowHighSplit[0]), 10, 32)
		if err != nil {
			log.Fatalf("unable to parse low %v", err)
		}
		high, err := strconv.ParseInt(strings.TrimSpace(lowHighSplit[1]), 10, 32)
		if err != nil {
			log.Fatalf("unable to parse high %v", err)
		}
		res = append(res, password{
			policy: passwordPolicy{
				letter: rune(letterPart[0]),
				low:    int(low),
				high:   int(high),
			},
			password: strings.TrimSpace(passwordPart),
		})
	})
	return res
}
