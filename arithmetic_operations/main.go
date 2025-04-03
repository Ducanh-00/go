package main

import (
	"arithmetic_operations/calculator"
	"fmt"
)

func main() {
	calc := &calculator.Calculator{A: 10, B: 5}
	addResult := calculator.PerformCalculation(calc, calc.Add)
	subtractResult := calculator.PerformCalculation(calc, calc.Subtract)
	multiplyResult := calculator.PerformCalculation(calc, calc.Multiply)
	divideResult := calculator.PerformCalculation(calc, calc.Divide)

	fmt.Printf("Addition: %d\n", addResult)
	fmt.Printf("Subtraction: %d\n", subtractResult)
	fmt.Printf("Multiplication: %d\n", multiplyResult)
	fmt.Printf("Division: %d\n", divideResult)

	fmt.Printf("Performing operations using a slice of Calculator structs:\n")
	number := []calculator.Calculator{
		{A: 10, B: 5},
		{A: 20, B: 10},
		{A: 30, B: 15},
		{A: 40, B: 20},
	}

	for _, num := range number {
		fmt.Printf("Addition: %d\n", calculator.PerformCalculation(&num, num.Add))
	}
}
