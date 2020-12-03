package aoc

import (
	"testing"
)

func TestCountTreeCollissions2(t *testing.T) {
	const expected = 7
	pattern, rowLen := parseTreeMap("inputs/day_3_sample.txt")
	actual := CountTreeCollissions(pattern, rowLen, 3, 1)
	if expected != actual {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.FailNow()
	}
}
func TestCountTreeCollissionsFinal(t *testing.T) {
	const expected = 272
	pattern, rowLen := parseTreeMap("inputs/day_3.txt")
	actual := CountTreeCollissions(pattern, rowLen, 3, 1)
	if expected != actual {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.FailNow()
	}
}

func TestCountTreeCollissionsFinal2(t *testing.T) {
	const expected = 3898725600
	pattern, rowLen := parseTreeMap("inputs/day_3.txt")
	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	product := MultiplySlopes(slopes, pattern, rowLen)
	if product != expected {
		t.Logf("product did not match expected")
		t.Fail()
	}
}
