package ml

import (
	"fmt"
	"math"
	"sort"
)

func GuessTrueValue(values []float64, epochs int, lr float64) float64 {
	if len(values) == 0 {
		panic("Input list cannot be empty.")
	}
	sortedValues := make([]float64, len(values))
	copy(sortedValues, values)
	sort.Float64s(sortedValues)
	guess := sortedValues[len(sortedValues)/2]
	for i := 0; i < epochs; i++ {
		v := sortedValues[i%len(sortedValues)]
		diff := v - guess
		weight := 1 / (1 + math.Abs(diff))
		guess += lr * weight * diff
		lr *= 0.999
		fmt.Printf("#%d: guess=%.6f, sample=%.6f, diff=%.6f, weight=%.6f, lr=%.6f\n", i, guess, v, diff, weight, lr)
	}
	return guess
}
