package aoc

import (
	"testing"
)

const sampleProgram = `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`

func TestAccumulatorValueAtInfiniteLoop(t *testing.T) {
	r := lineReaderFromString(sampleProgram)
	const expected = 5
	actual := accumulatorValueAtInifiniteLoop(r)
	if actual != expected {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.Fail()
	}
}

func TestAccumulatorValueAtInfiniteLoopFinal(t *testing.T) {
	r, close := lineReaderFromFile("./inputs/day_8.txt")
	defer close()
	const expected = 1548
	actual := accumulatorValueAtInifiniteLoop(r)
	if actual != expected {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.Fail()
	}
}

func TestAdjustProgramAndRunToEnd(t *testing.T) {
	r := lineReaderFromString(sampleProgram)
	const expected = 8
	actual := adjustProgramAndRunToEnd(r)
	if actual != expected {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.Fail()
	}
}

func TestAdjustProgramAndRunToEndFinal(t *testing.T) {
	r, close := lineReaderFromFile("./inputs/day_8.txt")
	defer close()
	const expected = 1375
	actual := adjustProgramAndRunToEnd(r)
	if actual != expected {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.Fail()
	}
}
