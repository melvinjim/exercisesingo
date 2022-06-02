package main

import (
	"calculatorofconsol/models"
	"fmt"
)

func main() {
	var response string

	fmt.Print("You want to perform an addition, subtraction, multiplication or division?-->")
	fmt.Scanln(&response)

	if response == "add" {
		models.Add()
	}

	if response == "subtraction" {
		models.Subtraction()
	}

	if response == "multiplication" {
		models.Multiplication()
	}

	if response == "division" {
		models.Division(10, 2)
	}
}
