package aoc

import (
	"math"
)

func findFirstNumberWithoutSumInPreamble(numbers []int, preambleSize int) int {
	set := make(map[int]int, 0)
	for i, num := range numbers {
		if len(set) == preambleSize {
			valid := false
			for v := range set {
				diff := num - v
				if _, ok := set[diff]; ok {
					valid = true
				}
			}
			if !valid {
				return num
			}
		}
		set[num] = i
		if len(set) < preambleSize {
			continue
		} else if len(set) > preambleSize {
			t, tk := intptr(math.MaxInt32), intptr(-1)
			for k, v := range set {
				if v < *t {
					*t = v
					*tk = k
				}
			}
			delete(set, *tk)
		}
	}
	panic("no solution")
}

func intptr(i int) *int {
	return &i
}

func findEnchryptionWeakness(numbers []int, preambleSize int) int {
	invalid := findFirstNumberWithoutSumInPreamble(numbers, preambleSize)
	for i, v := range numbers {
		sum := v
		for j, v2 := range numbers[i+1:] {
			sum += v2
			if sum == invalid {
				return addMinMax(numbers[i : i+j+2])
			} else if sum > invalid {
				break
			}
		}
	}
	panic("no solution")
}

func addMinMax(numbers []int) int {
	min := math.MaxInt64
	max := 0
	for _, v := range numbers {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min + max
}
