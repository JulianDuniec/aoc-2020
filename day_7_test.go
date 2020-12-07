package aoc

import "testing"

const sample = `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`

func TestFindPossibleContainersCount(t *testing.T) {
	expected := 4
	f := lineReaderFromString(sample)

	actual := findPossibleContainersCount(f, "shiny gold")

	if expected != actual {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.Fail()
	}
}

func TestFindPossibleContainersCountFinal(t *testing.T) {
	expected := 335
	f, close := lineReaderFromFile("./inputs/day_7.txt")
	defer close()
	actual := findPossibleContainersCount(f, "shiny gold")

	if expected != actual {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.Fail()
	}
}

func TestFindTotalChildCount(t *testing.T) {
	expected := 32
	f := lineReaderFromString(sample)

	actual := findTotalChildCount(f, "shiny gold")

	if expected != actual {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.Fail()
	}
}

func TestFindTotalChildCountFinal(t *testing.T) {
	expected := 2431
	f, close := lineReaderFromFile("./inputs/day_7.txt")
	defer close()
	actual := findTotalChildCount(f, "shiny gold")

	if expected != actual {
		t.Logf("Expected %d did not match actual %d", expected, actual)
		t.Fail()
	}
}
