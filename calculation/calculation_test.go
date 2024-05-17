package calculate_test

import (
	"math"
	calculate "math-skills/calculation"
	"testing"
)

func TestAverage(t *testing.T) {
	tests := []struct {
		name string
		args []float64
		want float64
	}{
		{"Case 1", []float64{1, 2, 3, 4, 5}, 3},
		{"Case 2", []float64{1, 2, 3, 4}, 2.5},
		{"Case 3", []float64{}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate.Average(tt.args); got != tt.want {
				t.Errorf("Average() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMedian(t *testing.T) {
	tests := []struct {
		name string
		args []float64
		want float64
	}{
		{"Case 1", []float64{1, 2, 3, 4, 5}, 3},
		{"Case 2", []float64{1, 2, 3, 4}, 2.5},
		{"Case 3", []float64{}, 0},
		// Add more test cases here
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate.Median(tt.args); got != tt.want {
				t.Errorf("Median() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVariance(t *testing.T) {
	tests := []struct {
		name string
		args []float64
		want float64
	}{
		{"Case 1", []float64{1, 2, 3, 4, 5}, 2},
		{"Case 2", []float64{1, 1, 1, 1, 1}, 0},
		{"Case 3", []float64{}, 0},
		// Add more test cases here
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate.Variance(tt.args); got != tt.want {
				t.Errorf("Variance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvFloatToInt(t *testing.T) {
	tests := []struct {
		name    string
		input   float64
		wantInt int
		wantErr bool
	}{
		{"Case 1", 123.456, 123, false},
		{"Case 2", -123.456, -123, false},
		{"Case 3", math.MaxFloat64, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotInt, gotErr := calculate.ConvFloatToInt(tt.input)
			if gotInt != tt.wantInt || (gotErr != false) != tt.wantErr {
				t.Errorf("ConvFloatToInt() = %v, %v; want %v, %v", gotInt, gotErr, tt.wantInt, tt.wantErr)
			}
		})
	}
}
