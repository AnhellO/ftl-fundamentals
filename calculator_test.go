package calculator_test

import (
	"calculator"
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type testCase struct {
	name          string
	a, b, want    float64
	expectedError bool
}

func TestAdd(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{name: "Sum two positive numbers", a: 2, b: 2, want: 4},
		{name: "Sum positive and negative number", a: 1, b: -1, want: 0},
		{name: "Sum positive number and zero", a: 5, b: 0, want: 5},
		{name: "Sum negative number and zero", a: -5, b: 0, want: -5},
		{name: "Sum two negative numbers", a: -5, b: -10, want: -15},
		{name: "Sum two fractional numbers", a: -5.5, b: -10.9, want: -16.4},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Case -> %s\nAdd(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{name: "Subtract two positive numbers", a: 2, b: 2, want: 0},
		{name: "Subtract two positive numbers with negative result", a: 1, b: 10, want: -9},
		{name: "Subtract zero", a: 5, b: 0, want: 5},
		{name: "Subtract zero to negative number", a: -5, b: 0, want: -5},
		{name: "Subtract two negative numbers", a: -5, b: -10, want: 5},
		{name: "Subtract two fractional numbers", a: -5.5, b: -10.9, want: 5.4},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Case -> %s\nSubtract(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{name: "Multiply two positive numbers", a: 2, b: 2, want: 4},
		{name: "Multiply positive and negative number", a: 1, b: -1, want: -1},
		{name: "Multiply by zero", a: 5, b: 0, want: 0},
		{name: "Multiply negative number by zero", a: -5, b: 0, want: 0},
		{name: "Multiply two negative numbers", a: -5, b: -10, want: 50},
		{name: "Multiply two fractional numbers", a: -5.5, b: -10.9, want: 59.95},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Case -> %s\nMultiply(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{name: "Divide two positive numbers", a: 4, b: 2, want: 2},
		{name: "Divide positive and negative number", a: 1, b: -1, want: -1},
		{name: "Divide by zero", a: 5, b: 0, want: -1.0, expectedError: true},
		{name: "Divide negative number by zero", a: -5, b: 0, want: -1.0, expectedError: true},
		{name: "Divide two negative numbers", a: -50, b: -10, want: 5},
		{name: "Divide two fractional numbers", a: -505.5, b: -10.5, want: 48.142857},
	}

	const tolerance = .000001
	opt := cmp.Comparer(func(x, y float64) bool {
		return math.Abs(x/y) > tolerance
	})

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		errReceived := err != nil
		if tc.expectedError != errReceived {
			t.Fatalf("Case -> %s\nDivide(%f, %f): unexpected error status '%v'", tc.name, tc.a, tc.b, errReceived)
		}

		if !tc.expectedError && !cmp.Equal(tc.want, got, opt) {
			t.Errorf("Case -> %s\nDivide(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}
