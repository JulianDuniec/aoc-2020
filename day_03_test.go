package aoc

import (
	"testing"
)

func TestCountTreeCollissions(t *testing.T) {
	const expected = 7
	r, close := lineReaderFromFile("inputs/day_3_sample.txt")
	defer close()
	pattern, rowLen := parseTreeMap(r)
	actual := countTreeCollissions(pattern, rowLen, 3, 1)
	if expected != actual {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.FailNow()
	}
}

func TestCountTreeCollissionsFinal(t *testing.T) {
	const expected = 272
	r, close := lineReaderFromFile("inputs/day_3.txt")
	defer close()
	pattern, rowLen := parseTreeMap(r)
	actual := countTreeCollissions(pattern, rowLen, 3, 1)
	if expected != actual {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.FailNow()
	}
}

func TestMultiplySlopes(t *testing.T) {
	const expected = 3898725600
	r, close := lineReaderFromFile("inputs/day_3.txt")
	defer close()
	pattern, rowLen := parseTreeMap(r)
	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	product := multiplySlopes(slopes, pattern, rowLen)
	if product != expected {
		t.Logf("product did not match expected")
		t.Fail()
	}
}
