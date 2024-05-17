package readfile

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadFile(input string) ([]float64, error) {
	// File Opening
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Error usage: go run main.go <data.txt>", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sliceArr []float64 
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		content, err := strconv.ParseFloat(line, 64)
		if err != nil {
			continue
		}
		sliceArr = append(sliceArr, content)

	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	return sliceArr, nil
}
