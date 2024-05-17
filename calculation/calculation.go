package calculate

import (
	"math"
	"sort"
)

func Average(arr []float64) float64 {
	if len(arr) == 0 {
		return 0
	}

	var sum float64
	for _, v := range arr {
		sum += v
	}
	return sum / float64(len(arr))
}

func Median(data []float64) float64 {
	sort.Float64s(data)
	num := len(data)
	if num == 0 {
		return 0
	}

	median := num / 2
	if num%2 == 0 {
		return (data[median-1] + data[median]) / 2
	} else {
		return data[median]
	}

}

func Variance(data []float64) float64 {
	num := len(data)
	if num == 0 {
		return 0
	}

	var sum float64
	for _, v := range data {
		sum += v
	}
	median := sum / float64(num)

	var squareDifsum float64
	for _, val := range data {
		diff := val - median
		squareDifsum += diff * diff
	}
	Variance := squareDifsum / float64(num)

	return Variance
}

const scalingFactor = 1

func ConvFloatToInt(f float64) (int, bool) {
	validnumber := f * scalingFactor

	if validnumber >= 9223372036854775807 || validnumber <= -9223372036854775807 {
		return 0, true
	}
	return int(math.Round((validnumber))), false

}
