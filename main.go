package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	// check for valid arguments
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go <number1> <operator> <number2>")
		return
	}

	// Convert the input args
	numbers := []float64{}
	for i, arg := range []string{os.Args[1], os.Args[3]} {
		num, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Printf("Error: Invalid number for argument %d/n", i+1)
			return
		}
		numbers = append(numbers, num)
	}

	num1, num2 := numbers[0], numbers[1]
	operator := os.Args[2]

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
		if b != 0 {
			return a / b
		} else {
			fmt.Println("Error: Division by zero")
			os.Exit(1) // Stop the program if there's a division by zero
		}
	case "%":
		return math.Mod(a, b)
	case "^":
		return math.Pow(a, b)
	default:
		fmt.Println("Invalid operator")
		os.Exit(1)
	}
	return 0
}
