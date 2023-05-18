package main

import "math"

func lcm_greater(a int, b int) int {
	max := int(math.Max(float64(a),float64(b)))
	min :=int(math.Min(float64(a),float64(b)))

	if max % min == 0 {
		return max;
	}
	return a * b;
}