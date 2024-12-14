package main

import (
	"fmt"
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
