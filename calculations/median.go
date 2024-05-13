package calculate

import "sort"

func MedianCalculation(numbers []float64) float64 {
	// Sort the numbers in ascending order
	sort.Float64s(numbers)

	n := len(numbers)
	if n == 0 {
		return 0
	}

	if n%2 == 1 {
		return numbers[n/2]
	}

	middle := n / 2
	return (numbers[middle-1] + numbers[middle]) / 2
}
