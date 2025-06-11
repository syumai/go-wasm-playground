package main

import (
	"calculator/internal/calculator/calc/calculator"
)

func init() {
	calculator.Exports.Add = func(a, b float64) float64 {
		return a + b
	}
	calculator.Exports.Subtract = func(a, b float64) float64 {
		return a - b
	}
	calculator.Exports.Multiply = func(a, b float64) float64 {
		return a * b
	}
	calculator.Exports.Divide = func(a, b float64) float64 {
		if b == 0 {
			return 0 // Return 0 for division by zero to avoid panic
		}
		return a / b
	}
}

func main() {}