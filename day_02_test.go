package aoc

import "testing"

func TestCountValidPasswordsFinal(t *testing.T) {
	expected := 542
	r, close := lineReaderFromFile("./inputs/day_2.txt")
	defer close()
	actual := countValidPasswords(r)
	if expected != actual {
		t.Logf("Expected did not match actual %d %d", expected, actual)
		t.FailNow()
	}
}

func TestCountValidPasswordsAlternateFinal(t *testing.T) {
	expected := 360
	r, close := lineReaderFromFile("./inputs/day_2.txt")
	defer close()
	actual := countValidPasswordsAlternate(r)
	if expected != actual {
		t.Logf("Expected did not match actual %d %d", expected, actual)
		t.FailNow()
	}
}
