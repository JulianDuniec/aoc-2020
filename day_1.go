package aoc

import (
	"log"
	"sort"
	"strconv"
)

func findSumsAndMultiply(reader lineReader) int {
	set := make(map[int]struct{}, 0)
	for {
		line, eof := reader.readLine()
		v, err := strconv.ParseInt(line, 10, 32)
		if err != nil {
			log.Fatalf("invalid input %s %v", line, err)
		}
		set[int(v)] = struct{}{}
		if eof {
			break
		}
	}
	for v := range set {
		diff := 2020 - v
		if _, ok := set[diff]; ok {
			return v * diff
		}
	}
	panic("no solution")
}

func findThreeSumsAndMultiply(reader lineReader) int {
	entries := []int{}
	for {
		line, eof := reader.readLine()
		v, err := strconv.ParseInt(line, 10, 32)
		if err != nil {
			log.Fatalf("invalid input %s %v", line, err)
		}
		entries = append(entries, int(v))
		if eof {
			break
		}
	}
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
