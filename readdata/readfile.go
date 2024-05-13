package readdata

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadData(input string) ([]float64, error) {
	file, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sliceArr []float64
	for scanner.Scan() {
		line := scanner.Text()
		content, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Error:", err)
			return nil, err
		}
		sliceArr = append(sliceArr, content)

	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	return sliceArr, nil
}
