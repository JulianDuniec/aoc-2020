package aoc

import "testing"

func TestValidatePassports(t *testing.T) {
	const expected = 2
	r, close := lineReaderFromFile("./inputs/day_4_sample.txt")
	defer close()
	actual := validatePassports(r, false)
	if actual != expected {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.FailNow()
	}
}

func TestValidatePassportsFinal(t *testing.T) {
	const expected = 233
	r, close := lineReaderFromFile("./inputs/day_4.txt")
	defer close()
	actual := validatePassports(r, false)
	if actual != expected {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.FailNow()
	}
}

func TestValidatePassportsPartTwo(t *testing.T) {
	const expected = 4 // 4 invalid, 4 valid
	r, close := lineReaderFromFile("./inputs/day_4_pt2_sample.txt")
	defer close()
	actual := validatePassports(r, true)
	if actual != expected {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.FailNow()
	}
}

func TestValidatePassportsPartTwoFinal(t *testing.T) {
	const expected = 111
	r, close := lineReaderFromFile("./inputs/day_4.txt")
	defer close()
	actual := validatePassports(r, true)
	if actual != expected {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.FailNow()
	}
}
