package main

import "testing"

func testcalculate(t *testing.T) {
	if calculate(13) != 15 {
		t.Error("Expecting 13 + 2 to equal 15")
	}
}

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{10, 12},
		{-2, 0},
		{9000, 9002},
		{2, 4},
		{0, 2},
	}

	for _, test := range tests {
		if output := calculate(test.input); output != test.expected {
			t.Errorf("Test Failed: {%v} inputted, {%v} expected, {%v} received", test.input, test.expected, output)
			t.Fail()
		}
	}
}
