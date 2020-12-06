package aoc

import (
	"testing"
)

func TestSumAndMultiply(t *testing.T) {
	expected := 514579
	r := lineReaderFromString(`1721
979
366
299
675
1456`)
	actual := findSumsAndMultiply(r)
	if actual != expected {
		t.Logf("Expected did not match actual %d %d", expected, actual)
		t.FailNow()
	}
}

func TestSumAndMultiplyFinal(t *testing.T) {
	expected := 259716
	r, close := lineReaderFromFile("./inputs/day_1.txt")
	defer close()
	actual := findSumsAndMultiply(r)
	if expected != actual {
		t.Logf("Expected did not match actual %d %d", expected, actual)
		t.FailNow()
	}
}

func TestThreeSumsAndMultiply(t *testing.T) {
	expected := 241861950
	r := lineReaderFromString(`1721
979
366
299
675
1456`)
	actual := findThreeSumsAndMultiply(r)
	if actual != expected {
		t.Logf("Expected did not match actual %d %d", expected, actual)
		t.FailNow()
	}
}

func TestThreeSumsAndMultiplyFinal(t *testing.T) {
	expected := 120637440
	r, close := lineReaderFromFile("./inputs/day_1.txt")
	defer close()
	actual := findThreeSumsAndMultiply(r)
	if expected != actual {
		t.Logf("Expected did not match actual %d %d", expected, actual)
	}
}
