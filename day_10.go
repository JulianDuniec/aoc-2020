package aoc

import (
	"sort"
)

func multiplyJoltDifferenceCount(reader lineReader) int {
	adapters := readNumbers(reader)
	adapters = append([]int{0}, adapters...)
	sort.Ints(adapters)
	diff3 := 1
	diff1 := 0
	for i := 1; i < len(adapters); i++ {
		diff := adapters[i] - adapters[i-1]
		if diff == 3 {
			diff3++
		} else if diff == 1 {
			diff1++
		}
	}
	return diff3 * diff1
}

func countDistinctArragements(reader lineReader) int {
	adapters := readNumbers(reader)
	adapters = append([]int{0}, adapters...)
	sort.Ints(adapters)

	memo := map[int]int{}
	var findPathCount func(int) int
	findPathCount = func(i int) int {
		if memo[i] != 0 {
			return memo[i]
		}
		if i == len(adapters)-1 {
			return 1
		}
		var sum int
		for ci := i + 1; ci < len(adapters); ci++ {
			if adapters[ci]-adapters[i] > 3 {
				break
			}
			n := findPathCount(ci)
			sum += n
			memo[ci] = n
		}
		return sum
	}
	return findPathCount(0)
}
