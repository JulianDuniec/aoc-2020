package aoc

import "testing"

func TestSumDistinctPositiveAnswers(t *testing.T) {
	expected := 11
	r, close := lineReaderFromFile("./inputs/day_6_sample.txt")
	defer close()
	actual := sumDistinctPositiveAnswers(r)
	if expected != actual {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.Fail()
	}
}

func TestSumDistinctPositiveAnswersFinal(t *testing.T) {
	expected := 6504
	r, close := lineReaderFromFile("./inputs/day_6.txt")
	defer close()
	actual := sumDistinctPositiveAnswers(r)
	if expected != actual {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.Fail()
	}
}

func TestSumIntersectingPositiveAnswers(t *testing.T) {
	expected := 6
	r, close := lineReaderFromFile("./inputs/day_6_sample.txt")
	defer close()
	actual := sumIntersectingPositiveAnswers(r)
	if expected != actual {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.Fail()
	}
}

func TestSumIntersectingPositiveAnswersFinal(t *testing.T) {
	expected := 3351
	r, close := lineReaderFromFile("./inputs/day_6.txt")
	defer close()
	actual := sumIntersectingPositiveAnswers(r)
	if expected != actual {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.Fail()
	}
}
