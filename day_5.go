package aoc

import (
	"sort"
	"strings"
)

func decodeSeatID(encoded string) int {
	rows := encoded[0:7]
	low, high := 0, 127
	for _, c := range rows {
		dp := (high - low) / 2
		if c == 'F' {
			high = low + int(dp)
		} else if c == 'B' {
			low = high - int(dp)
		}
	}
	row := low

	columns := encoded[7:]
	low, high = 0, 7
	for _, c := range columns {
		dp := (high - low) / 2
		if c == 'L' {
			high = low + int(dp)
		} else if c == 'R' {
			low = high - int(dp)
		}
	}
	column := low
	return row*8 + column
}

func findHighestSeatID(file string) int {
	max := 0
	iterateLines(file, func(line string) {
		line = strings.TrimSpace(line)
		seatID := decodeSeatID(line)
		if seatID > max {
			max = seatID
		}
	})
	return max
}

func findMissingSeatID(file string) int {
	seatIDs := make([]int, 0)
	iterateLines(file, func(line string) {
		line = strings.TrimSpace(line)
		seatID := decodeSeatID(line)
		seatIDs = append(seatIDs, seatID)
	})
	sort.Slice(seatIDs, func(i, j int) bool {
		return seatIDs[i] < seatIDs[j]
	})
	prev := -1
	for _, seatID := range seatIDs {
		if prev == -1 {
			prev = seatID
			continue
		}
		diff := seatID - prev
		if diff == 2 {
			return seatID - 1
		}
		prev = seatID
	}
	panic("no solution")
}
