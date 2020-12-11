package aoc

import (
	"reflect"
)

func findAvailableSeatsAtEquilibrium(reader lineReader, adj adjacency, threshold int) int {
	m := parseSeatLayout(reader)
	for {
		mp := advanceSequence(m, adj, threshold)
		if !reflect.DeepEqual(m, mp) {
			m = mp
			continue
		}
		return sumAvailableSeats(mp)
	}
}

func parseSeatLayout(reader lineReader) [][]rune {
	matrix := [][]rune{}
	for i := 0; ; i++ {
		line, eof := reader.readLine()
		l := []rune{}

		for _, c := range line {
			l = append(l, c)
		}
		matrix = append(matrix, l)
		if eof {
			break
		}
	}
	return matrix
}

type adjacency func(int, int, [][]rune) int

func advanceSequence(matrix [][]rune, adj adjacency, threshold int) [][]rune {
	res := [][]rune{}
	for i, row := range matrix {
		resline := []rune{}
		for j, v := range row {
			if v == '.' {
				resline = append(resline, v)
				continue
			}
			na := adj(i, j, matrix)
			if v == 'L' {
				if na == 0 {
					resline = append(resline, '#')
					continue
				}
				resline = append(resline, v)
			}
			if v == '#' {
				if na >= threshold {
					resline = append(resline, 'L')
					continue
				}
				resline = append(resline, v)
			}
		}
		res = append(res, resline)
	}
	return res
}

func numOccupiedAdjacent(i, j int, m [][]rune) int {
	sum := 0
	si := maxInt(i-1, 0)
	sit := minInt(i+1, len(m)-1)
	for ; si <= sit; si++ {
		r := m[si]
		sj := maxInt(j-1, 0)
		sjt := minInt(j+1, len(r)-1)
		for ; sj <= sjt; sj++ {
			if si == i && sj == j {
				continue
			}
			if m[si][sj] == '#' {
				sum++
			}
		}
	}
	return sum
}

func numOccupiedVisible(i, j int, m [][]rune) int {
	type dir struct{ i, j int }
	dirs := []dir{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	sum := 0
	mlen := len(m)
	rlen := len(m[0])
	for _, dir := range dirs {
		di, dj := i, j
		for {
			di, dj = dir.i+di, dir.j+dj
			if di < 0 || di > mlen-1 {
				break
			}
			if dj < 0 || dj > rlen-1 {
				break
			}
			v := m[di][dj]
			if v == '#' {
				sum++
				break
			}
			if v == 'L' {
				break
			}
		}
	}
	return sum
}

func sumAvailableSeats(r [][]rune) int {
	sum := 0
	for _, l := range r {
		for _, c := range l {
			if c == '#' {
				sum++
			}
		}
	}
	return sum
}
