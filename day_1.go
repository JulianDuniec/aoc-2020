package aoc

import (
	"sort"
)

// FindSumsAndMultiply https://adventofcode.com/2020/day/1
func FindSumsAndMultiply(entries []int) int {
	len := len(entries)
	set := make(map[int]struct{}, len)
	for _, v := range entries {
		set[v] = struct{}{}
	}
	for _, v := range entries {
		diff := 2020 - v
		if _, ok := set[diff]; ok {
			return v * diff
		}
	}
	panic("no solution")
}

// FindThreeSumsAndMultiply https://adventofcode.com/2020/day/1#part2
func FindThreeSumsAndMultiply(entries []int) int {
	sort.Slice(entries, func(i, j int) bool {
		return entries[i] < entries[j]
	})
	len := len(entries)
	for i := 0; i < len-2; i++ {
		j := i + 1
		k := len - 1
		for j < k {
			a, b, c := entries[i], entries[j], entries[k]
			sum := a + b + c
			if sum == 2020 {
				return a * b * c
			} else if sum < 2020 {
				j++
			} else if sum > 2020 {
				k--
			}
		}
	}
	panic("no solution")
}
