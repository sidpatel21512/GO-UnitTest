package main

import "testing"

func Test_lcm(t * testing.T) {
	inputs := []struct {
		a   int
		b   int
		expected int
	}{
		// Prime Numbers
		{a: 3, b: 5, expected: 15},
		{a: 2, b: 3, expected: 6},
		{a: 9, b: 7, expected: 63},	
		// Greater Numbers
		{a: 2, b: 4, expected: 4},	
		{a: 3, b: 9, expected: 9},	
		{a: 63, b: 21, expected: 63},	
		// GCD Numbers
		{a: 15, b: 20, expected: 60},	
		{a: 14, b: 21, expected: 42},	
		{a: 18, b: 12, expected: 36},	
		// Validation
			// One number is 0
		{a: 0, b: 0, expected: 0},
			// One number is 1
		{a: 6, b: 1, expected: 6},
			// Negative Numbers	
		{a: 2, b: -5, expected: 10},
			// Float Numbers	
		// {a: 1.75, b: 8.25, expected: 10},

	}

	for _,input := range inputs{
		result := lcm(input.a, input.b)
		if result != input.expected {
			t.Errorf("\"lcm(%d, %d)\" FAILED, expected -> %v, got -> %v", input.a, input.b, input.expected, result)
		} else {
			t.Logf("\"lcm(%d, %d)\" SUCCEDED, expected -> %v, got -> %v", input.a, input.b, input.expected, result)
		}
	}	
}