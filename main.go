package main

import (
	"fmt"
	"math"
	calculate "math-skills/calculation"
	"math-skills/readfile"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error usage: go run main.go <data.txt>")
		os.Exit(1)
	}

	input := os.Args[1]
	inputData, err := readfile.ReadFile(input)
	if err != nil {
		fmt.Println("Error occurred during inputing data: ", err)
		os.Exit(1)
	}

	var validData []float64
	for _, data := range inputData {
		number, overflow := calculate.ConvFloatToInt(data)
		if overflow {
			continue
		}
		validData = append(validData, float64(number))
	}

	average := calculate.Average(validData)
	median := calculate.Median(validData)
	variance := calculate.Variance(validData)
	stdVarian := math.Sqrt(variance)

	convmedian, _ := calculate.ConvFloatToInt(median)

	fmt.Printf("Average: %.0f\n", average)
	fmt.Printf("Median: %d\n", convmedian)
	fmt.Printf("Variance: %.0f\n", variance)
	fmt.Printf("Standard Deviation: %.0f\n", stdVarian)
}
