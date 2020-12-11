package aoc

import (
	"reflect"
	"testing"
)

const (
	adjacencyThresholdVisible = 5
	adjacencyThresholdNear    = 4
)

func TestAdvanceSeatLayout(t *testing.T) {
	sequence := []string{
		`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`,
		`#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`,
		`#.LL.L#.##
#LLLLLL.L#
L.L.L..L..
#LLL.LL.L#
#.LL.LL.LL
#.LLLL#.##
..L.L.....
#LLLLLLLL#
#.LLLLLL.L
#.#LLLL.##`,
		`#.##.L#.##
#L###LL.L#
L.#.#..#..
#L##.##.L#
#.##.LL.LL
#.###L#.##
..#.#.....
#L######L#
#.LL###L.L
#.#L###.##`,
		`#.#L.L#.##
#LLL#LL.L#
L.L.L..#..
#LLL.##.L#
#.LL.LL.LL
#.LL#L#.##
..L.L.....
#L#LLLL#L#
#.LLLLLL.L
#.#L#L#.##`,
		`#.#L.L#.##
#LLL#LL.L#
L.#.L..#..
#L##.##.L#
#.#L.LL.LL
#.#L#L#.##
..L.L.....
#L#L##L#L#
#.LLLLLL.L
#.#L#L#.##`,
	}
	for i := 1; i < len(sequence)-1; i++ {
		p := sequence[i-1]
		expected := parseSeatLayout(lineReaderFromString(sequence[i]))
		m := parseSeatLayout(lineReaderFromString(p))
		actual := advanceSequence(m, numOccupiedAdjacent, adjacencyThresholdNear)
		if reflect.DeepEqual(expected, actual) == false {
			t.Logf("Unexpected advancement in sequence %d-%d \nexpected: %v\nactual:  %v", i-1, i, expected, actual)
			t.FailNow()
		}
	}
}

func TestFindAvailableSeatsAtEquilibrium(t *testing.T) {
	expected := 37
	actual := findAvailableSeatsAtEquilibrium(lineReaderFromString(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`), numOccupiedAdjacent, 4)
	if actual != expected {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.Fail()
	}
}

func TestFindAvailableSeatsAtEquilibriumFinal(t *testing.T) {
	expected := 2448
	reader, close := lineReaderFromFile(`./inputs/day_11.txt`)
	defer close()
	actual := findAvailableSeatsAtEquilibrium(reader, numOccupiedAdjacent, adjacencyThresholdNear)
	if actual != expected {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.Fail()
	}
}

func TestAdvanceSeatLayoutVisible(t *testing.T) {
	sequence := []string{
		`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`,
		`#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`,
		`#.LL.LL.L#
#LLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLLL.L
#.LLLLL.L#`,
		`#.L#.##.L#
#L#####.LL
L.#.#..#..
##L#.##.##
#.##.#L.##
#.#####.#L
..#.#.....
LLL####LL#
#.L#####.L
#.L####.L#`,
		`#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##LL.LL.L#
L.LL.LL.L#
#.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLL#.L
#.L#LL#.L#`,
		`#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##L#.#L.L#
L.L#.#L.L#
#.L####.LL
..#.#.....
LLL###LLL#
#.LLLLL#.L
#.L#LL#.L#`,
		`#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##L#.#L.L#
L.L#.LL.L#
#.LLLL#.LL
..#.L.....
LLL###LLL#
#.LLLLL#.L
#.L#LL#.L#`,
	}
	for i := 1; i < len(sequence)-1; i++ {
		p := sequence[i-1]
		expected := parseSeatLayout(lineReaderFromString(sequence[i]))
		m := parseSeatLayout(lineReaderFromString(p))
		actual := advanceSequence(m, numOccupiedVisible, adjacencyThresholdVisible)
		if reflect.DeepEqual(expected, actual) == false {
			t.Logf("Unexpected advancement in sequence %d-%d \nexpected: %v\nactual:  %v", i-1, i, expected, actual)
			t.FailNow()
		}
	}
}

func TestCountAdjacent(t *testing.T) {
	m := parseSeatLayout(lineReaderFromString(`.......#.
...#.....
.#.......
.........
..#L....#
....#....
.........
#........
...#.....`))

	na := numOccupiedVisible(4, 3, m)
	if na != 8 {
		t.Logf("Expected %d did not match actual %d", 0, na)
		t.Fail()
	}
}

func TestFindAvailableSeatsAtEquilibriumVisibleFinal(t *testing.T) {
	expected := 2234
	reader, close := lineReaderFromFile(`./inputs/day_11.txt`)
	defer close()
	actual := findAvailableSeatsAtEquilibrium(reader, numOccupiedVisible, adjacencyThresholdVisible)
	if actual != expected {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.Fail()
	}
}
