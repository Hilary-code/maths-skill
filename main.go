package main

import (
	"fmt"
	"os"

	calculate "mathsskillproject/calculations"
	"mathsskillproject/readdata"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: usage go run main.go <data.txt>")
		os.Exit(1)
	}

	input := os.Args[1]
	inputData, err := readdata.ReadData(input)
	if err != nil {
		fmt.Println("Error occurred during reading", err)
		os.Exit(1)
	}
	average := calculate.CalculateAverage(inputData)
	median := calculate.MedianCalculation(inputData)
	variance := calculate.CalculateAverage(inputData)

	fmt.Printf("Average: %d\n", int(average+0.5))
	fmt.Printf("Median: %d\n", int(median+0.5))
	fmt.Printf("Variance: %d\n", int(variance+0.5))
}
