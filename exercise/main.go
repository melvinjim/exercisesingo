package main

import (
	"fmt"
)

func main() {
	var respuesta string
	var firstNumber int
	var secondNumber int

	fmt.Print("desea realizar una suma, resta, multiplicacion o division?-->")
	fmt.Scanln(&respuesta)

	if respuesta == "suma" {
		fmt.Print("¿Cuál es el primer numero desea sumar? ")
		fmt.Scanln(&firstNumber)
		fmt.Print("¿Cuál es el segundo numero que desea sumar? ")
		fmt.Scanln(&secondNumber)

		var suma = firstNumber + secondNumber
		fmt.Printf("el resultado es: %d\n", suma)
	}

	if respuesta == "resta" {
		fmt.Print("¿Cuál es el primer numero desea restar? ")
		fmt.Scanln(&firstNumber)
		fmt.Print("¿Cuál es el segundo numero que desea restar? ")
		fmt.Scanln(&secondNumber)

		var resta = firstNumber - secondNumber
		fmt.Printf("el resultado es: %d\n", resta)
	}

	if respuesta == "multiplicacion" {
		fmt.Print("¿Cuál es el primer desea multiplicar? ")
		fmt.Scanln(&firstNumber)
		fmt.Print("¿Cuál es el segundo numero que desea multiplicar? ")
		fmt.Scanln(&secondNumber)

		var multiplications = firstNumber * secondNumber
		fmt.Printf("el resultado es: %d\n", multiplications)
	}

	if respuesta == "division" {
		fmt.Print("¿Cuál es el primer desea dividir? ")
		fmt.Scanln(&firstNumber)
		fmt.Print("¿Cuál es el segundo numero que desea dividir? ")
		fmt.Scanln(&secondNumber)

		var dividir = firstNumber / secondNumber
		fmt.Printf("el resultado es: %d\n", dividir)
	}
}
