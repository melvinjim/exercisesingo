package models

import (
	"fmt"
)

//multiplicacion de dos datos numericos 
func Multiplication(firstNumber, secondNumber int){
	operation := firstNumber * secondNumber
	fmt.Printf(" --hey--- my first number was: %d and my second number was: %d and I give a total of: %d", firstNumber, secondNumber, operation)
}