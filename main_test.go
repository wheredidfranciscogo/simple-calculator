package main

import (
	"math"
	"testing"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		num1, num2 float64
		operator   string
		expected   float64
	}{
		{10, 5, "+", 15},
		{10, 5, "-", 5},
		{10, 5, "*", 50},
		{10, 5, "/", 2},
		{10, 5, "%", 0},             // 10 % 5 = 0
		{2, 3, "^", math.Pow(2, 3)}, // 2^3 = 8
	}

	for _, tt := range tests {
		t.Run(tt.operator, func(t *testing.T) {
			result := calculate(tt.num1, tt.num2, tt.operator)
			if math.Abs(result-tt.expected) > 1e-9 {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestCalculateEdgeCases(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for division by zero but didn't get one")
		}
	}()
	calculate(10, 0, "/") // This should panic
}

func TestCalculateModulus(t *testing.T) {
	result := calculate(10, 3, "%")
	expected := 1.0
	if math.Abs(result-expected) > 1e-9 {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}
