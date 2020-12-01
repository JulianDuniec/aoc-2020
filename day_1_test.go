package aoc

import (
	"testing"
)

func TestSumAndMultiply(t *testing.T) {
	res := FindSumsAndMultiply([]int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	})
	if res != 514579 {
		t.Logf("invalid answer %d", res)
		t.Fail()
	}
}

// BenchmarkSumAndMultiply-8   	  154111	      7520 ns/op	    3273 B/op	       7 allocs/op
func BenchmarkSumAndMultiply(b *testing.B) {
	input := inputIntSlice("./inputs/day_1.txt")
	for i := 0; i < b.N; i++ {
		FindSumsAndMultiply(input)
	}
}

func TestSumAndMultiplyFinal(t *testing.T) {
	expected := 259716
	actual := FindSumsAndMultiply(inputIntSlice("./inputs/day_1.txt"))
	if expected != actual {
		t.Logf("Expected did not match actual %d %d", expected, actual)
	}
}

func TestThreeSumsAndMultiply(t *testing.T) {
	res := FindThreeSumsAndMultiply([]int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	})
	if res != 241861950 {
		t.Logf("invalid answer %d", res)
		t.Fail()
	}
}

// BenchmarkThreeSumsAndMultiply-8   	  206443	      5464 ns/op	      64 B/op	       2 allocs/op
func BenchmarkThreeSumsAndMultiply(b *testing.B) {
	input := inputIntSlice("./inputs/day_1.txt")
	for i := 0; i < b.N; i++ {
		FindThreeSumsAndMultiply(input)
	}
}

func TestThreeSumsAndMultiplyFinal(t *testing.T) {
	expected := 120637440
	actual := FindThreeSumsAndMultiply(inputIntSlice("./inputs/day_1.txt"))
	if expected != actual {
		t.Logf("Expected did not match actual %d %d", expected, actual)
	}
}
