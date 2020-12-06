package aoc

import (
	"strings"
)

func SumDistinctPositiveAnswers(file string) int {
	sum := 0
	distinct := make(map[rune]struct{}, 0)
	iterateLines(file, func(line string) {
		line = strings.TrimSpace(line)
		if line == "" {
			sum += len(distinct)
			distinct = make(map[rune]struct{}, 0)
			return
		}
		for _, char := range line {
			distinct[char] = struct{}{}
		}
	})
	sum += len(distinct)
	return sum
}

func SumIntersectingPositiveAnswers(file string) int {
	sum := 0
	var intersecting []rune
	iterateLines(file, func(line string) {
		line = strings.TrimSpace(line)
		if line == "" {
			sum += len(intersecting)
			intersecting = nil
			return
		}
		if intersecting == nil {
			intersecting = []rune(line)
			return
		}
	outer:
		for i := len(intersecting) - 1; i >= 0; i-- {
			c1 := intersecting[i]
			for _, c2 := range line {
				if c1 == c2 {
					continue outer
				}
			}
			intersecting = append(intersecting[:i], intersecting[i+1:]...)
		}
	})
	sum += len(intersecting)
	return sum
}
