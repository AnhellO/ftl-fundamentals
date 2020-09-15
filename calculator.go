// Package calculator provides a library for simple calculations in Go.
package calculator

import "errors"

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
// betwee the second.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return -1, errors.New("Invalid division by 0")
	}

	return a / b, nil
}
