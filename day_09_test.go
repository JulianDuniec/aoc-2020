package aoc

import "testing"

const day9Sample = `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`

func TestFindFirstNumberWithoutSumInPreamble(t *testing.T) {
	r := lineReaderFromString(day9Sample)
	expected := 127
	actual := findFirstNumberWithoutSumInPreamble(readNumbers(r), 5)
	if actual != expected {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.Fail()
	}
}

func TestFindFirstNumberWithoutSumInPreambleFinal(t *testing.T) {
	r, close := lineReaderFromFile("./inputs/day_9.txt")
	defer close()
	expected := 26796446
	actual := findFirstNumberWithoutSumInPreamble(readNumbers(r), 25)
	if actual != expected {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.Fail()
	}
}

func TestFindEnchryptionWeakness(t *testing.T) {
	r := lineReaderFromString(day9Sample)
	expected := 62
	actual := findEnchryptionWeakness(readNumbers(r), 5)
	if actual != expected {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.Fail()
	}
}

func TestFindEnchryptionWeaknessFinal(t *testing.T) {
	r, close := lineReaderFromFile("./inputs/day_9.txt")
	defer close()
	expected := 3353494
	actual := findEnchryptionWeakness(readNumbers(r), 25)
	if actual != expected {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.Fail()
	}
}
