package aoc

import (
	"strings"
)

// CountTreeCollissions https://adventofcode.com/2020/day/3
func CountTreeCollissions(pattern []int32, rowLen int, dx, dy int) int {
	ix, iy := 0, 0
	len := len(pattern)
	sum := 0
	for iy < len {
		l := pattern[iy]
		mod := ix % rowLen
		if l&(1<<mod) > 0 {
			sum++
		}
		ix += dx
		iy += dy
	}
	return sum
}

// MultiplySlopes https://adventofcode.com/2020/day/3#part2
func MultiplySlopes(slopes [][]int, pattern []int32, rowLen int) int {
	product := 1
	for _, v := range slopes {
		product *= CountTreeCollissions(pattern, rowLen, v[0], v[1])
	}
	return product
}

func parseTreeMap(file string) ([]int32, int) {
	res := make([]int32, 0)
	rowLen := 0
	iterateLines(file, func(line string) {
		line = strings.TrimSpace(line)
		l := 0
		rowLen = len(line)
		for i, char := range line {
			if char == '#' {
				l ^= (1 << i)
			}
		}
		res = append(res, int32(l))
	})
	return res, rowLen
}
