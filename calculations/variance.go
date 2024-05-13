package calculate

func calculateVariance(numbers []float64) float64 {
	mean := calculateMean(numbers)
	var sumOfSquaredDeviations float64

	for _, num := range numbers {
		deviation := num - mean
		sumOfSquaredDeviations += deviation * deviation
	}

	n := float64(len(numbers))
	variance := sumOfSquaredDeviations / (n - 1)

	return variance
}

func calculateMean(numbers []float64) float64 {
	var sum float64
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}
