package aoc

import (
	"testing"
)

func TestDecodeSeatID(t *testing.T) {
	expected := 357
	actual := decodeSeatID("FBFBBFFRLR")
	if actual != expected {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.Fail()
	}
}

func TestFindHighestSeatID(t *testing.T) {
	expected := 955
	r, close := lineReaderFromFile("./inputs/day_5.txt")
	defer close()
	actual := findHighestSeatID(r)
	if actual != expected {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.Fail()
	}
}

func TestFindMissingSeatID(t *testing.T) {
	expected := 569
	r, close := lineReaderFromFile("./inputs/day_5.txt")
	defer close()
	actual := findMissingSeatID(r)
	if actual != expected {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.Fail()
	}
}
