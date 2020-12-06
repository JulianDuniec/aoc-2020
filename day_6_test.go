package aoc

import "testing"

func TestSumDistinctPositiveAnswers(t *testing.T) {
	expected := 11
	actual := SumDistinctPositiveAnswers("./inputs/day_6_sample.txt")
	if expected != actual {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.Fail()
	}
}

func TestSumDistinctPositiveAnswersFinal(t *testing.T) {
	expected := 6504
	actual := SumDistinctPositiveAnswers("./inputs/day_6.txt")
	if expected != actual {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.Fail()
	}
}

func TestSumIntersectingPositiveAnswers(t *testing.T) {
	expected := 6
	actual := SumIntersectingPositiveAnswers("./inputs/day_6_sample.txt")
	if expected != actual {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.Fail()
	}
}

func TestSumIntersectingPositiveAnswersFinal(t *testing.T) {
	expected := 3351
	actual := SumIntersectingPositiveAnswers("./inputs/day_6.txt")
	if expected != actual {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.Fail()
	}
}
