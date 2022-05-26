package main

import (
	"fmt"
)

func main() {
	var response string
	var firstNumber int
	var secondNumber int

	fmt.Print("You want to perform an addition, subtraction, multiplication or division?-->")
	fmt.Scanln(&response)

	if response == "add" {
		fmt.Print("¿What is the first number you want to add?")
		fmt.Scanln(&firstNumber)
		fmt.Print("¿What is the second number you want to add?")
		fmt.Scanln(&secondNumber)

		var add = firstNumber + secondNumber
		fmt.Printf("The result is: %d\n", add)
	}

	if response == "subtraction" {
		fmt.Print("¿What is the first number you want to subtract?")
		fmt.Scanln(&firstNumber)
		fmt.Print("¿What is the second number you want to subtract?")
		fmt.Scanln(&secondNumber)

		var subtraction = firstNumber - secondNumber
		fmt.Printf("The result is: %d\n", subtraction)
	}

	if response == "multiplication" {
		fmt.Print("¿What is the first number you want to multiply? ")
		fmt.Scanln(&firstNumber)
		fmt.Print("¿What is the second number you want to multiply? ")
		fmt.Scanln(&secondNumber)

		var multiplications = firstNumber * secondNumber
		fmt.Printf("The result is: %d\n", multiplications)
	}

	if response == "division" {
		fmt.Print("¿What is the first number you want to divide?")
		fmt.Scanln(&firstNumber)
		fmt.Print("¿What is the first number you want to divide?")
		fmt.Scanln(&secondNumber)

		var dividir = firstNumber / secondNumber
		fmt.Printf("The result is: %d\n", dividir)
	}
}