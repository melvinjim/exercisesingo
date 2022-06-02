package models

import (
	"fmt"
)

func Add() int {
	var firstNumber int
	var secondNumber int

	fmt.Print("¿What is the first number you want to add?\n")
	fmt.Scanln(&firstNumber)
	fmt.Println(firstNumber)
	fmt.Print("¿What is the second number you want to add?\n")
	fmt.Scanln(&secondNumber)
	fmt.Println(secondNumber)

	add := firstNumber + secondNumber
	fmt.Printf("The result is: %d\n\n", add)

	return add
}

func Subtraction() int {
	var firstNumber int
	var secondNumber int

	fmt.Print("¿What is the first number you want to subtract?\n")
	fmt.Scanln(&firstNumber)
	fmt.Println(firstNumber)
	fmt.Print("¿What is the second number you want to subtract?\n")
	fmt.Scanln(&secondNumber)
	fmt.Println(secondNumber)

	var subtraction = firstNumber - secondNumber
	fmt.Printf("The result is: %d\n\n", subtraction)

	return subtraction
}

func Multiplication() int {
	var firstNumber int
	var secondNumber int

	fmt.Print("¿What is the first number you want to multiply?\n")
	fmt.Scanln(&firstNumber)
	fmt.Println(firstNumber)
	fmt.Print("¿What is the second number you want to multiply?\n")
	fmt.Scanln(&secondNumber)
	fmt.Println(secondNumber)

	var multiplications = firstNumber * secondNumber
	fmt.Printf("The result is: %d\n\n", multiplications)

	return multiplications
}

func Division(firstNumber, secondNumber int) int {

	fmt.Print("¿What is the first number you want to divide?\n")
	fmt.Scanln(&firstNumber)
	fmt.Println(firstNumber)
	fmt.Print("¿What is the first number you want to divide?\n")
	fmt.Scanln(&secondNumber)
	fmt.Println(secondNumber)

	var dividir = firstNumber / secondNumber
	fmt.Printf("The result is: %d\n\n", dividir)

	return dividir
}
