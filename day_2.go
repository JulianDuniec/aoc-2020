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

func countValidPasswords(reader lineReader) int {
	count := 0
	for {
		line, eof := reader.readLine()
		password := parsePassword(line)
		if isPasswordValid(password) {
			count++
		}
		if eof {
			break
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

func countValidPasswordsAlternate(reader lineReader) int {
	count := 0
	for {
		line, eof := reader.readLine()
		password := parsePassword(line)
		if isPasswordValidAlternate(password) {
			count++
		}
		if eof {
			break
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

func parsePassword(line string) password {
	split := strings.Split(line, ":")
	policyPart := split[0]
	passwordPart := split[1]

	policySplit := strings.Split(policyPart, " ")
	lowHighPart := policySplit[0]
	letterPart := policySplit[1]

	lowHighSplit := strings.Split(lowHighPart, "-")
	low, err := strconv.ParseInt(lowHighSplit[0], 10, 32)
	if err != nil {
		log.Fatalf("unable to parse low %v", err)
	}
	high, err := strconv.ParseInt(lowHighSplit[1], 10, 32)
	if err != nil {
		log.Fatalf("unable to parse high %v", err)
	}
	return password{
		policy: passwordPolicy{
			letter: rune(letterPart[0]),
			low:    int(low),
			high:   int(high),
		},
		password: strings.TrimSpace(passwordPart),
	}
}
