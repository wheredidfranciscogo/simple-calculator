package main

import (
	"fmt"
	"math"
)

func main() {
	var num1, num2 float64
	var operator string

	// Prompt for the first number
	fmt.Print("Enter the first number: ")
	_, err := fmt.Scanln(&num1)
	if err != nil {
		fmt.Println("Error: Invalid input")
		return
	}

	// Prompt for the operator
	fmt.Print("Enter the operator (+, -, *, /, %, ^): ")
	_, err = fmt.Scanln(&operator)
	if err != nil {
		fmt.Println("Error: Invalid input")
		return
	}

	// Prompt for the second number
	fmt.Print("Enter the second number: ")
	_, err = fmt.Scanln(&num2)
	if err != nil {
		fmt.Println("Error: Invalid input")
		return
	}

	result := calculate(num1, num2, operator)

	fmt.Printf("Result of %.2f %s %.2f = %.2f\n", num1, operator, num2, result)
}

func calculate(a float64, b float64, operator string) float64 {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			panic("division by zero")
		}
		return a / b
	case "%":
		return math.Mod(a, b)
	case "^":
		return math.Pow(a, b)
	default:
		panic("invalid operator")
	}
}
