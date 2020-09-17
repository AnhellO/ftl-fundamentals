// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the first
// from the second.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers and returns the result of multiplying the first
// with the second.
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two numbers and returns the result of dividing the first
// between the second. This function can return an error if the divider is '0'.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return -1, errors.New("Invalid division by 0")
	}

	return a / b, nil
}

// SquareRoot takes a number and returns its square root.
// This function can return an error if the number is negative.
func SquareRoot(n float64) (float64, error) {
	if n < 0 {
		return -1, errors.New("Invalid negative number")
	}

	return math.Sqrt(n), nil
}
