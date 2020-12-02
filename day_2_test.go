package aoc

import "testing"

func TestCountValidPasswords(t *testing.T) {
	tcs := []struct {
		input    []password
		expected int
	}{
		{
			input: []password{
				{
					policy: passwordPolicy{
						low:    1,
						high:   3,
						letter: 'a',
					},
					password: "abcde",
				},
			},
			expected: 1,
		},
		{
			input: []password{
				{
					policy: passwordPolicy{
						low:    1,
						high:   3,
						letter: 'b',
					},
					password: "cdefg",
				},
			},
			expected: 0,
		},
		{
			input: []password{
				{
					policy: passwordPolicy{
						low:    2,
						high:   9,
						letter: 'c',
					},
					password: "ccccccccc",
				},
			},
			expected: 1,
		},
		{
			input: []password{
				{
					policy: passwordPolicy{
						low:    1,
						high:   3,
						letter: '😊',
					},
					password: "😊1😊2",
				},
			},
			expected: 1,
		},
	}
	for _, tc := range tcs {
		actual := CountValidPasswords(tc.input)
		if actual != tc.expected {
			t.Logf("unexpected count %d, expected %d for %v", actual, tc.expected, tc.input)
			t.Fail()
		}
	}
}

func TestCountValidPasswordsFinal(t *testing.T) {
	expected := 542
	actual := CountValidPasswords(parseInputAsPasswords("./inputs/day_2.txt"))
	if expected != actual {
		t.Logf("Expected did not match actual %d %d", expected, actual)
		t.FailNow()
	}
}

func TestCountValidPasswordsAlternate(t *testing.T) {
	tcs := []struct {
		input    []password
		expected int
	}{
		{
			input: []password{
				{
					policy: passwordPolicy{
						low:    1,
						high:   3,
						letter: 'a',
					},
					password: "abcde",
				},
			},
			expected: 1,
		},
		{
			input: []password{
				{
					policy: passwordPolicy{
						low:    1,
						high:   3,
						letter: 'b',
					},
					password: "cdefg",
				},
			},
			expected: 0,
		},
		{
			input: []password{
				{
					policy: passwordPolicy{
						low:    2,
						high:   9,
						letter: 'c',
					},
					password: "ccccccccc",
				},
			},
			expected: 0,
		},
		{
			input: []password{
				{
					policy: passwordPolicy{
						low:    1,
						high:   3,
						letter: '😊',
					},
					password: "😊122",
				},
			},
			expected: 1,
		},
		{
			input: []password{
				{
					policy: passwordPolicy{
						low:    1,
						high:   3,
						letter: '😊',
					},
					password: "😊1😊2",
				},
			},
			expected: 0,
		},
	}

	for _, tc := range tcs {
		actual := CountValidPasswordsAlternate(tc.input)
		if actual != tc.expected {
			t.Logf("unexpected count %d, expected %d for %v", actual, tc.expected, tc.input)
			t.Fail()
		}
	}
}

func TestCountValidPasswordsAlternateFinal(t *testing.T) {
	expected := 360
	actual := CountValidPasswordsAlternate(parseInputAsPasswords("./inputs/day_2.txt"))
	if expected != actual {
		t.Logf("Expected did not match actual %d %d", expected, actual)
		t.FailNow()
	}
}
