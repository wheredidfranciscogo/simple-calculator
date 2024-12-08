package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go <number1> <operator> <number2>")
		return
	}
	// Convert the first number
	num1, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Error: Invalid number for num1")
		return
	}

	// Convert the second number
	num2, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println("Error: Invalid number for num2")
		return
	}

	operator := os.Args[2]

	if operator == "+" {
		result := num1 + num2
		fmt.Println("Result:", result)
	} else {
		fmt.Println("Unsupported operator:", operator)
	}

	fmt.Println("Number 1:", num1)
	fmt.Println("Number 2:", num2)
	fmt.Println("Operator:", os.Args[2])
	// fmt.Println("Arguments provided:", os.Args[1:]) // what is the meaning of this 1:
}
