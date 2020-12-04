package aoc

import "testing"

func TestValidatePassports(t *testing.T) {
	const expected = 2
	passports := parsePassports("inputs/day_4_sample.txt")
	actual := ValidatePassports(passports)
	if actual != expected {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.FailNow()
	}
}

func TestValidatePassportsFinal(t *testing.T) {
	const expected = 233
	passports := parsePassports("inputs/day_4.txt")
	actual := ValidatePassports(passports)
	if actual != expected {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.FailNow()
	}
}

func TestValidatePassportsPartTwo(t *testing.T) {
	const expected = 4 // 4 invalid, 4 valid
	passports := parsePassports("inputs/day_4_pt2_sample.txt")
	actual := ValidatePassportsStrict(passports)
	if actual != expected {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.FailNow()
	}
}

func TestValidatePassportsPartTwoFinal(t *testing.T) {
	const expected = 111
	passports := parsePassports("inputs/day_4.txt")
	actual := ValidatePassportsStrict(passports)
	if actual != expected {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.FailNow()
	}
}
