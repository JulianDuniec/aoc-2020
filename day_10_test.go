package aoc

import "testing"

const day10Sample = `16
10
15
5
1
11
7
19
6
12
4`

const day10SampleAlt = `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`

func TestCalculateJolt(t *testing.T) {
	r := lineReaderFromString(day10Sample)
	expected := 35
	actual := multiplyJoltDifferenceCount(r)
	if actual != expected {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.Fail()
	}
}

func TestCalculateJoltAlt(t *testing.T) {
	r := lineReaderFromString(day10SampleAlt)
	expected := 220
	actual := multiplyJoltDifferenceCount(r)
	if actual != expected {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.Fail()
	}
}
func TestCalculateJoltFinal(t *testing.T) {
	r, close := lineReaderFromFile("./inputs/day_10.txt")
	defer close()
	expected := 2277
	actual := multiplyJoltDifferenceCount(r)
	if actual != expected {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.Fail()
	}
}

func TestCountDistinctArrangements(t *testing.T) {
	r := lineReaderFromString(day10Sample)
	expected := 8
	actual := countDistinctArragements(r)
	if actual != expected {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.Fail()
	}
}

func TestCountDistinctArrangementsAlt(t *testing.T) {
	r := lineReaderFromString(day10SampleAlt)
	expected := 19208
	actual := countDistinctArragements(r)
	if actual != expected {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.Fail()
	}
}

func TestCountDistinctArrangementsFinal(t *testing.T) {
	r, close := lineReaderFromFile("./inputs/day_10.txt")
	defer close()
	expected := 37024595836928
	actual := countDistinctArragements(r)
	if actual != expected {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.Fail()
	}
}
