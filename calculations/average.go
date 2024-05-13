package calculate

func CalculateAverage(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	average := sum / float64(len(numbers))
	return average
}
